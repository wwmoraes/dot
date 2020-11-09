package dot

import "github.com/emicklei/dot/attributes"

// Node represents a graph node
type Node interface {
	attributes.Object
	attributes.Serializable
	Graph() Graph
	Edge(to Node) Edge
	EdgeWithAttributes(to Node, attributes attributes.Reader) Edge
	EdgesTo(to Node) []Edge
}
