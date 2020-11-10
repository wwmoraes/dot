package dot

import (
	"testing"
)

func TestUIDGenerator(t *testing.T) {
	expectedLength := 16
	generator := newUIDGenerator(expectedLength)

	if generator == nil {
		t.Error("generator is nil")
		return
	}

	// test default length
	{
		random := generator.String()
		gotLength := len(random)
		if gotLength != expectedLength {
			t.Errorf("got [%v] want [%v]", gotLength, expectedLength)
		}
	}

	// test explicit length
	{
		expectedLength = 32
		random := generator.Stringn(expectedLength)
		gotLength := len(random)
		if gotLength != expectedLength {
			t.Errorf("got [%v] want [%v]", gotLength, expectedLength)
		}
	}
}
