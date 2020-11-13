package attributes

import (
	"io"
)

// Serializable is implemented by values that can be printed as string or
// written directly into an IO device
type Serializable interface {
	io.WriterTo
	String() (string, error)
}
