package dot

import (
	"bytes"
	"fmt"
	"io"

	"github.com/emicklei/dot/attributes"
)

// GraphData represents a dot graph with nodes and edges.
type GraphData struct {
	*attributes.Attributes
	id        string
	graphType GraphType
	strict    bool
	generator *UIDGenerator
	nodes     map[string]Node
	edgesFrom map[string][]Edge
	subgraphs map[string]Graph
	parent    Graph
	sameRank  map[string][]Node
	//
	nodeInitializer func(Node)
	edgeInitializer func(Edge)
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
		options.Type = GraphTypeDirected
	}

	if options.Type == GraphTypeSub && options.parent == nil {
		panic("cannot create subgraph without parent")
	} else if options.Type != GraphTypeSub && options.parent != nil {
		panic("cannot create graph with parent")
	}

	graph := &GraphData{
		id:              options.ID,
		Attributes:      attributes.NewAttributes(),
		graphType:       options.Type,
		strict:          options.Strict,
		generator:       NewUIDGenerator(24),
		nodes:           map[string]Node{},
		edgesFrom:       map[string][]Edge{},
		subgraphs:       map[string]Graph{},
		sameRank:        map[string][]Node{},
		nodeInitializer: options.NodeInitializer,
		edgeInitializer: options.EdgeInitializer,
	}

	return graph
}

// ID returns the immutable id
func (g *GraphData) ID() string {
	return g.id
}

func (g *GraphData) Subgraph(options *GraphOptions) Graph {
	if options == nil {
		options = &GraphOptions{}
	}

	// enforce subgraph type
	options.Type = GraphTypeSub

	// set parent
	options.parent = g

	sub := NewGraph(options)
	sub.SetAttributeString(attributes.KeyLabel, sub.ID())

	// save on parent with the generated ID
	g.subgraphs[sub.ID()] = sub

	return sub
}

// Label sets the "label" attribute value.
func (g *GraphData) Label(label string) Graph {
	g.SetAttribute("label", attributes.NewString(label))
	return g
}

// Root returns the top-level graph if this was a subgraph.
func (g *GraphData) Root() Graph {
	if g.parent == nil {
		return g
	}
	return g.parent.Root()
}

// FindSubgraph returns the subgraph of the graph or one from its parents.
func (g *GraphData) FindSubgraph(id string) (Graph, bool) {
	sub, ok := g.subgraphs[id]
	if !ok && g.parent != nil {
		return g.parent.FindSubgraph(id)
	}
	return sub, ok
}

// Subgraph returns the Graph with the given id ; creates one if absent.
// The label attribute is also set to the id ; use Label() to overwrite it.
//
// if id is "-", a randonly generated ID will be set
func (g *GraphData) SubgraphOld(id string, options ...GraphOption) Graph {
	sub, ok := g.subgraphs[id]
	if ok {
		return sub
	}

	sub = g.Subgraph(&GraphOptions{
		ID:              id,
		NodeInitializer: g.nodeInitializer,
		EdgeInitializer: g.edgeInitializer,
	})

	// for consistency with Node creation behavior.
	sub.SetAttributeString("label", id)
	return sub
}

func (g *GraphData) FindNode(id string) (Node, bool) {
	if n, ok := g.nodes[id]; ok {
		return n, ok
	}
	if g.parent == nil {
		return &NodeData{id: "void"}, false
	}
	return g.parent.FindNode(id)
}

// NodeInitializer sets a function that is called (if not nil) when a Node is implicitly created.
func (g *GraphData) NodeInitializer(callback func(n Node)) {
	g.nodeInitializer = callback
}

// EdgeInitializer sets a function that is called (if not nil) when an Edge is implicitly created.
func (g *GraphData) EdgeInitializer(callback func(e Edge)) {
	g.edgeInitializer = callback
}

