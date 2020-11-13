package dottest

import (
	"io"
	"testing"
)

// TestByteWrite asserts the outcome of a WriteTo operation on a ByteWriter
func TestByteWrite(tb testing.TB, value io.WriterTo, limit int, wantErr error, wantN int64, wantString string) {
	tb.Helper()

	writer := NewByteWriter(tb, limit, wantErr)

	TestWrite(tb, writer, value, wantErr, wantN, wantString)
}
