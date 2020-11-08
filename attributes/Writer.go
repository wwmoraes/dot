package attributes

import (
	"fmt"
)

// Writer graph object attribute read-write access
type Writer interface {
	Reader
	SetAttribute(key Key, value fmt.Stringer)
	SetAttributes(attributeMap Map)
	DeleteAttribute(key Key)
}
