// Package dottest provides helper values to create tests in a controlled and
// replicable environment, such as write errors
package dottest

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/wwmoraes/dot/attributes"
)

// ErrLimit means that the test writer has reached its limit
var ErrLimit = errors.New("unable to write more: the limit has been reached")

// WriterStringer is implemented by io.Writer + fmt.Stringer values
type WriterStringer interface {
	fmt.Stringer
	io.Writer
}

// TestByteWrite asserts the outcome of a WriteTo operation on a ByteWriter
func TestByteWrite(tb testing.TB, value io.WriterTo, limit int, wantErr error, wantN int64, wantString string) {
	tb.Helper()

	writer := NewByteWriter(tb, limit, wantErr)

	TestWrite(tb, writer, value, wantErr, wantN, wantString)
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

// MustGetSerializableString returns value as string or Fatal's the test on error
func MustGetSerializableString(tb testing.TB, value attributes.Serializable) string {
	tb.Helper()

	gotString, gotErr := value.String()

	if gotErr != nil {
		tb.Fatalf("unexpected %++v String() error: %v", value, gotErr)
	}

	return gotString
}

// MustGetFlattenSerializableString returns value as a flattened string or
// Fatal's the test on error
func MustGetFlattenSerializableString(tb testing.TB, value attributes.Serializable) string {
	tb.Helper()

	return Flatten(tb, MustGetSerializableString(tb, value))
}
