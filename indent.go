package dot

import (
	"fmt"
	"io"
)

// IndentWriter io.Writer that indents content with the set level
type IndentWriter struct {
	level  int
	writer io.Writer
}

// NewIndentWriter new instance of IndentWriter
func NewIndentWriter(w io.Writer) *IndentWriter {
	return &IndentWriter{level: 0, writer: w}
}

// Indent increase the indent level and write a tab
func (i *IndentWriter) Indent() {
	i.level++
	fmt.Fprint(i.writer, "\t")
}

// BackIndent decrease the indent level
func (i *IndentWriter) BackIndent() {
	i.level--
}

// IndentWhile executes a block on an indented section
func (i *IndentWriter) IndentWhile(block func()) {
	i.Indent()
	block()
	i.BackIndent()
}

// NewLineIndentWhile executes a block on a indented and newline separared section
func (i *IndentWriter) NewLineIndentWhile(block func()) {
	i.NewLine()
	i.IndentWhile(block)
	i.NewLine()
}

// NewLine adds a line break and indent the new line
func (i *IndentWriter) NewLine() {
	fmt.Fprint(i.writer, "\n")
	for j := 0; j < i.level; j++ {
		fmt.Fprint(i.writer, "\t")
	}
}

// Write makes it an io.Writer
func (i *IndentWriter) Write(data []byte) (n int, err error) {
	return i.writer.Write(data)
}

// WriteString writes the given string
func (i *IndentWriter) WriteString(s string) (n int, err error) {
	fmt.Fprint(i.writer, s)
	return len(s), nil
}
