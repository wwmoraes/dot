package dot

import (
	"fmt"

	"github.com/emicklei/dot/attributes"
)

// EdgeData represents a graph edge between two Nodes.
type EdgeData struct {
	*attributes.Attributes
	graph    Graph
	from, to Node
	id       string
}

// ID returns the immutable id
func (e *EdgeData) ID() string {
	return e.id
}

func (e *EdgeData) String() string {
	// TODO
	return e.id
}

func (e *EdgeData) From() Node {
	return e.from
}

func (e *EdgeData) To() Node {
	return e.to
}

// Label sets "label"=value and returns the Edge.
// Same as SetAttribute(attributes.KeyLabel, value)
func (e *EdgeData) Label(value fmt.Stringer) Edge {
	e.SetAttribute(attributes.KeyLabel, value)
	return e
}

// Solid sets the edge attribute "style" to "solid"
// Default style
func (e *EdgeData) Solid() Edge {
	e.SetAttribute(attributes.KeyStyle, attributes.NewString("solid"))
	return e
}

// Bold sets the edge attribute "style" to "bold"
func (e *EdgeData) Bold() Edge {
	e.SetAttribute(attributes.KeyStyle, attributes.NewString("bold"))
	return e
}

// Dashed sets the edge attribute "style" to "dashed"
func (e *EdgeData) Dashed() Edge {
	e.SetAttribute(attributes.KeyStyle, attributes.NewString("dashed"))
	return e
}

// Dotted sets the edge attribute "style" to "dotted"
func (e *EdgeData) Dotted() Edge {
	e.SetAttribute(attributes.KeyStyle, attributes.NewString("dotted"))
	return e
}

// Edge returns a new Edge between the "to" node of this Edge and the argument Node
func (e *EdgeData) Edge(to Node) Edge {
	return e.EdgeWithAttributes(to, nil)
}

// EdgeWithAttributes returns a new Edge between the "to" node of this Edge and the argument Node
func (e *EdgeData) EdgeWithAttributes(to Node, attributes attributes.Reader) Edge {
	return e.graph.EdgeWithAttributes(e.to, to, attributes)
}

// EdgesTo returns all existing edges between the "to" Node of this Edge and the argument Node.
func (e *EdgeData) EdgesTo(to Node) []Edge {
	return e.graph.FindEdges(e.to, to)
}
