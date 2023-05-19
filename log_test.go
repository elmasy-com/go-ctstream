package ctstream

import (
	"strings"
	"testing"
)

func TestSize(t *testing.T) {

	for i := range Logs {

		size, err := Logs[i].Size()
		if err != nil {
			t.Fatalf("Size of %s failed: %s\n", Logs[i].Name, err)
		}

		t.Logf("Size of %s: %d\n", Logs[i].Name, size)
	}
}

func TestBatchSize(t *testing.T) {

	for i := range Logs {

		size, err := Logs[i].BatchSize()
		if err != nil && !strings.Contains(err.Error(), "429 Too Many Requests") {
			t.Fatalf("BatchSize of %s failed: %s\n", Logs[i].Name, err)
		}

		t.Logf("BatchSize of %s: %d\n", Logs[i].Name, size)
	}
}
