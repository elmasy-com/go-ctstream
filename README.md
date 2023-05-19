# go-ctstream

[![Go Reference](https://pkg.go.dev/badge/github.com/elmasy-com/go-ctstream.svg)](https://pkg.go.dev/github.com/elmasy-com/go-ctstream)
[![Go Report Card](https://goreportcard.com/badge/github.com/elmasy-com/go-ctstream)](https://goreportcard.com/report/github.com/elmasy-com/go-ctstream)

## Install

```
go get -u github.com/elmasy-com/go-ctstream
```

## Example

```go
    log := ctstream.FindByName("Oak2024H1")
	if log == nil {
		// Handle error
	}

    // Starts a new scan from index 0 with 10 concurrent fetcher.
	scanner, err := ctstream.NewScanner(log, 0, 10, false)
	if err != nil {
		// Handle error
	}

	for entry := range scanner.EntryChan {
		// Do something
	}

	for err := range s.ErrChan {
		// Handle errors
	}
```

## Test

```
go test -v -race
```