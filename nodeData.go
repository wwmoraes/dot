package dot

import (
	"fmt"
	"io"

	"github.com/wwmoraes/dot/attributes"
)

// nodeData represents a dot nodeData.
type nodeData struct {
	*attributes.Attributes
	graph Graph
	id    string
}

// ID returns the immutable id
func (thisNode *nodeData) ID() string {
	return thisNode.id
}

func (thisNode *nodeData) String() string {
	// TODO
	return thisNode.id
}

// Edge creates an Edge to a node
func (thisNode *nodeData) Edge(toNode Node) Edge {
	return thisNode.graph.EdgeWithAttributes(thisNode, toNode, nil)
}

// EdgeWithAttributes creates an Edge with the provided attributes to the a node
func (thisNode *nodeData) EdgeWithAttributes(toNode Node, attributes attributes.Reader) Edge {
	return thisNode.graph.EdgeWithAttributes(thisNode, toNode, attributes)
}

// EdgesTo returns all edges between this Node and the target Node
func (thisNode *nodeData) EdgesTo(toNode Node) []Edge {
	return thisNode.graph.FindEdges(thisNode, toNode)
}

func (thisNode *nodeData) WriteTo(device io.Writer) (n int64, err error) {
	written32, err := fmt.Fprintf(device, `"%s"`, thisNode.ID())
	n += int64(written32)
	if err != nil {
		return
	}

	written64, err := thisNode.Attributes.WriteTo(device)
	n += written64
	if err != nil {
		return
	}

	written32, err = fmt.Fprintf(device, ";")
	n += int64(written32)

	return
}
