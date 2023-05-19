package ctstream

import (
	"context"
	"sync"
	"testing"
)

func TestIndexPartitioner(t *testing.T) {

	s := Scanner{}
	s.wg = new(sync.WaitGroup)
	s.wg.Add(1)

	s.irc = make(chan *indexRange)

	s.start = 0
	s.End = 23
	s.BatchSize = 10
	s.ctx = context.TODO()

	go s.indexPartitioner()

	if ir := <-s.irc; ir.Start != 0 || ir.End != 9 {
		t.Fatalf("Invalid range: want: 0..9, got: %d..%d\n", ir.Start, ir.End)
	}

	if ir := <-s.irc; ir.Start != 10 || ir.End != 19 {
		t.Fatalf("Invalid range: want: 10..19, got: %d..%d\n", ir.Start, ir.End)
	}

	if ir := <-s.irc; ir.Start != 20 || ir.End != 23 {
		t.Fatalf("Invalid range: want: 20..23, got: %d..%d\n", ir.Start, ir.End)
	}
}

func TestNewScanner(t *testing.T) {

	log := FindByName("Oak2024H2")
	if log == nil {
		t.Fatalf("Failed to find Oak2024H2 log\n")
	}

	scanner, err := NewScanner(context.TODO(), log, 0, 20, false)
	if err != nil {
		t.Fatalf("Failed: %s\n", err)
	}

	t.Logf("Size of the log: %d\n", scanner.End)

	totalNum := 0

	for range scanner.EntryChan {
		totalNum++
	}

	for err = range scanner.ErrChan {
		t.Fatalf("Error: %s\n", err)
	}

	scanner.Stop()

	t.Logf("Total log: %d\n", totalNum)
}
