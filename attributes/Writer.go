package attributes

import (
	"fmt"

	"github.com/wwmoraes/dot/constants"
)

// Writer is implemented by attribute-based values that allows mutating them
type Writer interface {
	// SetAttribute sets the value for the attribute Key
	SetAttribute(key constants.Key, value fmt.Stringer)
	// SetAttributeString sets the string value for the attribute Key
	SetAttributeString(key constants.Key, value string)
	// SetAttributeLiteral sets the literal value for the attribute Key
	SetAttributeLiteral(key constants.Key, value string)
	// SetAttributeHTML sets the HTML value for the attribute Key
	SetAttributeHTML(key constants.Key, value string)
	// SetAttributes sets the value for multiple attributes
	SetAttributes(attributeMap Map)
	// SetAttributesString sets the string value for multiple attributes
	SetAttributesString(attributeMap MapString)
	// SetAttributesLiteral sets the literal value for multiple attributes
	SetAttributesLiteral(attributeMap MapString)
	// SetAttributesHTML sets the HTML value for multiple attributes
	SetAttributesHTML(attributeMap MapString)
	// DeleteAttribute unset the attribute Key
	DeleteAttribute(key constants.Key)
}
