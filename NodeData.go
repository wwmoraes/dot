package dot

import "github.com/emicklei/dot/attributes"

// NodeData represents a dot NodeData.
type NodeData struct {
	*attributes.Attributes
	graph Graph
	id    string
}

// ID returns the immutable id
func (n *NodeData) ID() string {
	return n.id
}

func (n *NodeData) String() string {
	// TODO
	return n.id
}

func (n *NodeData) Graph() Graph {
	return n.graph
}

// Label sets the attribute "label" to the given label
func (n *NodeData) Label(label string) Node {
	n.SetAttribute(attributes.KeyLabel, attributes.NewString(label))
	return n
}

// Box sets the attribute "shape" to "box"
func (n *NodeData) Box() Node {
	n.SetAttribute(attributes.KeyShape, *attributes.ShapeBox)
	return n
}

// Edge creates a Edge and returns it for chaining
func (n *NodeData) Edge(toNode Node) Edge {
	return n.graph.EdgeWithAttributes(n, toNode, nil)
}

// EdgeWithAttributes sets the given attributes and returns the Edge for chaining
func (n *NodeData) EdgeWithAttributes(toNode Node, attributes attributes.Reader) Edge {
	return n.graph.EdgeWithAttributes(n, toNode, attributes)
}

// EdgesTo returns all existing edges between this Node and the argument Node.
func (n *NodeData) EdgesTo(toNode Node) []Edge {
	return n.graph.FindEdges(n, toNode)
}
