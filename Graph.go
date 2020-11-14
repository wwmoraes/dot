package dot

import (
	"github.com/wwmoraes/dot/attributes"
)

// GraphOptionFn is a functor that mutates graph options
type GraphOptionFn func(GraphOptions) error

// NodeInitializerFn mutates Nodes during their creation time
type NodeInitializerFn func(Node)

// EdgeInitializerFn mutates Edges during their creation time
type EdgeInitializerFn func(StyledEdge)

// Graph is implemented by dot-compatible graph values
type Graph interface {
	attributes.Identity
	attributes.Styleable
	attributes.Serializable

	// Root returns the root graph (i.e. the topmost, without a parent graph)
	Root() Graph
	// Type returns the graph type: directed, undirected or sub
	Type() GraphType
	// FindSubgraph returns the subgraph of this graph or from one of its parents
	FindSubgraph(id string) (Graph, bool)
	// Subgraph creates a subgraph of this graph
	Subgraph(optionsFn ...GraphOptionFn) (Graph, error)
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
