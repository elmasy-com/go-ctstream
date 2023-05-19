package ctstream

import (
	"context"
	"errors"
	"fmt"
	"sync"

	ct "github.com/google/certificate-transparency-go"
	"github.com/google/certificate-transparency-go/x509"
)

var ErrNothingToDo = errors.New("nothing to do")

type indexRange struct {
	Start int64
	End   int64
}

type Entry struct {
	Index          int
	Precertificate bool
	Certificate    *x509.Certificate
}

type Scanner struct {
	EntryChan   chan *Entry // Channel to send/receive entries form the log
	ErrChan     chan error  // Channel to send/receive errors. Any error will gracefully stop the scanner.
	End         int         // End index, the size of the log
	BatchSize   int         // Batch size
	log         *Log
	irc         chan *indexRange
	ctx         context.Context
	cancel      context.CancelFunc
	wg          *sync.WaitGroup
	skipPrecert bool
	start       int
}

// Stop gracefully stops the scanner.
func (s *Scanner) Stop() {
	s.cancel()
	s.wg.Wait()
}

// indexPartitioner creates equal parts form a range.
// Batch is the "number of elements" in the part.
// eg.:
//
//	start: 0
//	end: 23
//	batch: 10
//	result: 0..9, 10..19, 20..23
func (s *Scanner) indexPartitioner() {

	defer s.wg.Done()
	defer close(s.irc)

	for s.start <= s.End {

		select {

		case <-s.ctx.Done():
			return
		default:

			// Remove one from the end index, because the start index is included
			// eg.: a batch with size of 10 starting from 0 is 0,1,2,3,4,5,6,7,8,9

			ir := indexRange{Start: int64(s.start), End: int64(s.start + s.BatchSize - 1)}

			// Check overflow
			if s.End-s.start < s.BatchSize {
				ir.End = int64(s.End)
			}

			s.irc <- &ir

			s.start += s.BatchSize
		}
	}
}

// fetcher is goroutine function to fetch entries and send the x509 cert's through the EntryChan
func (s *Scanner) fetcher() {

	defer s.wg.Done()

	lclient, err := s.log.LogClient()
	if err != nil {
		s.ErrChan <- fmt.Errorf("failed to create log client: %w", err)
		s.cancel()
		return
	}

	for ir := range s.irc {

		// Logs can retun fewer enries that the requested.
		for ir.Start < ir.End {

			select {
			case <-s.ctx.Done():
				return
			default:

				rawEntries, err := lclient.GetRawEntries(context.Background(), ir.Start, ir.End)
				if err != nil {
					s.ErrChan <- fmt.Errorf("failed to get raw entries %d..%d: %w", ir.Start, ir.End, err)
					s.cancel()
					return
				}

				for i := range rawEntries.Entries {

					rawEntry, err := ct.RawLogEntryFromLeaf(ir.Start, &rawEntries.Entries[i])
					if err != nil {
						s.ErrChan <- fmt.Errorf("failed to convert leaf to raw log entry: %w", err)
						s.cancel()
						return
					}

					ir.Start++

					if rawEntry.Leaf.TimestampedEntry.EntryType == ct.PrecertLogEntryType && s.skipPrecert {
						continue
					}

					entry, err := rawEntry.ToLogEntry()
					if err != nil {
						s.ErrChan <- fmt.Errorf("failed to convert raw to leg entry: %w", err)
						s.cancel()
						return
					}

					e := Entry{Index: int(entry.Index), Precertificate: rawEntry.Leaf.TimestampedEntry.EntryType == ct.PrecertLogEntryType}

					if e.Precertificate {
						if entry.Precert != nil && entry.Precert.TBSCertificate != nil {
							e.Certificate = entry.Precert.TBSCertificate
						}
					} else {
						if entry.X509Cert != nil {
							e.Certificate = entry.X509Cert
						}
					}

					s.EntryChan <- &e
				}
			}
		}
	}
}

// Starts the workers and closes the channels when finished
func (s *Scanner) startBackground(fetcherNum int) {

	s.wg.Add(1)
	go s.indexPartitioner()

	for i := 0; i < fetcherNum; i++ {

		s.wg.Add(1)
		go s.fetcher()
	}

	s.wg.Wait()
	close(s.EntryChan)
	close(s.ErrChan)
}

// NewScanner creates and starts a log scanner.
// The fparsed entries are returned in the Scanner.EntryChan channale.
// Any error will gracefully stops the scanner and returned in the Scanner.ErrChan. Multiple errors are possible.
// fetcherNum is the number of concurrent fetcher.
// If skipPrecert is true, than EntryChan contains only leaf certificates.
//
// Returns ErrNothingToDo if the start index is equal to the last index.
func NewScanner(ctx context.Context, log *Log, start int, fetcherNum int, skipPrecert bool) (*Scanner, error) {

	if log == nil {
		return nil, fmt.Errorf("log is nil")
	}
	if fetcherNum == 0 {
		return nil, fmt.Errorf("fetcher is zero")
	}

	logSize, err := log.Size()
	if err != nil {
		return nil, fmt.Errorf("failed to get log size: %w", err)
	}

	if start >= logSize-1 {
		return nil, ErrNothingToDo
	}

	batchSize, err := log.BatchSize()
	if err != nil {
		return nil, fmt.Errorf("failed to get batch size: %w", err)
	}

	scnnr := Scanner{}
	scnnr.log = log
	scnnr.start = start
	scnnr.skipPrecert = skipPrecert

	// indexRange channel
	scnnr.irc = make(chan *indexRange)

	// Buffered. Number of fetchers + indexPartitioner
	scnnr.ErrChan = make(chan error, fetcherNum+1)
	scnnr.EntryChan = make(chan *Entry)

	scnnr.ctx, scnnr.cancel = context.WithCancel(ctx)

	scnnr.wg = new(sync.WaitGroup)

	scnnr.End = logSize

	scnnr.BatchSize = batchSize

	go scnnr.startBackground(fetcherNum)

	return &scnnr, nil
}
