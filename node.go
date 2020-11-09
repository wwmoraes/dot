package dot

import "github.com/emicklei/dot/attributes"

// Node represents a dot Node.
type Node struct {
	*attributes.Attributes
	graph *Graph
	id    string
}

// ID returns the node immutable id
func (n *Node) ID() string {
	return n.id
}

// Label sets the attribute "label" to the given label
func (n *Node) Label(label string) *Node {
	n.SetAttribute(attributes.KeyLabel, attributes.NewString(label))
	return n
}

// Box sets the attribute "shape" to "box"
func (n *Node) Box() *Node {
	n.SetAttribute(attributes.KeyShape, *attributes.ShapeBox)
	return n
}

// Edge sets label=value and returns the Edge for chaining.
func (n *Node) Edge(toNode *Node, labels ...string) *Edge {
	return n.graph.Edge(n, toNode, labels...)
}

// EdgesTo returns all existing edges between this Node and the argument Node.
func (n *Node) EdgesTo(toNode *Node) []*Edge {
	return n.graph.FindEdges(n, toNode)
}
