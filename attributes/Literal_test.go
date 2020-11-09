package attributes

import (
	"reflect"
	"testing"
)

func TestLiteral(t *testing.T) {
	t.Run("outputs Literal string", func(t *testing.T) {
		expectedValue := `"left-justified text\l"`
		literalString := NewLiteral(expectedValue)
		gotValue := literalString.String()

		if !reflect.DeepEqual(gotValue, expectedValue) {
			t.Errorf("got [%v] want [%v]", gotValue, expectedValue)
		}
	})
}
