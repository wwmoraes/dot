package attributes

import (
	"reflect"
	"strings"
	"testing"
)

func TestAttributes(t *testing.T) {
	t.Run("equal on attribute re-set with same value", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttribute(KeyLabel, NewString("test"))

		gotMap := attributes.GetAttributes()
		expectedMap := Map{
			KeyLabel: NewString("test"),
		}

		if !reflect.DeepEqual(gotMap, expectedMap) {
			t.Errorf("got [%v] want [%v]", gotMap, expectedMap)
		}
	})
	t.Run("get attribute previously set", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttribute(KeyLabel, NewString("test"))

		expectedLabelValue := NewString("test")

		gotValue := attributes.GetAttribute(KeyLabel)
		if !reflect.DeepEqual(gotValue, expectedLabelValue) {
			t.Errorf("got [%v] want [%v]", gotValue, expectedLabelValue)
		}
	})
}

func TestAttributes_NewAttributes(t *testing.T) {
	t.Run("empty initialization", func(t *testing.T) {
		attributes := NewAttributes()
		expectedAttributes := Map{}
		if !reflect.DeepEqual(attributes.attributes, expectedAttributes) {
			t.Errorf("got [%v] want [%v]", attributes, expectedAttributes)
		}
	})
	t.Run("write nothing if empty", func(t *testing.T) {
		attributes := NewAttributes()
		expectedString := ""
		var gotStringBuilder strings.Builder
		attributes.Write(&gotStringBuilder, false)
		gotString := gotStringBuilder.String()

		if !reflect.DeepEqual(gotString, expectedString) {
			t.Errorf("got [%v] want [%v]", gotString, expectedString)
		}
	})
}

func TestAttributes_GetAttributes(t *testing.T) {
	t.Run("does not mutate using GetAttributes copy map", func(t *testing.T) {
		attributes := NewAttributes()
		value := NewString("test")
		attributes.SetAttribute(KeyLabel, value)

		indirectAttributes := attributes.GetAttributes()
		indirectAttributes[KeyClass] = NewString("my-class")

		got := attributes.GetAttributes()
		want := Map{
			KeyLabel: value,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	})
	t.Run("mutates using getAttributes map reference", func(t *testing.T) {
		attributes := NewAttributes()

		labelValue := NewString("test")
		attributes.SetAttribute(KeyLabel, labelValue)

		// mutate the map using the reference returned with the internal func
		indirectAttributes := attributes.getAttributes()
		classValue := NewString("my-class")
		indirectAttributes[KeyClass] = classValue

		// get a copy of the map using the public func
		got := attributes.GetAttributes()
		want := Map{
			KeyLabel: labelValue,
			KeyClass: classValue,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	})
	t.Run("get single attribute as string", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttribute(KeyLabel, NewHTML("<b>html label</b>"))

		got := attributes.GetAttributeString(KeyLabel)
		want := "<b>html label</b>"

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	})
}

func TestAttributes_Write(t *testing.T) {
	t.Run("writes single string attribute without brackets", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttribute(KeyLabel, NewString("test"))

		var gotStringBuilder strings.Builder
		attributes.Write(&gotStringBuilder, false)
		gotString := gotStringBuilder.String()
		expectedString := `label="test";`

		if !reflect.DeepEqual(gotString, expectedString) {
			t.Errorf("got [%v] want [%v]", gotString, expectedString)
		}
	})
	t.Run("writes single string attribute with brackets", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttribute(KeyLabel, NewString("test"))

		var gotStringBuilder strings.Builder
		attributes.Write(&gotStringBuilder, true)
		gotString := gotStringBuilder.String()
		expectedString := `[label="test"]`

		if !reflect.DeepEqual(gotString, expectedString) {
			t.Errorf("got [%v] want [%v]", gotString, expectedString)
		}
	})
	t.Run("writes single HTML attribute without brackets", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttributeHTML(KeyLabel, "<B>Hi</B>")

		var gotStringBuilder strings.Builder
		attributes.Write(&gotStringBuilder, false)
		gotString := gotStringBuilder.String()

		expectedString := "label=<<B>Hi</B>>;"

		if !reflect.DeepEqual(gotString, expectedString) {
			t.Errorf("got [%v] want [%v]", gotString, expectedString)
		}
	})
	t.Run("writes single HTML attribute with brackets", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttributeHTML(KeyLabel, "<B>Hi</B>")

		var gotStringBuilder strings.Builder
		attributes.Write(&gotStringBuilder, true)
		gotString := gotStringBuilder.String()

		expectedString := "[label=<<B>Hi</B>>]"

		if !reflect.DeepEqual(gotString, expectedString) {
			t.Errorf("got [%v] want [%v]", gotString, expectedString)
		}
	})
	t.Run("writes single Literal attribute without brackets", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttributeLiteral(KeyLabel, `"left text\l"`)

		var gotStringBuilder strings.Builder
		attributes.Write(&gotStringBuilder, false)
		gotString := gotStringBuilder.String()

		expectedString := `label="left text\l";`

		if !reflect.DeepEqual(gotString, expectedString) {
			t.Errorf("got [%v] want [%v]", gotString, expectedString)
		}
	})
	t.Run("writes single Literal attribute with brackets", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttributeLiteral(KeyLabel, `"left text\l"`)

		var gotStringBuilder strings.Builder
		attributes.Write(&gotStringBuilder, true)
		gotString := gotStringBuilder.String()

		expectedString := `[label="left text\l"]`

		if !reflect.DeepEqual(gotString, expectedString) {
			t.Errorf("got [%v] want [%v]", gotString, expectedString)
		}
	})
	t.Run("writes multi attributes without brackets", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttributesString(MapString{
			KeyClass: "my-class",
			KeyLabel: "my-label",
		})

		var gotStringBuilder strings.Builder
		attributes.Write(&gotStringBuilder, false)
		gotString := gotStringBuilder.String()

		expectedString := `class="my-class";label="my-label";`

		if !reflect.DeepEqual(gotString, expectedString) {
			t.Errorf("got [%v] want [%v]", gotString, expectedString)
		}
	})
	t.Run("writes multi attributes with brackets", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttributesString(MapString{
			KeyClass: "my-class",
			KeyLabel: "my-label",
		})

		var gotStringBuilder strings.Builder
		attributes.Write(&gotStringBuilder, true)
		gotString := gotStringBuilder.String()

		expectedString := `[class="my-class",label="my-label"]`

		if !reflect.DeepEqual(gotString, expectedString) {
			t.Errorf("got [%v] want [%v]", gotString, expectedString)
		}
	})
}

