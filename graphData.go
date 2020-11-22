package dot

import (
	"bytes"
	"fmt"
	"io"

	"github.com/wwmoraes/dot/attributes"
	"github.com/wwmoraes/dot/generators"
)

// graphData is a dot-compatible graph that stores child components, and
// auto-generates IDs internally
type graphData struct {
	*attributes.Attributes
	id        string
	graphType GraphType
	strict    bool
	generator generators.IDGenerator
	nodes     map[string]Node
	edgesFrom map[string][]StyledEdge
	subgraphs map[string]Graph
	parent    Graph
	sameRank  map[string][]Node
	//
	nodeInitializer NodeInitializerFn
	edgeInitializer EdgeInitializerFn
}

// New return a Graph after all option functions are processed
func New(optionsFn ...GraphOptionFn) (Graph, error) {
	options, err := NewGraphOptions(optionsFn...)

	if err != nil {
		return nil, err
	}

	return NewWithOptions(options)
}

// NewWithOptions returns a Graph using the options values
func NewWithOptions(options GraphOptions) (Graph, error) {
	if options.Generator() == nil {
		return nil, ErrGraphWithoutGenerator
	}

	if options.Type() == GraphTypeSub {
		if options.Parent() == nil {
			return nil, ErrSubgraphWithoutParent
		}
	} else {
		if options.Parent() != nil {
			return nil, ErrRootWithParent
		}
		if options.Cluster() {
			return nil, ErrRootAsCluster
		}
	}

	return &graphData{
		id:              options.ID(),
		parent:          options.Parent(),
		Attributes:      attributes.NewAttributes(),
		graphType:       options.Type(),
		strict:          options.Strict(),
		generator:       options.Generator(),
		nodes:           map[string]Node{},
		edgesFrom:       map[string][]StyledEdge{},
		subgraphs:       map[string]Graph{},
		sameRank:        map[string][]Node{},
		nodeInitializer: options.NodeInitializer(),
		edgeInitializer: options.EdgeInitializer(),
	}, nil
}

// ID returns the immutable id
func (thisGraph *graphData) ID() string {
	return thisGraph.id
}

// Root returns the root graph (i.e. the topmost, without a parent graph)
func (thisGraph *graphData) Root() Graph {
	if thisGraph.parent == nil {
		return thisGraph
	}

	return thisGraph.parent.Root()
}

// Type returns the graph type: directed, undirected or sub
func (thisGraph *graphData) Type() GraphType {
	return thisGraph.graphType
}

func (thisGraph *graphData) Subgraph(optionsFn ...GraphOptionFn) (Graph, error) {
	subgraphOptions := [...]GraphOptionFn{
		WithParent(thisGraph),
		WithGenerator(thisGraph.generator),
	}

	newOptionsFn := make([]GraphOptionFn, 0, len(subgraphOptions)+len(optionsFn))

	for _, subgraphOption := range subgraphOptions {
		newOptionsFn = append(newOptionsFn, subgraphOption)
	}

	newOptionsFn = append(newOptionsFn, optionsFn...)

	graph, err := New(newOptionsFn...)
	if err != nil {
		return nil, err
	}

	// save on parent with the generated ID
	thisGraph.subgraphs[graph.ID()] = graph

	return graph, nil
}

// FindSubgraph returns the subgraph of the graph or one from its parents.
func (thisGraph *graphData) FindSubgraph(id string) (Graph, bool) {
	if sub, ok := thisGraph.subgraphs[id]; ok {
		return sub, ok
	}
	if thisGraph.parent == nil {
		return nil, false
	}
	return thisGraph.parent.FindSubgraph(id)
}

func (thisGraph *graphData) FindNode(id string) (Node, bool) {
	if n, ok := thisGraph.nodes[id]; ok {
		return n, ok
	}
	if thisGraph.parent == nil {
		return nil, false
	}
	return thisGraph.parent.FindNode(id)
}

// Node returns the node created with this id or creates a new node if absent
// This method can be used as both a constructor and accessor.
// not thread safe!
func (thisGraph *graphData) Node(id string) Node {
	if n, ok := thisGraph.FindNode(id); ok {
		return n
	}
	if len(id) == 0 {
		id = thisGraph.generator.String()
	}
	n := &nodeData{
		id:         id,
		Attributes: attributes.NewAttributes(),
		graph:      thisGraph,
	}
	if thisGraph.nodeInitializer != nil {
		thisGraph.nodeInitializer(n)
	}
	// store local
	thisGraph.nodes[id] = n
	return n
}

