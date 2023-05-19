package ctstream

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/google/certificate-transparency-go/client"
	"github.com/google/certificate-transparency-go/jsonclient"
)

type Log struct {
	Name   string // Log name
	URI    string // Log URI
	PubKey string // Base64 encoded string of the log's public key
}

// KeyToDer returns the log's public key in DER format.
func (l *Log) KeyToDer() ([]byte, error) {

	if l == nil {
		return nil, fmt.Errorf("log is nil")
	}
	if l.PubKey == "" {
		return nil, fmt.Errorf("key is empty")
	}

	return base64.StdEncoding.DecodeString(l.PubKey)
}

// LogClient creates a new client.LogClient for parse logs.
func (l *Log) LogClient() (*client.LogClient, error) {

	if l == nil {
		return nil, fmt.Errorf("log is nil")
	}
	if l.URI == "" {
		return nil, fmt.Errorf("uri is empty")
	}

	keyDer, err := l.KeyToDer()
	if err != nil {
		return nil, fmt.Errorf("failed to decode key: %w", err)
	}

	return client.New(l.URI, http.DefaultClient, jsonclient.Options{PublicKeyDER: keyDer})
}

// Size returns the log size from the SignedTreeHead.
func (l *Log) Size() (int, error) {

	if l == nil {
		return 0, fmt.Errorf("log is nil")
	}
	if l.URI == "" {
		return 0, fmt.Errorf("uri is empty")
	}

	lclient, err := l.LogClient()
	if err != nil {
		return 0, fmt.Errorf("failed to create log client: %w", err)
	}

	sth, err := lclient.GetSTH(context.TODO())
	if err != nil {
		return 0, fmt.Errorf("failed to get signed tree head: %w", err)
	}

	return int(sth.TreeSize), nil
}

// BatchSize returns the maximum number of entries that can retrieved at once.
func (l *Log) BatchSize() (int, error) {

	if l == nil {
		return 0, fmt.Errorf("log is nil")
	}
	if l.URI == "" {
		return 0, fmt.Errorf("uri is empty")
	}

	lclient, err := l.LogClient()
	if err != nil {
		return 0, fmt.Errorf("failed to create log client: %w", err)
	}

	resp, err := lclient.GetRawEntries(context.Background(), 0, 10240)
	if err != nil {
		return 0, fmt.Errorf("failed to get raw entries: %w", err)
	}

	return len(resp.Entries), nil
}
