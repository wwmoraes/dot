package attributes

import (
	"reflect"
	"strings"
	"testing"

	"github.com/wwmoraes/dot/constants"
)

func TestAttributes(t *testing.T) {
	t.Run("equal on attribute re-set with same value", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttribute(constants.KeyLabel, NewString("test"))

		got := attributes.GetAttributes()

		want := Map{
			constants.KeyLabel: NewString("test"),
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
	t.Run("get attribute previously set", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttribute(constants.KeyLabel, NewString("test"))

		got, _ := attributes.GetAttribute(constants.KeyLabel)

		want := NewString("test")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
}

func TestAttributes_NewAttributes(t *testing.T) {
	t.Run("empty initialization", func(t *testing.T) {
		attributes := NewAttributes()

		got := attributes.GetAttributes()

		want := Map{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
	t.Run("write nothing if empty", func(t *testing.T) {
		attributes := NewAttributes()

		var gotStringBuilder strings.Builder
		_, err := attributes.WriteTo(&gotStringBuilder)
		if err != nil {
			t.Fatal(err)
		}
		got := gotStringBuilder.String()

		want := ""

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
}

func TestAttributes_NewAttributesFrom(t *testing.T) {
	t.Run("empty initialization", func(t *testing.T) {
		attributes := NewAttributesFrom(nil)

		got := attributes.GetAttributes()

		want := Map{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
	t.Run("copy attributes given on initialization", func(t *testing.T) {
		sourceAttributes := NewAttributes()
		sourceAttributes.SetAttributeString(constants.KeyLabel, "test-label")
		attributes := NewAttributesFrom(sourceAttributes)

		var gotStringBuilder strings.Builder
		_, err := attributes.WriteTo(&gotStringBuilder)
		if err != nil {
			t.Fatal(err)
		}
		got := gotStringBuilder.String()

		var wantStringBuilder strings.Builder
		_, err = sourceAttributes.WriteTo(&wantStringBuilder)
		if err != nil {
			t.Fatal(err)
		}
		want := wantStringBuilder.String()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
}

func TestAttributes_GetAttributes(t *testing.T) {
	t.Run("does not mutate using GetAttributes copy map", func(t *testing.T) {
		attributes := NewAttributes()
		value := NewString("test")
		attributes.SetAttribute(constants.KeyLabel, value)
		indirectAttributes := attributes.GetAttributes()
		indirectAttributes[constants.KeyClass] = NewString("my-classconstants.")

		got := attributes.GetAttributes()

		want := Map{
			constants.KeyLabel: value,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
	t.Run("mutates using getAttributes map reference", func(t *testing.T) {
		attributes := NewAttributes()
		labelValue := NewString("test")
		attributes.SetAttribute(constants.KeyLabel, labelValue)
		indirectAttributes := attributes.getAttributes()
		classValue := NewString("my-class")
		indirectAttributes[constants.KeyClass] = classValue

		got := attributes.GetAttributes()

		want := Map{
			constants.KeyLabel: labelValue,
			constants.KeyClass: classValue,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
	t.Run("get single attribute as string", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttribute(constants.KeyLabel, NewHTML("<b>html label</b>"))

		got := attributes.GetAttributeString(constants.KeyLabel)

		want := "<b>html label</b>"

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
	t.Run("get single unset attribute as string", func(t *testing.T) {
		attributes := NewAttributes()

		got := attributes.GetAttributeString(constants.KeyLabel)

		want := ""

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
}

func TestAttributes_WriteTo(t *testing.T) {
	t.Run("writes single string attribute without brackets", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttribute(constants.KeyLabel, NewString("test"))

		var gotStringBuilder strings.Builder
		_, err := attributes.WriteTo(&gotStringBuilder)
		if err != nil {
			t.Fatal(err)
		}
		got := gotStringBuilder.String()

		want := `[label="test"]`

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
	t.Run("writes single string attribute with brackets", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttribute(constants.KeyLabel, NewString("test"))

		var gotStringBuilder strings.Builder
		_, err := attributes.WriteTo(&gotStringBuilder)
		if err != nil {
			t.Fatal(err)
		}
		got := gotStringBuilder.String()

		want := `[label="test"]`

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
	t.Run("writes single HTML attribute without brackets", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttributeHTML(constants.KeyLabel, "<B>Hi</B>")

		var gotStringBuilder strings.Builder
		_, err := attributes.WriteTo(&gotStringBuilder)
		if err != nil {
			t.Fatal(err)
		}
		got := gotStringBuilder.String()

		want := "[label=<<B>Hi</B>>]"

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
	t.Run("writes single HTML attribute with brackets", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttributeHTML(constants.KeyLabel, "<B>Hi</B>")

		var gotStringBuilder strings.Builder
		_, err := attributes.WriteTo(&gotStringBuilder)
		if err != nil {
			t.Fatal(err)
		}
		got := gotStringBuilder.String()

		want := "[label=<<B>Hi</B>>]"

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
	t.Run("writes single Literal attribute without brackets", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttributeLiteral(constants.KeyLabel, `"left text\l"`)

		var gotStringBuilder strings.Builder
		_, err := attributes.WriteTo(&gotStringBuilder)
		if err != nil {
			t.Fatal(err)
		}
		got := gotStringBuilder.String()

		want := `[label="left text\l"]`

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
	t.Run("writes single Literal attribute with brackets", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttributeLiteral(constants.KeyLabel, `"left text\l"`)

		var gotStringBuilder strings.Builder
		_, err := attributes.WriteTo(&gotStringBuilder)
		if err != nil {
			t.Fatal(err)
		}
		got := gotStringBuilder.String()

		want := `[label="left text\l"]`

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
	t.Run("writes multi attributes without brackets", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttributesString(MapString{
			constants.KeyClass: "my-class",
			constants.KeyLabel: "my-label",
		})

		var gotStringBuilder strings.Builder
		_, err := attributes.WriteTo(&gotStringBuilder)
		if err != nil {
			t.Fatal(err)
		}
		got := gotStringBuilder.String()

		want := `[class="my-class",label="my-label"]`

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
	t.Run("writes multi attributes with brackets", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttributesString(MapString{
			constants.KeyClass: "my-class",
			constants.KeyLabel: "my-label",
		})

		var gotStringBuilder strings.Builder
		_, err := attributes.WriteTo(&gotStringBuilder)
		if err != nil {
			t.Fatal(err)
		}
		got := gotStringBuilder.String()

		want := `[class="my-class",label="my-label"]`

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
}

func TestAttributes_SetAttribute(t *testing.T) {
	t.Run("set attribute using single attribute set methods", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttributeString(constants.KeyClass, "my-class")
		attributes.SetAttributeHTML(constants.KeyLabel, "<b>my-label</b>")
		attributes.SetAttributeLiteral(constants.KeyXlabel, `"left text\l"`)
		attributes.SetAttribute(constants.KeyColor, NewString("black"))

		var gotStringBuilder strings.Builder
		_, err := attributes.WriteTo(&gotStringBuilder)
		if err != nil {
			t.Fatal(err)
		}
		got := gotStringBuilder.String()

		want := `[class="my-class",color="black",label=<<b>my-label</b>>,xlabel="left text\l"]`

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
	t.Run("set attribute using multi attribute set methods", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttributesString(MapString{
			constants.KeyClass: "my-class",
		})
		attributes.SetAttributesHTML(MapString{
			constants.KeyLabel: "<b>my-label</b>",
		})
		attributes.SetAttributesLiteral(MapString{
			constants.KeyXlabel: `"left text\l"`,
		})
		attributes.SetAttributes(Map{
			constants.KeyColor: NewString("black"),
		})

		var gotStringBuilder strings.Builder
		_, err := attributes.WriteTo(&gotStringBuilder)
		if err != nil {
			t.Fatal(err)
		}
		got := gotStringBuilder.String()

		want := `[class="my-class",color="black",label=<<b>my-label</b>>,xlabel="left text\l"]`

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
}

func TestAttributes_DeleteAttribute(t *testing.T) {
	t.Run("try to delete un-existent attribute", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.DeleteAttribute(constants.KeyClass)

		got := attributes.GetAttributes()

		want := Map{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
	t.Run("delete a set attribute", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttribute(constants.KeyLabel, NewString("test"))
		attributes.DeleteAttribute(constants.KeyLabel)

		got := attributes.GetAttributes()

		want := Map{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
}
