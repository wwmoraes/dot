package dot

import "github.com/wwmoraes/dot/attributes"

// Node is implemented by dot-compatible node values
type Node interface {
	attributes.Identity
	attributes.Styleable
	attributes.Serializable
	// Edge creates an Edge to a Node
	Edge(to Node) Edge
	// EdgeWithAttributes creates an Edge with the provided attributes to a Node
	EdgeWithAttributes(to Node, attributes attributes.Reader) Edge
	// EdgesTo returns all edges between this Node and the target Node
	EdgesTo(to Node) []Edge
}

// StyledNode is implemented by dot-compatible node values which have
// convenience styling methods
type StyledNode interface {
	// Box sets the node style to box
	Box() Node
}
