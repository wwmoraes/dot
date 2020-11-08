package dot

import (
	"fmt"

	"github.com/emicklei/dot/attributes"
)

// Edge represents a graph edge between two Nodes.
type Edge struct {
	*attributes.Attributes
	graph    *Graph
	from, to *Node
	id       string
}

// Label sets "label"=value and returns the Edge.
// Same as Attr("label",value)
func (e *Edge) Label(value fmt.Stringer) *Edge {
	e.SetAttribute(attributes.AttributeLabel, value)
	return e
}

// Solid sets the edge attribute "style" to "solid"
// Default style
func (e *Edge) Solid() *Edge {
	e.SetAttribute(attributes.AttributeStyle, attributes.NewString("solid"))
	return e
}

// Bold sets the edge attribute "style" to "bold"
func (e *Edge) Bold() *Edge {
	e.SetAttribute(attributes.AttributeStyle, attributes.NewString("bold"))
	return e
}

// Dashed sets the edge attribute "style" to "dashed"
func (e *Edge) Dashed() *Edge {
	e.SetAttribute(attributes.AttributeStyle, attributes.NewString("dashed"))
	return e
}

// Dotted sets the edge attribute "style" to "dotted"
func (e *Edge) Dotted() *Edge {
	e.SetAttribute(attributes.AttributeStyle, attributes.NewString("dotted"))
	return e
}

// Edge returns a new Edge between the "to" node of this Edge and the argument Node.
func (e *Edge) Edge(to *Node, labels ...string) *Edge {
	return e.graph.Edge(e.to, to, labels...)
}

// EdgesTo returns all existing edges between the "to" Node of thie Edge and the argument Node.
func (e *Edge) EdgesTo(to *Node) []*Edge {
	return e.graph.FindEdges(e.to, to)
}
