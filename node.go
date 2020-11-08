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

// Attr sets key=value and return the Node
func (n *Node) Attr(key attributes.Key, value string) *Node {
	n.SetAttribute(key, attributes.NewString(value))
	return n
}

// Label sets the attribute "label" to the given label
func (n *Node) Label(label string) *Node {
	return n.Attr("label", label)
}

// Box sets the attribute "shape" to "box"
func (n *Node) Box() *Node {
	return n.Attr("shape", "box")
}

// Edge sets label=value and returns the Edge for chaining.
func (n *Node) Edge(toNode *Node, labels ...string) *Edge {
	return n.graph.Edge(n, toNode, labels...)
}

// EdgesTo returns all existing edges between this Node and the argument Node.
func (n *Node) EdgesTo(toNode *Node) []*Edge {
	return n.graph.FindEdges(n, toNode)
}
