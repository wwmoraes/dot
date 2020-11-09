package dot

import "github.com/emicklei/dot/attributes"

type Node interface {
	attributes.Object
	attributes.Serializable
	Graph() Graph
	Edge(to Node) Edge
	EdgeWithAttributes(to Node, attributes attributes.Reader) Edge
	EdgesTo(to Node) []Edge
}
