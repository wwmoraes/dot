package dot

import (
	"fmt"

	"github.com/emicklei/dot/attributes"
)

// GraphOptions parameters used for graph creation
type GraphOptions struct {
	parent Graph
	// ID returns the immutable id
	ID string
	// Type the graph type (directed, undirected or sub)
	Type attributes.GraphType
	// Cluster forbids the creation of multi-edges i.e.:
	//
	// on directed graphs, only one one edge between a given pair of head and tail nodes is allowed
	//
	// on undirected graphs, only one edge between the same two nodes is allowed
	Strict bool
	// Cluster denotes if the graph is the special cluster subgraph, whose name
	// starts with "cluster_"
	Cluster bool
	// NodeInitializer applies defaults to newly created nodes
	NodeInitializer func(Node)
	// EdgeInitializer applies defaults to newly created edges
	EdgeInitializer func(Edge)
}

// Graph context/area that contains nodes and edges
type Graph interface {
	attributes.Object
	fmt.Stringer
	Root() Graph
	FindSubgraph(id string) (Graph, bool)
	Subgraph(options *GraphOptions) Graph
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