func TestAttributes_SetAttribute(t *testing.T) {
	t.Run("set attribute using single attribute set methods", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttributeString(KeyClass, "my-class")
		attributes.SetAttributeHTML(KeyLabel, "<b>my-label</b>")
		attributes.SetAttributeLiteral(KeyXlabel, `"left text\l"`)
		attributes.SetAttribute(KeyColor, NewString("black"))

		expectedString := `class="my-class";color="black";label=<<b>my-label</b>>;xlabel="left text\l";`
		var gotStringBuilder strings.Builder
		attributes.Write(&gotStringBuilder, false)
		gotString := gotStringBuilder.String()

		if !reflect.DeepEqual(gotString, expectedString) {
			t.Errorf("got [%v] want [%v]", gotString, expectedString)
		}
	})
	t.Run("set attribute using multi attribute set methods", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttributesString(MapString{
			KeyClass: "my-class",
		})
		attributes.SetAttributesHTML(MapString{
			KeyLabel: "<b>my-label</b>",
		})
		attributes.SetAttributesLiteral(MapString{
			KeyXlabel: `"left text\l"`,
		})
		attributes.SetAttributes(Map{
			KeyColor: NewString("black"),
		})

		expectedString := `class="my-class";color="black";label=<<b>my-label</b>>;xlabel="left text\l";`
		var gotStringBuilder strings.Builder
		attributes.Write(&gotStringBuilder, false)
		gotString := gotStringBuilder.String()

		if !reflect.DeepEqual(gotString, expectedString) {
			t.Errorf("got [%v] want [%v]", gotString, expectedString)
		}
	})
}

func TestAttributes_DeleteAttribute(t *testing.T) {
	t.Run("try to delete un-existant attribute", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.DeleteAttribute(KeyClass)

		gotMap := attributes.GetAttributes()
		expectedMap := Map{}

		if !reflect.DeepEqual(gotMap, expectedMap) {
			t.Errorf("got [%v] want [%v]", gotMap, expectedMap)
		}
	})
	t.Run("delete a set attribute", func(t *testing.T) {
		attributes := NewAttributes()
		attributes.SetAttribute(KeyLabel, NewString("test"))
		attributes.DeleteAttribute(KeyLabel)

		gotMap := attributes.GetAttributes()
		expectedMap := Map{}

		if !reflect.DeepEqual(gotMap, expectedMap) {
			t.Errorf("got [%v] want [%v]", gotMap, expectedMap)
		}
	})
}
