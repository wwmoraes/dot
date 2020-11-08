package dot

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/emicklei/dot/attributes"
)

// Graph represents a dot graph with nodes and edges.
type Graph struct {
	*attributes.Attributes
	id        string
	graphType string
	strict    bool
	generator *UIDGenerator
	nodes     map[string]*Node
	edgesFrom map[string][]*Edge
	subgraphs map[string]*Graph
	parent    *Graph
	sameRank  map[string][]*Node
	//
	nodeInitializer func(*Node)
	edgeInitializer func(*Edge)
}

// NewGraph return a new initialized Graph
//
// if id is "-", a randonly generated ID will be set
func NewGraph(options ...GraphOption) *Graph {
	graph := &Graph{
		Attributes: attributes.NewAttributes(),
		graphType:  Directed.Name,
		generator:  NewUIDGenerator(24),
		nodes:      map[string]*Node{},
		edgesFrom:  map[string][]*Edge{},
		subgraphs:  map[string]*Graph{},
		sameRank:   map[string][]*Node{},
	}
	for _, each := range options {
		each.Apply(graph)
	}

	if graph.id == "-" {
		graph.id = graph.generator.String()
	}

	return graph
}

// Label sets the "label" attribute value.
func (g *Graph) Label(label string) *Graph {
	g.SetAttribute("label", attributes.NewString(label))
	return g
}

func (g *Graph) beCluster() {
	g.id = "cluster_" + g.id
}

// Root returns the top-level graph if this was a subgraph.
func (g *Graph) Root() *Graph {
	if g.parent == nil {
		return g
	}
	return g.parent.Root()
}

// FindSubgraph returns the subgraph of the graph or one from its parents.
func (g *Graph) FindSubgraph(id string) (*Graph, bool) {
	sub, ok := g.subgraphs[id]
	if !ok {
		if g.parent != nil {
			return g.parent.FindSubgraph(id)
		}
	}
	return sub, ok
}

// Subgraph returns the Graph with the given id ; creates one if absent.
// The label attribute is also set to the id ; use Label() to overwrite it.
//
// if id is "-", a randonly generated ID will be set
func (g *Graph) Subgraph(id string, options ...GraphOption) *Graph {
	sub, ok := g.subgraphs[id]
	if ok {
		return sub
	}
	sub = NewGraph(Sub, &GraphIDOption{id})
	sub.SetAttribute("label", attributes.NewString(id)) // for consistency with Node creation behavior.
	for _, each := range options {
		each.Apply(sub)
	}
	sub.parent = g
	sub.edgeInitializer = g.edgeInitializer
	sub.nodeInitializer = g.nodeInitializer
	g.subgraphs[id] = sub
	return sub
}

func (g *Graph) findNode(id string) (*Node, bool) {
	if n, ok := g.nodes[id]; ok {
		return n, ok
	}
	if g.parent == nil {
		return &Node{id: "void"}, false
	}
	return g.parent.findNode(id)
}

// NodeInitializer sets a function that is called (if not nil) when a Node is implicitly created.
func (g *Graph) NodeInitializer(callback func(n *Node)) {
	g.nodeInitializer = callback
}

// EdgeInitializer sets a function that is called (if not nil) when an Edge is implicitly created.
func (g *Graph) EdgeInitializer(callback func(e *Edge)) {
	g.edgeInitializer = callback
}

// Node returns the node created with this id or creates a new node if absent.
// The node will have a label attribute with the id as its value. Use Label() to overwrite this.
// This method can be used as both a constructor and accessor.
// not thread safe!
func (g *Graph) Node(id string) *Node {
	if n, ok := g.findNode(id); ok {
		return n
	}
	if len(id) == 0 {
		id = g.generator.String()
	}
	n := &Node{
		id:         id,
		Attributes: attributes.NewAttributes(),
		graph:      g,
	}
	n.SetAttribute(attributes.AttributeLabel, attributes.NewString(id))
	if g.nodeInitializer != nil {
		g.nodeInitializer(n)
	}
	// store local
	g.nodes[id] = n
	return n
}

