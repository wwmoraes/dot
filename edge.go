package dot

import "github.com/emicklei/dot/attributes"

type Edge interface {
	attributes.Object
	attributes.Serializable
	From() Node
	To() Node
	Edge(to Node) Edge
	EdgeWithAttributes(to Node, attributes attributes.Reader) Edge
	EdgesTo(to Node) []Edge
	// TODO remove those
	Solid() Edge
	Bold() Edge
	Dashed() Edge
	Dotted() Edge
}
