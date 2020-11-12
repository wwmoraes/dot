package formatters

import (
	"fmt"
	"io"
)

// indentedWriter io.Writer that indents content with the set level
type indentedWriter struct {
	level  int
	writer io.Writer
}

// NewIndentedWriter new instance of IndentWriter
func NewIndentedWriter(w io.Writer) IndentedWriter {
	return &indentedWriter{level: 0, writer: w}
}

// Indent increase the indent level and write a tab
func (i *indentedWriter) Indent() {
	i.level++
	fmt.Fprint(i.writer, "\t")
}

// BackIndent decrease the indent level
func (i *indentedWriter) BackIndent() {
	i.level--
}

// IndentWhile executes a block on an indented section
func (i *indentedWriter) IndentWhile(block func()) {
	i.Indent()
	block()
	i.BackIndent()
}

// NewLineIndentWhile executes a block on a indented and newline separared section
func (i *indentedWriter) NewLineIndentWhile(block func()) {
	i.NewLine()
	i.IndentWhile(block)
	i.NewLine()
}

// NewLine adds a line break and indent the new line
func (i *indentedWriter) NewLine() {
	fmt.Fprint(i.writer, "\n")
	for j := 0; j < i.level; j++ {
		fmt.Fprint(i.writer, "\t")
	}
}

// Write makes it an io.Writer
func (i *indentedWriter) Write(data []byte) (n int, err error) {
	return i.writer.Write(data)
}

// WriteString writes the given string
func (i *indentedWriter) WriteString(s string) (n int, err error) {
	fmt.Fprint(i.writer, s)
	return len(s), nil
}
