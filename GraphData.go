package dot

import (
	"bytes"
	"fmt"
	"io"

	"github.com/emicklei/dot/attributes"
)

// graph represents a dot graph with nodes and edges.
type graph struct {
	*attributes.Attributes
	id        string
	graphType attributes.GraphType
	strict    bool
	generator *UIDGenerator
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
	generator := NewUIDGenerator(24)
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

	newGraph := &graph{
		id:              options.ID,
		parent:          options.parent,
		Attributes:      attributes.NewAttributes(),
		graphType:       options.Type,
		strict:          options.Strict,
		generator:       NewUIDGenerator(24),
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
func (thisGraph *graph) ID() string {
	return thisGraph.id
}

func (thisGraph *graph) Subgraph(options *GraphOptions) Graph {
	if options == nil {
		options = &GraphOptions{}
	}

	// enforce subgraph type
	options.Type = attributes.GraphTypeSub

	// set parent
	options.parent = thisGraph

	sub := NewGraph(options)
	sub.SetAttributeString(attributes.KeyLabel, sub.ID())

	// save on parent with the generated ID
	thisGraph.subgraphs[sub.ID()] = sub

	return sub
}

// FindSubgraph returns the subgraph of the graph or one from its parents.
func (thisGraph *graph) FindSubgraph(id string) (Graph, bool) {
	if sub, ok := thisGraph.subgraphs[id]; ok {
		return sub, ok
	}
	if thisGraph.parent == nil {
		return nil, false
	}
	return thisGraph.parent.FindSubgraph(id)
}

func (thisGraph *graph) FindNode(id string) (Node, bool) {
	if n, ok := thisGraph.nodes[id]; ok {
		return n, ok
	}
	if thisGraph.parent == nil {
		return nil, false
	}
	return thisGraph.parent.FindNode(id)
}

// Node returns the node created with this id or creates a new node if absent.
// The node will have a label attribute with the id as its value. Use Label() to overwrite this.
// This method can be used as both a constructor and accessor.
// not thread safe!
func (thisGraph *graph) Node(id string) Node {
	if n, ok := thisGraph.FindNode(id); ok {
		return n
	}
	if len(id) == 0 {
		id = thisGraph.generator.String()
	}
	n := &node{
		id:         id,
		Attributes: attributes.NewAttributes(),
		graph:      thisGraph,
	}
	n.SetAttribute(attributes.KeyLabel, attributes.NewString(id))
	if thisGraph.nodeInitializer != nil {
		thisGraph.nodeInitializer(n)
	}
	// store local
	thisGraph.nodes[id] = n
	return n
}

// Edge creates a new edge between two nodes
func (thisGraph *graph) Edge(fromNode, toNode Node) StyledEdge {
	return thisGraph.EdgeWithAttributes(fromNode, toNode, nil)
}

// Edge creates a new edge between two nodes, and set the given attributes
func (thisGraph *graph) EdgeWithAttributes(fromNode, toNode Node, attr attributes.Reader) StyledEdge {
	e := &edge{
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
func (thisGraph *graph) FindEdges(fromNode, toNode Node) (found []Edge) {
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
func (thisGraph *graph) AddToSameRank(group string, nodes ...Node) {
	thisGraph.sameRank[group] = append(thisGraph.sameRank[group], nodes...)
}

// String returns the source in dot notation.
func (thisGraph *graph) String() string {
	b := new(bytes.Buffer)
	thisGraph.Write(b)
	return b.String()
}

func (thisGraph *graph) Write(w io.Writer) {
	thisGraph.IndentedWrite(NewIndentWriter(w))
}

// IndentedWrite write the graph to a writer using simple TAB indentation.
func (thisGraph *graph) IndentedWrite(w *IndentWriter) {
	if thisGraph.strict {
		fmt.Fprint(w, "strict ")
	}
	fmt.Fprintf(w, `%s "%s" {`, thisGraph.graphType, thisGraph.id)
	w.NewLineIndentWhile(func() {
		// subgraphs
		for _, key := range thisGraph.sortedSubgraphsKeys() {
			each := thisGraph.subgraphs[key]
			each.IndentedWrite(w)
		}
		// graph attributes
		thisGraph.Attributes.Write(w, false)
		w.NewLine()
		// graph nodes
		for _, key := range thisGraph.sortedNodesKeys() {
			each := thisGraph.nodes[key]
			fmt.Fprintf(w, `"%s"`, each.ID())
			each.Write(w, true)
			fmt.Fprintf(w, ";")
			w.NewLine()
		}
		// graph edges
		denoteEdge := "->"
		if thisGraph.graphType == "graph" {
			denoteEdge = "--"
		}
		for _, each := range thisGraph.sortedEdgesFromKeys() {
			all := thisGraph.edgesFrom[each]
			for _, each := range all {
				fmt.Fprintf(w, `"%s"%s"%s"`, each.From().ID(), denoteEdge, each.To().ID())
				each.Write(w, true)
				fmt.Fprint(w, ";")
				w.NewLine()
			}
		}
		for _, nodes := range thisGraph.sameRank {
			str := ""
			for _, n := range nodes {
				str += fmt.Sprintf(`"%s";`, n.ID())
			}
			fmt.Fprintf(w, "{rank=same; %s};", str)
			w.NewLine()
		}
	})
	fmt.Fprintf(w, "}")
	w.NewLine()
}

// VisitNodes visits all nodes recursively
func (thisGraph *graph) VisitNodes(callback func(node Node) (done bool)) {
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
func (thisGraph *graph) FindNodeByID(id string) (foundNode Node, found bool) {
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
func (thisGraph *graph) FindNodes() (nodes []Node) {
	var foundNodes []Node
	thisGraph.VisitNodes(func(node Node) (done bool) {
		foundNodes = append(foundNodes, node)
		return false
	})
	return foundNodes
}