// Edge creates a new edge between two nodes
func (thisGraph *graphData) Edge(fromNode, toNode Node) StyledEdge {
	return thisGraph.EdgeWithAttributes(fromNode, toNode, nil)
}

// Edge creates a new edge between two nodes, and set the given attributes
func (thisGraph *graphData) EdgeWithAttributes(fromNode, toNode Node, attr attributes.Reader) StyledEdge {
	e := &edgeData{
		from:       fromNode,
		to:         toNode,
		Attributes: attributes.NewAttributesFrom(attr),
		graph:      thisGraph}

	if thisGraph.edgeInitializer != nil {
		thisGraph.edgeInitializer(e)
	}

	thisGraph.edgesFrom[fromNode.ID()] = append(thisGraph.edgesFrom[fromNode.ID()], e)

	return e
}

// FindEdges finds all edges in the graph that go from the fromNode to the toNode.
// Otherwise, returns an empty slice.
func (thisGraph *graphData) FindEdges(fromNode, toNode Node) (found []Edge) {
	found = make([]Edge, 0)
	if edges, ok := thisGraph.edgesFrom[fromNode.ID()]; ok {
		for _, e := range edges {
			if e.To().ID() == toNode.ID() {
				found = append(found, e)
			}
		}
	}
	return found
}

// AddToSameRank adds the given nodes to the specified rank group, forcing them to be rendered in the same row
func (thisGraph *graphData) AddToSameRank(group string, nodes ...Node) {
	thisGraph.sameRank[group] = append(thisGraph.sameRank[group], nodes...)
}

// String returns the graph transformed into string dot notation
func (thisGraph *graphData) String() (string, error) {
	var b bytes.Buffer

	_, err := thisGraph.WriteTo(&b)

	return b.String(), err
}

func (thisGraph *graphData) WriteTo(w io.Writer) (n int64, err error) {
	var written32 int
	var written64 int64

	// write strict tag
	written64, err = thisGraph.writeStrictTagTo(w)
	n += written64
	if err != nil {
		return
	}

	// graph type
	written32, err = fmt.Fprintf(w, `%s`, thisGraph.graphType)
	n += int64(written32)
	if err != nil {
		return
	}

	// write graph id
	if len(thisGraph.id) > 0 {
		written32, err = fmt.Fprintf(w, ` "%s"`, thisGraph.id)
		n += int64(written32)
		if err != nil {
			return
		}
	}

	// write open block
	written32, err = fmt.Fprint(w, " {")
	n += int64(written32)
	if err != nil {
		return
	}

	// write graph body
	written64, err = thisGraph.writeGraphBodyTo(w)
	n += written64
	if err != nil {
		return
	}

	// write close block
	written32, err = fmt.Fprintf(w, "}")
	n += int64(written32)

	return
}

// VisitNodes visits all nodes recursively
func (thisGraph *graphData) VisitNodes(callback func(node Node) (done bool)) {
	for _, node := range thisGraph.nodes {
		done := callback(node)
		if done {
			return
		}
	}

	for _, subGraph := range thisGraph.subgraphs {
		subGraph.VisitNodes(callback)
	}
}

// FindNodeByID return node by id
func (thisGraph *graphData) FindNodeByID(id string) (foundNode Node, found bool) {
	thisGraph.VisitNodes(func(node Node) (done bool) {
		if node.ID() == id {
			found = true
			foundNode = node
			return true
		}
		return false
	})
	return
}

// FindNodes returns all nodes recursively
func (thisGraph *graphData) FindNodes() (nodes []Node) {
	var foundNodes []Node
	thisGraph.VisitNodes(func(node Node) (done bool) {
		foundNodes = append(foundNodes, node)
		return false
	})
	return foundNodes
}

// HasSubgraphs returns true if the graph has any subgraphs
func (thisGraph *graphData) HasSubgraphs() bool {
	return len(thisGraph.subgraphs) > 0
}

// HasNodes returns true if the graph has any nodes
func (thisGraph *graphData) HasNodes() bool {
	return len(thisGraph.nodes) > 0
}

