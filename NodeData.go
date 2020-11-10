package dot

import "github.com/wwmoraes/dot/attributes"

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
