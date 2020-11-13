package dottest

import (
	"strings"
	"testing"
)

// Flatten removes line breaks and indentation
func Flatten(tb testing.TB, s string) string {
	tb.Helper()

	s = strings.ReplaceAll(s, "  ", "")
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, "\n", "")

	return s
}
