package attributes

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

// Map collection of dot component attributes
type Map map[Key]fmt.Stringer

// MapString collection of dot component attributes as primitive strings
type MapString map[Key]string

// Attributes dot component attributes data
type Attributes struct {
	attributes Map
}

// NewAttributes creates an empty attributes map
func NewAttributes() *Attributes {
	return &Attributes{
		attributes: make(Map),
	}
}

// NewAttributesFrom creates an attributes map with the provided attributes
func NewAttributesFrom(attr Reader) *Attributes {
	if attr == nil {
		return NewAttributes()
	}

	return &Attributes{
		attributes: attr.GetAttributes(),
	}
}

// getAttributes returns a reference to the internal attributes map
func (dotObjectData *Attributes) getAttributes() Map {
	return dotObjectData.attributes
}

func (dotObjectData *Attributes) GetAttribute(key Key) fmt.Stringer {
	return dotObjectData.attributes[key]
// GetAttribute returns the attribute value or nil if unset
}

// GetAttributeString returns the attribute as string or an empty string if unset
func (dotObjectData *Attributes) GetAttributeString(key Key) string {
	return dotObjectData.attributes[key].String()
}

// GetAttributes returns a copy of all attributes
func (dotObjectData *Attributes) GetAttributes() Map {
	newMap := make(Map, len(dotObjectData.attributes))
	for key, value := range dotObjectData.attributes {
		newMap[key] = value
	}
	return newMap
}

// Write transforms attributes into dot notation and writes on the given writer
func (dotObjectData *Attributes) Write(device io.Writer, mustBracket bool) {
	if len(dotObjectData.attributes) == 0 {
		return
	}

	if mustBracket {
		fmt.Fprint(device, "[")
	}
	first := true
	// first collect keys
	keys := []Key{}
	for k := range dotObjectData.attributes {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return strings.Compare(string(keys[i]), string(keys[j])) < 0
	})

	for _, k := range keys {
		if !first {
			if mustBracket {
				fmt.Fprint(device, ",")
			} else {
				fmt.Fprint(device, ";")
			}
		}
		switch attributeData := dotObjectData.attributes[k].(type) {
		case *HTML:
			fmt.Fprintf(device, "%s=<%s>", k, attributeData.value)
		case *Literal:
			fmt.Fprintf(device, "%s=%s", k, attributeData.value)
		default:
			fmt.Fprintf(device, "%s=%q", k, attributeData.String())
		}
		first = false
	}
	if mustBracket {
		fmt.Fprint(device, "]")
	} else {
		fmt.Fprint(device, ";")
	}
}

// SetAttribute sets the value for the attribute Key
func (dotObjectData *Attributes) SetAttribute(key Key, value fmt.Stringer) {
	dotObjectData.attributes[key] = value
}

// SetAttributeString sets the string value for the attribute Key
func (dotObjectData *Attributes) SetAttributeString(key Key, value string) {
	dotObjectData.attributes[key] = NewString(value)
}

// SetAttributeLiteral sets the literal value for the attribute Key
func (dotObjectData *Attributes) SetAttributeLiteral(key Key, value string) {
	dotObjectData.attributes[key] = NewLiteral(value)
}

// SetAttributeHTML sets the HTML value for the attribute Key
func (dotObjectData *Attributes) SetAttributeHTML(key Key, value string) {
	dotObjectData.attributes[key] = NewHTML(value)
}

// SetAttributes sets the value for multiple attributes
func (dotObjectData *Attributes) SetAttributes(attributeMap Map) {
	for k, v := range attributeMap {
		dotObjectData.attributes[k] = v
	}
}

// SetAttributesString sets the string value for multiple attributes
func (dotObjectData *Attributes) SetAttributesString(attributeMap MapString) {
	for k, v := range attributeMap {
		dotObjectData.attributes[k] = NewString(v)
	}
}

// SetAttributesLiteral sets the literal value for multiple attributes
func (dotObjectData *Attributes) SetAttributesLiteral(attributeMap MapString) {
	for k, v := range attributeMap {
		dotObjectData.attributes[k] = NewLiteral(v)
	}
}

// SetAttributesHTML sets the HTML value for multiple attributes
func (dotObjectData *Attributes) SetAttributesHTML(attributeMap MapString) {
	for k, v := range attributeMap {
		dotObjectData.attributes[k] = NewHTML(v)
	}
}

// DeleteAttribute unset the attribute Key
func (dotObjectData *Attributes) DeleteAttribute(key Key) {
	delete(dotObjectData.attributes, key)
}
