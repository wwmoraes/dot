package attributes

import (
	"fmt"
)

// Reader is implemented by attribute-based values that allows reading them
type Reader interface {
	GetAttribute(key Key) fmt.Stringer
	// GetAttribute returns the attribute value or nil if unset
	// GetAttributeString returns the attribute as string or an empty string if unset
	GetAttributeString(key Key) string
	// GetAttributes returns a copy of all attributes
	GetAttributes() Map
}
