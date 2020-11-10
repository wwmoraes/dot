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

// Graph return this node's graph
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

// Edge creates an Edge to a node
func (thisNode *node) Edge(toNode Node) Edge {
	return thisNode.graph.EdgeWithAttributes(thisNode, toNode, nil)
}

// EdgeWithAttributes creates an Edge with the provided attributes to the a node
func (thisNode *node) EdgeWithAttributes(toNode Node, attributes attributes.Reader) Edge {
	return thisNode.graph.EdgeWithAttributes(thisNode, toNode, attributes)
}

// EdgesTo returns all edges between this Node and the target Node
func (thisNode *node) EdgesTo(toNode Node) []Edge {
	return thisNode.graph.FindEdges(thisNode, toNode)
}
