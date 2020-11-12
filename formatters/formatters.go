package formatters

import "io"

// IndentedWriter is implemented by passthrough writers with indent helpers
type IndentedWriter interface {
	io.Writer
	io.StringWriter
	Indent()
	BackIndent()
	IndentWhile(block func())
	NewLineIndentWhile(block func())
	NewLine()
}
