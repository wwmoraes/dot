package dottest

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

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

func (fw *byteWriter) String() string {
	return fw.buffer.String()
}

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
