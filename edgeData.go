package dot

import (
	"bytes"
	"fmt"
	"io"

	"github.com/wwmoraes/dot/attributes"
)

type edgeData struct {
	*attributes.Attributes
	graph      Graph
	from, to   Node
	internalID string
}

// String returns the graph transformed into string dot notation
func (thisEdge *edgeData) String() (string, error) {
	var b bytes.Buffer

	_, err := thisEdge.WriteTo(&b)

	return b.String(), err
}

func (thisEdge *edgeData) From() Node {
	return thisEdge.from
}

func (thisEdge *edgeData) To() Node {
	return thisEdge.to
}

// Solid sets the edge attribute "style" to "solid"
// Default style
func (thisEdge *edgeData) Solid() Edge {
	thisEdge.SetAttribute(attributes.KeyStyle, attributes.NewString("solid"))
	return thisEdge
}

// Bold sets the edge attribute "style" to "bold"
func (thisEdge *edgeData) Bold() Edge {
	thisEdge.SetAttribute(attributes.KeyStyle, attributes.NewString("bold"))
	return thisEdge
}

// Dashed sets the edge attribute "style" to "dashed"
func (thisEdge *edgeData) Dashed() Edge {
	thisEdge.SetAttribute(attributes.KeyStyle, attributes.NewString("dashed"))
	return thisEdge
}

// Dotted sets the edge attribute "style" to "dotted"
func (thisEdge *edgeData) Dotted() Edge {
	thisEdge.SetAttribute(attributes.KeyStyle, attributes.NewString("dotted"))
	return thisEdge
}

// Edge returns a new Edge between the "to" node of this Edge and the argument Node
func (thisEdge *edgeData) Edge(to Node) Edge {
	return thisEdge.EdgeWithAttributes(to, nil)
}

// EdgeWithAttributes returns a new Edge between the "to" node of this Edge and the argument Node
func (thisEdge *edgeData) EdgeWithAttributes(to Node, attributes attributes.Reader) Edge {
	return thisEdge.graph.EdgeWithAttributes(thisEdge.to, to, attributes)
}

// EdgesTo returns all existing edges between the "to" Node of this Edge and the argument Node.
func (thisEdge *edgeData) EdgesTo(to Node) []Edge {
	return thisEdge.graph.FindEdges(thisEdge.to, to)
}

func (thisEdge *edgeData) WriteTo(device io.Writer) (n int64, err error) {
	denoteEdge := attributes.EdgeTypeUndirected

	if thisEdge.graph.Root().Type() == attributes.GraphTypeDirected {
		denoteEdge = attributes.EdgeTypeDirected
	}

	written32, err := fmt.Fprintf(device, `"%s"%s"%s"`, thisEdge.From().ID(), denoteEdge, thisEdge.To().ID())
	n += int64(written32)
	if err != nil {
		return
	}

	written64, err := thisEdge.Attributes.WriteTo(device)
	n += written64
	if err != nil {
		return
	}

	written32, err = fmt.Fprint(device, ";")
	n += int64(written32)
	return
}
