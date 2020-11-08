package attributes

// Literal renders the content as is, without any quoting or escaping.
// The caller is fully responsible for properly quoting and escaping the content.
// This is useful for some control mechanisms of Graphviz, as \l to left-align
// labels, e.g. Literal(`"my left-aligned text\l"`)
type Literal struct {
	value string
}

// NewLiteral creates a literal string that'll be written as is
func NewLiteral(value string) *Literal {
	return &Literal{value}
}

// String returns the literal string
func (data Literal) String() string {
	return data.value
}
