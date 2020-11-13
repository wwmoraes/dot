package dot

import (
	"github.com/wwmoraes/dot/attributes"
	"github.com/wwmoraes/dot/generators"
)

// GraphOptions contains the parameters used for graph creation
type GraphOptions struct {
	parent Graph
	// Generator is used to create IDs for objects to prevent colision
	Generator generators.IDGenerator
	// ID immutable id
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
	EdgeInitializer func(StyledEdge)
}

// Graph is implemented by dot-compatible graph values
type Graph interface {
	attributes.Identity
	attributes.Styleable
	attributes.Serializable

	// Root returns the root graph (i.e. the topmost, without a parent graph)
	Root() Graph
	// Type returns the graph type: directed, undirected or sub
	Type() attributes.GraphType
	// FindSubgraph returns the subgraph of this graph or from one of its parents
	FindSubgraph(id string) (Graph, bool)
	// Subgraph creates a subgraph of this graph
	Subgraph(options *GraphOptions) Graph
	// Node gets a node by id, or creates a new one if it doesn't exist
	Node(id string) Node
	// Edge creates a new edge between the two provided nodes
	Edge(n1, n2 Node) StyledEdge
	// Edge creates a new edge between the two provided nodes, and also set the
	// given attributes
	EdgeWithAttributes(n1, n2 Node, attributes attributes.Reader) StyledEdge
	// FindEdges gets all edges in the graph between the two provided nodes
	FindEdges(fromNode, toNode Node) (found []Edge)
	// FindNode gets a node by id
	FindNode(id string) (Node, bool)
	// VisitNodes runs the provided function on all nodes recursively
	VisitNodes(callback func(node Node) (done bool))
	// AddToSameRank adds the given nodes to the specified rank group, forcing
	// them to be rendered in the same row
	AddToSameRank(group string, nodes ...Node)
	// FindNodeByID return node by id
	FindNodeByID(id string) (foundNode Node, found bool)
	// FindNodes returns all nodes recursively
	FindNodes() (nodes []Node)
	// HasSubgraphs returns true if the graph has any subgraphs
	HasSubgraphs() bool
	// HasNodes returns true if the graph has any nodes
	HasNodes() bool
	// HasEdges returns true if the graph has any nodes
	HasEdges() bool
	// HasSameRankNodes returns true if the graph has nodes grouped as same rank
	HasSameRankNodes() bool
	// IsStrict return true if the graph is set as strict
	IsStrict() bool
}