// Node returns the node created with this id or creates a new node if absent.
// The node will have a label attribute with the id as its value. Use Label() to overwrite this.
// This method can be used as both a constructor and accessor.
// not thread safe!
func (g *GraphData) Node(id string) Node {
	if n, ok := g.FindNode(id); ok {
		return n
	}
	if len(id) == 0 {
		id = g.generator.String()
	}
	n := &NodeData{
		id:         id,
		Attributes: attributes.NewAttributes(),
		graph:      g,
	}
	n.SetAttribute(attributes.KeyLabel, attributes.NewString(id))
	if g.nodeInitializer != nil {
		g.nodeInitializer(n)
	}
	// store local
	g.nodes[id] = n
	return n
}

// Edge creates a new edge between two nodes
func (g *GraphData) Edge(fromNode, toNode Node) Edge {
	return g.EdgeWithAttributes(fromNode, toNode, nil)
}

// Edge creates a new edge between two nodes, and set the given attributes
func (g *GraphData) EdgeWithAttributes(fromNode, toNode Node, attr attributes.Reader) Edge {
	// assume fromNode owner == toNode owner
	// if fromNode.Graph() != toNode.Graph() { // 1 or 2 are subgraphs
	// 	edgeOwner := commonParentOf(fromNode.Graph(), toNode.Graph())
	// 	if edgeOwner.ID() != g.ID() {
	// 		return edgeOwner.EdgeWithAttributes(fromNode, toNode, attr)
	// 	}
	// }

	e := &EdgeData{
		from:       fromNode,
		to:         toNode,
		id:         g.generator.String(),
		Attributes: attributes.NewAttributesFrom(attr),
		graph:      g}

	if g.edgeInitializer != nil {
		g.edgeInitializer(e)
	}

	g.edgesFrom[fromNode.ID()] = append(g.edgesFrom[fromNode.ID()], e)

	return e
}

// FindEdges finds all edges in the graph that go from the fromNode to the toNode.
// Otherwise, returns an empty slice.
func (g *GraphData) FindEdges(fromNode, toNode Node) (found []Edge) {
	// if fromNode.Graph() != toNode.Graph() {
	// 	edgeOwner := commonParentOf(fromNode.Graph(), toNode.Graph())
	// 	return edgeOwner.FindEdges(fromNode, toNode)
	// }
	found = make([]Edge, 0)
	if edges, ok := g.edgesFrom[fromNode.ID()]; ok {
		for _, e := range edges {
			if e.To().ID() == toNode.ID() {
				found = append(found, e)
			}
		}
	}
	return found
}

// func commonParentOf(one, two Graph) Graph {
// 	// TODO
// 	return one.Root()
// }

// AddToSameRank adds the given nodes to the specified rank group, forcing them to be rendered in the same row
func (g *GraphData) AddToSameRank(group string, nodes ...Node) {
	g.sameRank[group] = append(g.sameRank[group], nodes...)
}

// String returns the source in dot notation.
func (g *GraphData) String() string {
	b := new(bytes.Buffer)
	g.Write(b)
	return b.String()
}

func (g *GraphData) Write(w io.Writer) {
	g.IndentedWrite(NewIndentWriter(w))
}

// IndentedWrite write the graph to a writer using simple TAB indentation.
func (g *GraphData) IndentedWrite(w *IndentWriter) {
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
			fmt.Fprintf(w, `"%s"`, each.ID())
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
				fmt.Fprintf(w, `"%s"%s"%s"`, each.From().ID(), denoteEdge, each.To().ID())
				each.Write(w, true)
				fmt.Fprint(w, ";")
				w.NewLine()
			}
		}
		for _, nodes := range g.sameRank {
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
func (g *GraphData) VisitNodes(callback func(node Node) (done bool)) {
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
func (g *GraphData) FindNodeByID(id string) (foundNode Node, found bool) {
	g.VisitNodes(func(node Node) (done bool) {
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
func (g *GraphData) FindNodes() (nodes []Node) {
	var foundNodes []Node
	g.VisitNodes(func(node Node) (done bool) {
		foundNodes = append(foundNodes, node)
		return false
	})
	return foundNodes
}
