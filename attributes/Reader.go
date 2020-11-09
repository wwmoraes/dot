package attributes

import (
	"fmt"
)

// Reader graph object attribute read-only access
type Reader interface {
	GetAttribute(key Key) fmt.Stringer
	GetAttributeString(key Key) string
	GetAttributes() Map
}
