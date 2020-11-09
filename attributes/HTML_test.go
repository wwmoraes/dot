package attributes

import (
	"reflect"
	"testing"
)

func TestHTML(t *testing.T) {
	t.Run("outputs HTML string", func(t *testing.T) {
		expectedValue := "<B>Hi</B>"
		htmlString := NewHTML(expectedValue)
		gotValue := htmlString.String()

		if !reflect.DeepEqual(gotValue, expectedValue) {
			t.Errorf("got [%v] want [%v]", gotValue, expectedValue)
		}
	})
}
