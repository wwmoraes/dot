package dottest

import (
	"testing"

	"github.com/wwmoraes/dot/attributes"
)

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
