package formatters

import (
	"bytes"
	"errors"
	"io"
	"regexp"
	"strings"
)

var ErrNoWriter = errors.New("the underlying writer is not valid")

type regexpRule struct {
	regex              *regexp.Regexp
	replacement        []byte
	indent             bool
	adjustIndentBefore int
	adjustIndentAfter  int
}

func (thisRule *regexpRule) Apply(data []byte, level int) ([]byte, int) {
	// skip entirely if the regex doesn't match
	if !thisRule.regex.Match(data) {
		return data, level
	}

	var newData bytes.Buffer

	// indent with pre/post-adjust
	level = level + thisRule.adjustIndentBefore
	if thisRule.indent {
		newData.WriteString(strings.Repeat("  ", level))
	}
	level = level + thisRule.adjustIndentAfter

	newData.Write(thisRule.regex.ReplaceAllLiteral(data, thisRule.replacement))

	return newData.Bytes(), level
}

var masterRules []regexpRule = []regexpRule{
	{
		// break lines and increase level on opening blocks
		regex:             regexp.MustCompile(`{`),
		replacement:       []byte("{\n"),
		adjustIndentAfter: 1,
	},
	{
		// indent and increase level of anonymous blocks
		regex:              regexp.MustCompile(`^{`),
		replacement:        []byte("{"),
		indent:             true,
		adjustIndentBefore: -1,
		adjustIndentAfter:  1,
	},
	{
		// break lines after semicolon
		regex:       regexp.MustCompile(`;`),
		replacement: []byte(";\n"),
	},
	{
		// break lines and decrease indent on closing blocks
		regex:              regexp.MustCompile(`}`),
		replacement:        []byte("}\n"),
		indent:             true,
		adjustIndentBefore: -1,
	},
	{
		// indent global graph attributes
		regex:       regexp.MustCompile(`\bgraph\b`),
		replacement: []byte(`graph`),
		indent:      true,
	},
	{
		// indent subgraphs
		regex:       regexp.MustCompile(`\bsubgraph\b`),
		replacement: []byte(`subgraph`),
		indent:      true,
	},
	{
		// indent global node attributes
		regex:       regexp.MustCompile(`\bnode\b`),
		replacement: []byte(`node`),
		indent:      true,
	},
	{
		// indent global edge attributes
		regex:       regexp.MustCompile(`\bedge\b`),
		replacement: []byte(`edge`),
		indent:      true,
	},
	{
		// indent nodes and edges
		regex:       regexp.MustCompile(`^"`),
		replacement: []byte(`"`),
		indent:      true,
	},
}

type prettyWriter struct {
	writer io.Writer
	rules  []regexpRule
	level  int
}

// NewPrettyWriter creates a new instance of a prettifier formatter for dot code
func NewPrettyWriter(writer io.Writer) io.Writer {
	return &prettyWriter{
		writer: writer,
		rules:  masterRules,
	}
}

// Write writes data bytes to the underlying writer
func (thisWriter *prettyWriter) Write(data []byte) (n int, err error) {
	if thisWriter.writer == nil {
		return 0, ErrNoWriter
	}

	for _, rule := range thisWriter.rules {
		data, thisWriter.level = rule.Apply(data, thisWriter.level)
	}
	return thisWriter.writer.Write(data)
}
