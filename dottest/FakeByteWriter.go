package dottest

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing"
)

// ErrLimit means that the test writer has reached its limit
var ErrLimit = errors.New("unable to write more: the limit has been reached")

// ByteWriter is implemented by values that can have only bytes written to
type ByteWriter interface {
	fmt.Stringer
	io.Writer
}

type byteWriter struct {
	buffer bytes.Buffer
	err    error
	count  int
	limit  int
}

// NewByteWriter can be written bytes up to limit times before returning err
func NewByteWriter(tb testing.TB, limit int, err error) ByteWriter {
	tb.Helper()

	return &byteWriter{
		limit: limit,
		err:   err,
	}
}

// String returns the buffer value as string
func (fw *byteWriter) String() string {
	return fw.buffer.String()
}

// Write writes the data into buffer, or returns an error if over the limit
func (fw *byteWriter) Write(p []byte) (int, error) {
	if fw.count < fw.limit {
		fw.count++
		n, err := fw.buffer.Write(p)
		if err != nil {
			return n, fmt.Errorf("failed to write to buffer: %w", err)
		}
		if n != len(p) {
			return n, fmt.Errorf("partial write on the buffer happened")
		}
		return len(p), nil
	}

	if fw.err != nil {
		return 0, fw.err
	}

	return 0, ErrLimit
}
