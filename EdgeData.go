package dot

import (
	"fmt"

	"github.com/emicklei/dot/attributes"
)

type edge struct {
	*attributes.Attributes
	graph      Graph
	from, to   Node
	internalID string
}

func (thisEdge *edge) String() string {
	// TODO
	return thisEdge.internalID
}

func (thisEdge *edge) From() Node {
	return thisEdge.from
}

func (thisEdge *edge) To() Node {
	return thisEdge.to
}

// Label sets "label"=value and returns the Edge.
// Same as SetAttribute(attributes.KeyLabel, value)
func (thisEdge *edge) Label(value fmt.Stringer) Edge {
	thisEdge.SetAttribute(attributes.KeyLabel, value)
	return thisEdge
}

// Solid sets the edge attribute "style" to "solid"
// Default style
func (thisEdge *edge) Solid() Edge {
	thisEdge.SetAttribute(attributes.KeyStyle, attributes.NewString("solid"))
	return thisEdge
}

// Bold sets the edge attribute "style" to "bold"
func (thisEdge *edge) Bold() Edge {
	thisEdge.SetAttribute(attributes.KeyStyle, attributes.NewString("bold"))
	return thisEdge
}

// Dashed sets the edge attribute "style" to "dashed"
func (thisEdge *edge) Dashed() Edge {
	thisEdge.SetAttribute(attributes.KeyStyle, attributes.NewString("dashed"))
	return thisEdge
}

// Dotted sets the edge attribute "style" to "dotted"
func (thisEdge *edge) Dotted() Edge {
	thisEdge.SetAttribute(attributes.KeyStyle, attributes.NewString("dotted"))
	return thisEdge
}

// Edge returns a new Edge between the "to" node of this Edge and the argument Node
func (thisEdge *edge) Edge(to Node) Edge {
	return thisEdge.EdgeWithAttributes(to, nil)
}

// EdgeWithAttributes returns a new Edge between the "to" node of this Edge and the argument Node
func (thisEdge *edge) EdgeWithAttributes(to Node, attributes attributes.Reader) Edge {
	return thisEdge.graph.EdgeWithAttributes(thisEdge.to, to, attributes)
}

// EdgesTo returns all existing edges between the "to" Node of this Edge and the argument Node.
func (thisEdge *edge) EdgesTo(to Node) []Edge {
	return thisEdge.graph.FindEdges(thisEdge.to, to)
}
