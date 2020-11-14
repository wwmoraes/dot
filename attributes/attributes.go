package attributes

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/wwmoraes/dot/constants"
)

// Map collection of dot component attributes
type Map map[constants.Key]fmt.Stringer

// MapString collection of dot component attributes as primitive strings
type MapString map[constants.Key]string

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

// GetAttribute returns the attribute value or nil if unset
func (dotObjectData *Attributes) GetAttribute(key constants.Key) (fmt.Stringer, bool) {
	attr, found := dotObjectData.attributes[key]
	return attr, found
}

// GetAttributeString returns the attribute as string or an empty string if unset
func (dotObjectData *Attributes) GetAttributeString(key constants.Key) string {
	attr, ok := dotObjectData.GetAttribute(key)

	if !ok {
		return ""
	}

	return attr.String()
}

// GetAttributes returns a copy of all attributes
func (dotObjectData *Attributes) GetAttributes() Map {
	newMap := make(Map, len(dotObjectData.attributes))
	for key, value := range dotObjectData.attributes {
		newMap[key] = value
	}
	return newMap
}

// SetAttribute sets the value for the attribute Key
func (dotObjectData *Attributes) SetAttribute(key constants.Key, value fmt.Stringer) {
	dotObjectData.attributes[key] = value
}

// SetAttributeString sets the string value for the attribute Key
func (dotObjectData *Attributes) SetAttributeString(key constants.Key, value string) {
	dotObjectData.attributes[key] = NewString(value)
}

// SetAttributeLiteral sets the literal value for the attribute Key
func (dotObjectData *Attributes) SetAttributeLiteral(key constants.Key, value string) {
	dotObjectData.attributes[key] = NewLiteral(value)
}

// SetAttributeHTML sets the HTML value for the attribute Key
func (dotObjectData *Attributes) SetAttributeHTML(key constants.Key, value string) {
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
func (dotObjectData *Attributes) DeleteAttribute(key constants.Key) {
	delete(dotObjectData.attributes, key)
}

// WriteTo writes formatted attributes data to the writer
func (dotObjectData *Attributes) WriteTo(writer io.Writer) (n int64, err error) {
	if !dotObjectData.HasAttributes() {
		return 0, nil
	}

	// first collect keys
	keys := []constants.Key{}
	for k := range dotObjectData.attributes {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return strings.Compare(string(keys[i]), string(keys[j])) < 0
	})

	// write keys
	var stringAttributes = make([]string, len(keys))
	for i, k := range keys {
		switch attributeData := dotObjectData.attributes[k].(type) {
		case *HTML:
			stringAttributes[i] = fmt.Sprintf("%s=<%s>", k, attributeData.String())
		case *Literal:
			stringAttributes[i] = fmt.Sprintf("%s=%s", k, attributeData.String())
		default:
			stringAttributes[i] = fmt.Sprintf("%s=%q", k, attributeData.String())
		}
	}
	written32, err := fmt.Fprintf(writer, "[%s]", strings.Join(stringAttributes, ","))
	n += int64(written32)
	return
}

// HasAttributes returns true if there's any attribute set
func (dotObjectData *Attributes) HasAttributes() bool {
	return len(dotObjectData.attributes) > 0
}
