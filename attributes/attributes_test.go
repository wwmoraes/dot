package attributes

import (
	"reflect"
	"strings"
	"testing"
)

func TestAttributes_empty(t *testing.T) {
	attributes := NewAttributes()
	expectedAttributes := Map{}
	if !reflect.DeepEqual(attributes.attributes, expectedAttributes) {
		t.Errorf("got [%v] want [%v]", attributes, expectedAttributes)
	}

	expectedString := ""
	var gotStringBuilder strings.Builder
	attributes.Write(&gotStringBuilder, false)
	gotString := gotStringBuilder.String()

	if !reflect.DeepEqual(gotString, expectedString) {
		t.Errorf("got [%v] want [%v]", gotString, expectedString)
	}
}

func TestAttributes_withValue(t *testing.T) {
	attributes := NewAttributes()
	expectedLabelValue := NewString("test")
	expectedMap := Map{
		AttributeLabel: expectedLabelValue,
	}

	attributes.SetAttribute(AttributeLabel, NewString("test"))
	if !reflect.DeepEqual(attributes.attributes, expectedMap) {
		t.Errorf("got [%v] want [%v]", attributes, expectedMap)
	}

	// test marshalling a single attribute map without brackets
	{
		expectedString := `label="test";`
		var gotStringBuilder strings.Builder
		attributes.Write(&gotStringBuilder, false)
		gotString := gotStringBuilder.String()

		if !reflect.DeepEqual(gotString, expectedString) {
			t.Errorf("got [%v] want [%v]", gotString, expectedString)
		}
	}

	// test marshalling a single attribute map with brackets
	{
		expectedString := `[label="test"]`
		var gotStringBuilder strings.Builder
		attributes.Write(&gotStringBuilder, true)
		gotString := gotStringBuilder.String()

		if !reflect.DeepEqual(gotString, expectedString) {
			t.Errorf("got [%v] want [%v]", gotString, expectedString)
		}
	}

	// test getting a single value
	{
		gotValue := attributes.GetAttribute(AttributeLabel)
		if !reflect.DeepEqual(gotValue, expectedLabelValue) {
			t.Errorf("got [%v] want [%v]", gotValue, expectedLabelValue)
		}
	}

	// test removing an unset value
	{
		attributes.DeleteAttribute(AttributeClass)
		gotMap := attributes.GetAttributes()

		if !reflect.DeepEqual(gotMap, expectedMap) {
			t.Errorf("got [%v] want [%v]", gotMap, expectedMap)
		}
	}

	// test removing a set value
	{
		attributes.DeleteAttribute(AttributeLabel)
		gotMap := attributes.GetAttributes()
		expectedMap = Map{}

		if !reflect.DeepEqual(gotMap, expectedMap) {
			t.Errorf("got [%v] want [%v]", gotMap, expectedMap)
		}
	}

	// test marshalling HTML label
	{
		expectedValue := NewHTML("<B>Hi</B>")
		attributes.SetAttribute(AttributeLabel, expectedValue)

		expectedString := `label=<<B>Hi</B>>;`
		var gotStringBuilder strings.Builder
		attributes.Write(&gotStringBuilder, false)
		gotString := gotStringBuilder.String()

		if !reflect.DeepEqual(gotString, expectedString) {
			t.Errorf("got [%v] want [%v]", gotString, expectedString)
		}
	}

	// test literal label
	{
		expectedValue := NewLiteral(`"left-justified text\l"`)
		attributes.SetAttribute(AttributeLabel, expectedValue)

		expectedString := `label="left-justified text\l";`
		var gotStringBuilder strings.Builder
		attributes.Write(&gotStringBuilder, false)
		gotString := gotStringBuilder.String()

		if !reflect.DeepEqual(gotString, expectedString) {
			t.Errorf("got [%v] want [%v]", gotString, expectedString)
		}
	}

	// test setting multiple attributes
	{
		expectedMap := Map{
			AttributeClass: NewString("my-class"),
			AttributeLabel: NewString("my-label"),
		}
		attributes.SetAttributes(expectedMap)
		gotMap := attributes.GetAttributes()

		if !reflect.DeepEqual(gotMap, expectedMap) {
			t.Errorf("got [%v] want [%v]", gotMap, expectedMap)
		}
	}

	// test marshalling a multi attribute map without brackets
	{
		expectedString := `class="my-class";label="my-label";`
		var gotStringBuilder strings.Builder
		attributes.Write(&gotStringBuilder, false)
		gotString := gotStringBuilder.String()

		if !reflect.DeepEqual(gotString, expectedString) {
			t.Errorf("got [%v] want [%v]", gotString, expectedString)
		}
	}
	// test marshalling a multi attribute map with brackets
	{
		expectedString := `[class="my-class",label="my-label"]`
		var gotStringBuilder strings.Builder
		attributes.Write(&gotStringBuilder, true)
		gotString := gotStringBuilder.String()

		if !reflect.DeepEqual(gotString, expectedString) {
			t.Errorf("got [%v] want [%v]", gotString, expectedString)
		}
	}
}

func TestHTML(t *testing.T) {
	expectedValue := "<B>Hi</B>"
	htmlString := NewHTML(expectedValue)
	gotValue := htmlString.String()

	if !reflect.DeepEqual(gotValue, expectedValue) {
		t.Errorf("got [%v] want [%v]", gotValue, expectedValue)
	}
}

func TestLiteral(t *testing.T) {
	expectedValue := `"left-justified text\l"`
	literalString := NewLiteral(expectedValue)
	gotValue := literalString.String()

	if !reflect.DeepEqual(gotValue, expectedValue) {
		t.Errorf("got [%v] want [%v]", gotValue, expectedValue)
	}
}
