package formatters

import (
	"errors"
	"io"
	"regexp"
	"strings"
)

type regexpRule struct {
	regex              *regexp.Regexp
	replacement        string
	expand             bool
	indent             bool
	adjustIndentBefore int
	adjustIndentAfter  int
}

func (thisRule *regexpRule) Apply(dataString string, level int) (string, int) {
	// skip entirely if the regex doesn't match
	if !thisRule.regex.MatchString(dataString) {
		return dataString, level
	}

	var newDataString strings.Builder

	// indent with pre/post-adjust
	level = level + thisRule.adjustIndentBefore
	if thisRule.indent {
		newDataString.WriteString(strings.Repeat("  ", level))
	}
	level = level + thisRule.adjustIndentAfter

	if thisRule.expand {
		newDataString.WriteString(thisRule.regex.ReplaceAllString(dataString, thisRule.replacement))
	} else {
		newDataString.WriteString(thisRule.regex.ReplaceAllLiteralString(dataString, thisRule.replacement))
	}

	return newDataString.String(), level
}

var masterRules []regexpRule = []regexpRule{
	{
		// break lines and increase level on opening blocks
		regex:             regexp.MustCompile(`{`),
		replacement:       "{\n",
		adjustIndentAfter: 1,
	},
	{
		// indent and increase level of anonymous blocks
		regex:              regexp.MustCompile(`^{`),
		replacement:        "{",
		indent:             true,
		adjustIndentBefore: -1,
		adjustIndentAfter:  1,
	},
	{
		// break lines after semicolon
		regex:       regexp.MustCompile(`;`),
		replacement: ";\n",
	},
	{
		// break lines and decrease indent on closing blocks
		regex:              regexp.MustCompile(`}`),
		replacement:        "}\n",
		indent:             true,
		adjustIndentBefore: -1,
	},
	{
		// indent global graph attributes
		regex:       regexp.MustCompile(`\bgraph\b`),
		replacement: `graph`,
		indent:      true,
	},
	{
		// indent subgraphs
		regex:       regexp.MustCompile(`\bsubgraph\b`),
		replacement: `subgraph`,
		indent:      true,
	},
	{
		// indent global node attributes
		regex:       regexp.MustCompile(`\bnode\b`),
		replacement: `node`,
		indent:      true,
	},
	{
		// indent global edge attributes
		regex:       regexp.MustCompile(`\bedge\b`),
		replacement: `edge`,
		indent:      true,
	},
	{
		// indent nodes and edges
		regex:       regexp.MustCompile(`^"`),
		replacement: `"`,
		indent:      true,
	},
	{
		// indent attributes directly set on the graph
		regex:       regexp.MustCompile(`^([a-zA-Z0-9 ]+=)`),
		replacement: `$1`,
		indent:      true,
		expand:      true,
	},
	// {
	// 	regex:        regexp.MustCompile(`(.*)`),
	// 	replacement:  "^$1$",
	// 	indentBefore: false,
	// },
}

type prettyWriter struct {
	byteWriter   io.Writer
	stringWriter io.StringWriter
	rules        []regexpRule
	level        int
}

// NewPrettyWriter creates a new instance of a prettifier formatter for dot code
func NewPrettyWriter(writer io.Writer) io.Writer {
	if stringWriter, ok := writer.(io.StringWriter); ok {
		return &prettyWriter{
			stringWriter: stringWriter,
			rules:        masterRules,
		}
	}

	return &prettyWriter{
		byteWriter: writer,
		rules:      masterRules,
	}
}

// Write writes data bytes to the underlying writer
func (thisWriter *prettyWriter) Write(data []byte) (n int, err error) {
	return thisWriter.processWrite(string(data))
}

// WriteString writes the string data to the underlying writer
func (thisWriter *prettyWriter) WriteString(stringData string) (n int, err error) {
	return thisWriter.processWrite(stringData)
}

// processWrite prettifies the string data before forwarding it
func (thisWriter *prettyWriter) processWrite(stringData string) (n int, err error) {
	for _, rule := range thisWriter.rules {
		stringData, thisWriter.level = rule.Apply(stringData, thisWriter.level)
	}
	return thisWriter.forwardWrite(stringData)
}

// forwardWrite sends the modified string data to the underlying writer
func (thisWriter *prettyWriter) forwardWrite(stringData string) (n int, err error) {
	if thisWriter.stringWriter != nil {
		return thisWriter.stringWriter.WriteString(stringData)
	}

	if thisWriter.byteWriter != nil {
		return thisWriter.byteWriter.Write([]byte(stringData))
	}

	return 0, errors.New("Unimplemented")
}
