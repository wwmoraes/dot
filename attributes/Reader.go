package attributes

import (
	"fmt"
	"io"
)

// Reader graph object attribute read-only access
type Reader interface {
	GetAttribute(key Key) fmt.Stringer
	GetAttributes() Map
	Write(device io.Writer, mustBracket bool)
}
