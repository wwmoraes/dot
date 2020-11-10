package dot

import "github.com/emicklei/dot/attributes"

// Node is implemented by dot-compatible node values
type Node interface {
	attributes.Object
	attributes.Serializable
	// Graph returns this node's graph
	Graph() Graph
	// Edge creates an Edge to a Node
	Edge(to Node) Edge
	// EdgeWithAttributes creates an Edge with the provided attributes to a Node
	EdgeWithAttributes(to Node, attributes attributes.Reader) Edge
	// EdgesTo returns all edges between this Node and the target Node
	EdgesTo(to Node) []Edge
}
