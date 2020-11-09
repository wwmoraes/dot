package attributes

import (
	"fmt"
)

// Writer graph object attribute read-write access
type Writer interface {
	Reader
	SetAttribute(key Key, value fmt.Stringer)
	SetAttributeString(key Key, value string)
	SetAttributeLiteral(key Key, value string)
	SetAttributeHTML(key Key, value string)
	SetAttributes(attributeMap Map)
	SetAttributesString(attributeMap MapString)
	SetAttributesLiteral(attributeMap MapString)
	SetAttributesHTML(attributeMap MapString)
	DeleteAttribute(key Key)
}
