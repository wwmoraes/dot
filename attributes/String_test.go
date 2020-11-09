package attributes

import (
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	t.Run("outputs plain string", func(t *testing.T) {
		expectedValue := "my-label"
		literalString := NewString(expectedValue)
		gotValue := literalString.String()

		if !reflect.DeepEqual(gotValue, expectedValue) {
			t.Errorf("got [%v] want [%v]", gotValue, expectedValue)
		}
	})
}
