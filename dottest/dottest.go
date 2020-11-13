// Package dottest provides helper values to create tests in a controlled and
// replicable environment, such as write errors
package dottest

import (
	"errors"
	"fmt"
	"io"
	"testing"
)

// WriterStringer is implemented by io.Writer + fmt.Stringer values
type WriterStringer interface {
	fmt.Stringer
	io.Writer
}

// TestWrite writes value into writer and then assert the result
func TestWrite(tb testing.TB, writer WriterStringer, value io.WriterTo, wantErr error, wantN int64, wantString string) {
	tb.Helper()

	gotN, gotErr := value.WriteTo(writer)
	gotString := writer.String()

	if gotN != wantN {
		tb.Errorf("got [\n%v\n] want [\n%v\n]", gotN, wantN)
	}

	if !errors.Is(gotErr, wantErr) {
		tb.Errorf("got [\n%v\n] want [\n%v\n]", gotErr, wantErr)
	}

	if gotString != wantString {
		tb.Errorf("got [\n%v\n] want [\n%v\n]", gotString, wantString)
	}
}
