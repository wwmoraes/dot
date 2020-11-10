package dot

import "github.com/wwmoraes/dot/attributes"

// Edge is implemented by dot-compatible edge values
type Edge interface {
	attributes.Styleable
	attributes.Serializable
	// From returns the tail node this Edge is connected from
	From() Node
	// From returns the head node this Edge is connected to
	To() Node
	// Edge creates an Edge to a Node using the head node of this Edge as tail
	Edge(to Node) Edge
	// Edge creates an Edge with the provided attributes to a Node using the head
	// node of this Edge as tail
	EdgeWithAttributes(to Node, attributes attributes.Reader) Edge
	// EdgesTo returns all edges between the head Node of this Edge and the target
	// Node
	EdgesTo(to Node) []Edge
}

// StyledEdge is implemented by dot-compatible edge values which have
// convenience styling methods
type StyledEdge interface {
	Edge
	// Solid sets the edge style to solid
	Solid() Edge
	// Solid sets the edge style to bold
	Bold() Edge
	// Solid sets the edge style to dashed
	Dashed() Edge
	// Solid sets the edge style to dotted
	Dotted() Edge
}
