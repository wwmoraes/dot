package attributes

import (
	"fmt"
	"io"
)

type Serializable interface {
	fmt.Stringer
	Write(device io.Writer, mustBracket bool)
}
