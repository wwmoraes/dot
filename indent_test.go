package dot

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestIndentWriter(t *testing.T) {
	b := new(bytes.Buffer)
	i := NewIndentWriter(b)
	if _, err := i.WriteString("doc {"); err != nil {
		t.Errorf("unable to open doc block: %w", err)
		return
	}
	i.NewLineIndentWhile(func() {
		fmt.Fprint(i, "chapter {")
		i.NewLineIndentWhile(func() {
			fmt.Fprint(i, "chapter text")
		})
		if _, err := i.WriteString("}"); err != nil {
			t.Errorf("unable to close chapter block: %w", err)
			return
		}
	})
	if _, err := i.WriteString("}"); err != nil {
		t.Errorf("unable to close doc block: %w", err)
		return
	}
	got := b.String()
	want := "doc {\n\tchapter {\n\t\tchapter text\n\t}\n}"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		t.Fail()
	}
}
