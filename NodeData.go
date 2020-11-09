package dot

import "github.com/emicklei/dot/attributes"

// node represents a dot node.
type node struct {
	*attributes.Attributes
	graph Graph
	id    string
}

// ID returns the immutable id
func (thisNode *node) ID() string {
	return thisNode.id
}

func (thisNode *node) String() string {
	// TODO
	return thisNode.id
}

func (thisNode *node) Graph() Graph {
	return thisNode.graph
}

// Label sets the attribute "label" to the given label
func (thisNode *node) Label(label string) Node {
	thisNode.SetAttribute(attributes.KeyLabel, attributes.NewString(label))
	return thisNode
}

// Box sets the attribute "shape" to "box"
func (thisNode *node) Box() Node {
	thisNode.SetAttribute(attributes.KeyShape, *attributes.ShapeBox)
	return thisNode
}

// Edge creates a Edge and returns it for chaining
func (thisNode *node) Edge(toNode Node) Edge {
	return thisNode.graph.EdgeWithAttributes(thisNode, toNode, nil)
}

// EdgeWithAttributes sets the given attributes and returns the Edge for chaining
func (thisNode *node) EdgeWithAttributes(toNode Node, attributes attributes.Reader) Edge {
	return thisNode.graph.EdgeWithAttributes(thisNode, toNode, attributes)
}

// EdgesTo returns all existing edges between this Node and the argument Node.
func (thisNode *node) EdgesTo(toNode Node) []Edge {
	return thisNode.graph.FindEdges(thisNode, toNode)
}