// Edge creates a new edge between two nodes.
// Nodes can be have multiple edges to the same other node (or itself).
// If one or more labels are given then the "label" attribute is set to the edge.
func (g *Graph) Edge(fromNode, toNode *Node, labels ...string) *Edge {
	// assume fromNode owner == toNode owner
	edgeOwner := g
	if fromNode.graph != toNode.graph { // 1 or 2 are subgraphs
		edgeOwner = commonParentOf(fromNode.graph, toNode.graph)
	}
	e := &Edge{
		from:       fromNode,
		to:         toNode,
		id:         g.generator.String(),
		Attributes: attributes.NewAttributes(),
		graph:      edgeOwner}
	if len(labels) > 0 {
		e.SetAttribute("label", attributes.NewString(strings.Join(labels, ",")))
	}
	if g.edgeInitializer != nil {
		g.edgeInitializer(e)
	}
	edgeOwner.edgesFrom[fromNode.id] = append(edgeOwner.edgesFrom[fromNode.id], e)
	return e
}

// FindEdges finds all edges in the graph that go from the fromNode to the toNode.
// Otherwise, returns an empty slice.
func (g *Graph) FindEdges(fromNode, toNode *Node) (found []*Edge) {
	found = make([]*Edge, 0)
	edgeOwner := g
	if fromNode.graph != toNode.graph {
		edgeOwner = commonParentOf(fromNode.graph, toNode.graph)
	}
	if edges, ok := edgeOwner.edgesFrom[fromNode.id]; ok {
		for _, e := range edges {
			if e.to.id == toNode.id {
				found = append(found, e)
			}
		}
	}
	return found
}

func commonParentOf(one *Graph, two *Graph) *Graph {
	// TODO
	return one.Root()
}

// AddToSameRank adds the given nodes to the specified rank group, forcing them to be rendered in the same row
func (g *Graph) AddToSameRank(group string, nodes ...*Node) {
	g.sameRank[group] = append(g.sameRank[group], nodes...)
}

// String returns the source in dot notation.
func (g *Graph) String() string {
	b := new(bytes.Buffer)
	g.Write(b)
	return b.String()
}

func (g *Graph) Write(w io.Writer) {
	g.IndentedWrite(NewIndentWriter(w))
}

// IndentedWrite write the graph to a writer using simple TAB indentation.
func (g *Graph) IndentedWrite(w *IndentWriter) {
	if g.strict {
		fmt.Fprint(w, "strict ")
	}
	fmt.Fprintf(w, `%s "%s" {`, g.graphType, g.id)
	w.NewLineIndentWhile(func() {
		// subgraphs
		for _, key := range g.sortedSubgraphsKeys() {
			each := g.subgraphs[key]
			each.IndentedWrite(w)
		}
		// graph attributes
		g.Attributes.Write(w, false)
		w.NewLine()
		// graph nodes
		for _, key := range g.sortedNodesKeys() {
			each := g.nodes[key]
			fmt.Fprintf(w, `"%s"`, each.id)
			each.Write(w, true)
			fmt.Fprintf(w, ";")
			w.NewLine()
		}
		// graph edges
		denoteEdge := "->"
		if g.graphType == "graph" {
			denoteEdge = "--"
		}
		for _, each := range g.sortedEdgesFromKeys() {
			all := g.edgesFrom[each]
			for _, each := range all {
				fmt.Fprintf(w, `"%s"%s"%s"`, each.from.id, denoteEdge, each.to.id)
				each.Write(w, true)
				fmt.Fprint(w, ";")
				w.NewLine()
			}
		}
		for _, nodes := range g.sameRank {
			str := ""
			for _, n := range nodes {
				str += fmt.Sprintf(`"%s";`, n.id)
			}
			fmt.Fprintf(w, "{rank=same; %s};", str)
			w.NewLine()
		}
	})
	fmt.Fprintf(w, "}")
	w.NewLine()
}

// VisitNodes visits all nodes recursively
func (g *Graph) VisitNodes(callback func(node *Node) (done bool)) {
	for _, node := range g.nodes {
		done := callback(node)
		if done {
			return
		}
	}

	for _, subGraph := range g.subgraphs {
		subGraph.VisitNodes(callback)
	}
}

// FindNodeByID return node by id
func (g *Graph) FindNodeByID(id string) (foundNode *Node, found bool) {
	g.VisitNodes(func(node *Node) (done bool) {
		if node.id == id {
			found = true
			foundNode = node
			return true
		}
		return false
	})
	return
}

// FindNodes returns all nodes recursively
func (g *Graph) FindNodes() (nodes []*Node) {
	var foundNodes []*Node
	g.VisitNodes(func(node *Node) (done bool) {
		foundNodes = append(foundNodes, node)
		return false
	})
	return foundNodes
}
