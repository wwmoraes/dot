package attributes

import (
	"fmt"
	"io"
)

// Serializable is implemented by values that can be printed as string or
// written directly into an IO device
type Serializable interface {
	fmt.Stringer
	Write(device io.Writer, mustBracket bool)
}
