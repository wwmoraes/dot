package dot

import (
	"bytes"
	"fmt"
	"io"

	"github.com/wwmoraes/dot/attributes"
	"github.com/wwmoraes/dot/formatters"
	"github.com/wwmoraes/dot/generators"
)

// graphData is a dot-compatible graph that stores child components, and
// auto-generates IDs internally
type graphData struct {
	*attributes.Attributes
	id        string
	graphType attributes.GraphType
	strict    bool
	generator generators.IDGenerator
	nodes     map[string]Node
	edgesFrom map[string][]StyledEdge
	subgraphs map[string]Graph
	parent    Graph
	sameRank  map[string][]Node
	//
	nodeInitializer func(Node)
	edgeInitializer func(StyledEdge)
}

// NewGraph return a new initialized Graph
//
// if id is "-", a randonly generated ID will be set
func NewGraph(options *GraphOptions) Graph {
	generator := generators.NewRandTimeIDGenerator(24)
	if options == nil {
		options = &GraphOptions{}
	}

	if options.ID == "-" {
		options.ID = generator.String()
	}

	if options.Cluster {
		options.ID = fmt.Sprintf("cluster_%s", options.ID)
	}

	if options.Type == "" {
		options.Type = attributes.GraphTypeDirected
	}

	if options.Type == attributes.GraphTypeSub && options.parent == nil {
		panic("cannot create subgraph without parent")
	} else if options.Type != attributes.GraphTypeSub && options.parent != nil {
		panic("cannot create graph with parent")
	}

	if options.Generator == nil {
		options.Generator = generators.NewRandTimeIDGenerator(24)
	}

	newGraph := &graphData{
		id:              options.ID,
		parent:          options.parent,
		Attributes:      attributes.NewAttributes(),
		graphType:       options.Type,
		strict:          options.Strict,
		generator:       options.Generator,
		nodes:           map[string]Node{},
		edgesFrom:       map[string][]StyledEdge{},
		subgraphs:       map[string]Graph{},
		sameRank:        map[string][]Node{},
		nodeInitializer: options.NodeInitializer,
		edgeInitializer: options.EdgeInitializer,
	}

	return newGraph
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
func (thisGraph *graphData) Type() attributes.GraphType {
	return thisGraph.graphType
}

func (thisGraph *graphData) Subgraph(options *GraphOptions) Graph {
	if options == nil {
		options = &GraphOptions{}
	}

	// enforce subgraph type
	options.Type = attributes.GraphTypeSub

	// set parent
	options.parent = thisGraph

	// share generator
	options.Generator = thisGraph.generator

	sub := NewGraph(options)

	// save on parent with the generated ID
	thisGraph.subgraphs[sub.ID()] = sub

	return sub
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
		internalID: thisGraph.generator.String(),
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

// String returns the source in dot notation.
func (thisGraph *graphData) String() string {
	b := new(bytes.Buffer)
	thisGraph.Write(b)
	return b.String()
}

func (thisGraph *graphData) Write(w io.Writer) {
	if thisGraph.strict {
		fmt.Fprint(w, "strict ")
	}

	// open graph
	fmt.Fprintf(w, `%s`, thisGraph.graphType)
	if len(thisGraph.id) > 0 {
		fmt.Fprintf(w, ` "%s"`, thisGraph.id)
	}
	fmt.Fprint(w, " {")

	// write attributes
	if thisGraph.HasAttributes() {
		fmt.Fprintf(w, "graph ")
		thisGraph.Attributes.WriteAttributes(w)
		fmt.Fprintf(w, ";")
	}

	// iterate and write subgraphs
	for _, key := range thisGraph.sortedSubgraphsKeys() {
		each := thisGraph.subgraphs[key]
		each.Write(w)
	}

	// iterate and write nodes
	for _, key := range thisGraph.sortedNodesKeys() {
		each := thisGraph.nodes[key]
		each.Write(w)
	}

	// iterate and write node groups
	for _, nodes := range thisGraph.sameRank {
		// open group
		fmt.Fprintf(w, "{")
		// set rank attribute
		fmt.Fprintf(w, "rank=same;")
		// write group nodes
		for _, n := range nodes {
			n.Write(w)
		}
		// close group
		fmt.Fprintf(w, "}")
	}

	// iterate and write edges
	for _, each := range thisGraph.sortedEdgesFromKeys() {
		all := thisGraph.edgesFrom[each]
		for _, each := range all {
			each.Write(w)
		}
	}

	// close graph
	fmt.Fprintf(w, "}")
}

// IndentedWrite write the graph to a writer using simple TAB indentation.
func (thisGraph *graphData) IndentedWrite(w formatters.IndentedWriter) {
	if thisGraph.strict {
		fmt.Fprint(w, "strict ")
	}
	fmt.Fprintf(w, `%s "%s" {`, thisGraph.graphType, thisGraph.id)
	w.NewLineIndentWhile(func() {
		// subgraphs
		for _, key := range thisGraph.sortedSubgraphsKeys() {
			each := thisGraph.subgraphs[key]
			each.IndentedWrite(w)
			w.NewLine()
		}
		// graph attributes
		thisGraph.Attributes.WriteAttributes(w)
		// graph nodes
		for _, key := range thisGraph.sortedNodesKeys() {
			w.NewLine()
			each := thisGraph.nodes[key]
			each.Write(w)
		}
		// graph edges
		for _, each := range thisGraph.sortedEdgesFromKeys() {
			all := thisGraph.edgesFrom[each]
			for _, each := range all {
				w.NewLine()
				each.Write(w)
			}
		}
		for _, nodes := range thisGraph.sameRank {
			w.NewLine()
			str := ""
			for _, n := range nodes {
				str += fmt.Sprintf(`"%s";`, n.ID())
			}
			fmt.Fprintf(w, "{rank=same; %s};", str)
		}
	})
	fmt.Fprintf(w, "}")
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
