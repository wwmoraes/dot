package dot

import (
	"fmt"

	"github.com/emicklei/dot/attributes"
)

// Graph context/area that contains nodes and edges
type Graph interface {
	attributes.Object
	fmt.Stringer
	Root() Graph
	FindSubgraph(id string) (Graph, bool)
	Subgraph(options *GraphOptions) Graph
	// SubgraphOld(id string, options ...GraphOption) Graph
	Node(id string) Node
	Edge(n1, n2 Node) Edge
	EdgeWithAttributes(n1, n2 Node, attributes attributes.Reader) Edge
	FindEdges(fromNode, toNode Node) (found []Edge)
	FindNode(id string) (Node, bool)
	IndentedWrite(w *IndentWriter)
	VisitNodes(callback func(node Node) (done bool))
	AddToSameRank(group string, nodes ...Node)
	FindNodeByID(id string) (foundNode Node, found bool)
	FindNodes() (nodes []Node)
}
