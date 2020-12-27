package dot

// StyledNode is implemented by dot-compatible node values which have
// convenience styling methods
type StyledNode interface {
	Node
	// Box sets the node style to box
	Box() Node
}
