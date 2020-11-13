package attributes

import (
	"fmt"
)

// Reader is implemented by attribute-based values that allows reading them
type Reader interface {
	// GetAttribute returns the attribute value or nil if unset
	GetAttribute(key Key) (fmt.Stringer, bool)
	// GetAttributeString returns the attribute as string or an empty string if unset
	GetAttributeString(key Key) string
	// GetAttributes returns a copy of all attributes
	GetAttributes() Map
	// HasAttributes returns true if there's any attribute set
	HasAttributes() bool
}
