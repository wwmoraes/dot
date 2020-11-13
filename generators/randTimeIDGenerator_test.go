package generators

import (
	"testing"
)

func TestRandTimeIDGenerator(t *testing.T) {
	t.Run("generate default length random string", func(t *testing.T) {
		expectedLength := 16
		generator := NewRandTimeIDGenerator(expectedLength)

		if generator == nil {
			t.Error("generator is nil")
			t.FailNow()
		}

		random := generator.String()
		gotLength := len(random)

		if gotLength != expectedLength {
			t.Errorf("got [%v] want [%v]", gotLength, expectedLength)
		}
	})

	t.Run("generate provided length random string", func(t *testing.T) {
		expectedLength := 16
		generator := NewRandTimeIDGenerator(expectedLength)

		if generator == nil {
			t.Error("generator is nil")
			t.FailNow()
		}

		expectedLength = expectedLength * 2
		random := generator.Stringn(expectedLength)
		gotLength := len(random)

		if gotLength != expectedLength {
			t.Errorf("got [%v] want [%v]", gotLength, expectedLength)
		}
	})
}