// HasEdges returns true if the graph has any edges
func (thisGraph *graphData) HasEdges() bool {
	return len(thisGraph.edgesFrom) > 0
}

// HasSameRankNodes returns true if the graph has nodes grouped as same rank
func (thisGraph *graphData) HasSameRankNodes() bool {
	return len(thisGraph.sameRank) > 0
}

func (thisGraph *graphData) writeGraphBodyTo(w io.Writer) (n int64, err error) {
	// write attributes
	written64, err := thisGraph.writeAttributesTo(w)
	n += written64
	if err != nil {
		return
	}

	// write subgraphs
	written64, err = thisGraph.writeSubgraphsTo(w)
	n += written64
	if err != nil {
		return
	}

	// write nodes
	written64, err = thisGraph.writeNodesTo(w)
	n += written64
	if err != nil {
		return
	}

	// node groups
	written64, err = thisGraph.writeSameRankNodesTo(w)
	n += written64
	if err != nil {
		return
	}

	// write edges
	written64, err = thisGraph.writeEdgesTo(w)
	n += written64
	if err != nil {
		return
	}

	return n, err
}

func (thisGraph *graphData) writeAttributesTo(w io.Writer) (n int64, err error) {
	if !thisGraph.HasAttributes() {
		return 0, nil
	}

	written32, err := fmt.Fprintf(w, "graph ")
	n += int64(written32)
	if err != nil {
		return n, err
	}

	written64, err := thisGraph.Attributes.WriteTo(w)
	n += written64
	if err != nil {
		return n, err
	}

	written32, err = fmt.Fprintf(w, ";")
	n += int64(written32)
	if err != nil {
		return n, err
	}

	return n, err
}

func (thisGraph *graphData) writeSubgraphsTo(w io.Writer) (n int64, err error) {
	if !thisGraph.HasSubgraphs() {
		return 0, nil
	}

	for _, key := range thisGraph.sortedSubgraphsKeys() {
		each := thisGraph.subgraphs[key]
		written64, err := each.WriteTo(w)
		n += written64
		if err != nil {
			return n, err
		}
	}

	return n, err
}

func (thisGraph *graphData) writeNodesTo(w io.Writer) (n int64, err error) {
	if !thisGraph.HasNodes() {
		return 0, nil
	}

	for _, key := range thisGraph.sortedNodesKeys() {
		each := thisGraph.nodes[key]
		written64, err := each.WriteTo(w)
		n += written64
		if err != nil {
			return n, err
		}
	}

	return n, err
}

func (thisGraph *graphData) writeEdgesTo(w io.Writer) (n int64, err error) {
	if !thisGraph.HasEdges() {
		return 0, nil
	}

	for _, each := range thisGraph.sortedEdgesFromKeys() {
		all := thisGraph.edgesFrom[each]
		for _, each := range all {
			written64, err := each.WriteTo(w)
			n += written64
			if err != nil {
				return n, err
			}
		}
	}

	return n, err
}

func (thisGraph *graphData) writeSameRankNodesTo(w io.Writer) (n int64, err error) {
	if !thisGraph.HasSameRankNodes() {
		return 0, nil
	}

	for _, nodes := range thisGraph.sameRank {
		// open group
		written32, err := fmt.Fprintf(w, "{")
		n += int64(written32)
		if err != nil {
			return n, err
		}
		// set rank attribute
		written32, err = fmt.Fprintf(w, "rank=same;")
		n += int64(written32)
		if err != nil {
			return n, err
		}
		// write group nodes
		for _, node := range nodes {
			written64, err := node.WriteTo(w)
			n += written64
			if err != nil {
				return n, err
			}
		}
		// close group
		written32, err = fmt.Fprintf(w, "}")
		n += int64(written32)
		if err != nil {
			return n, err
		}
	}

	return n, err
}

func (thisGraph *graphData) writeStrictTagTo(w io.Writer) (n int64, err error) {
	if !thisGraph.IsStrict() {
		return 0, nil
	}

	written32, err := fmt.Fprint(w, "strict ")
	n += int64(written32)
	if err != nil {
		return n, err
	}

	return n, err
}

// IsStrict return true if the graph is set as strict
func (thisGraph *graphData) IsStrict() bool {
	return thisGraph.strict
}
