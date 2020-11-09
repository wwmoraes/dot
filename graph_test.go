package dot

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/emicklei/dot/attributes"
)

func TestEmpty(t *testing.T) {
	di := NewGraph(Directed)
	if got, want := flatten(di.String()), `digraph "" {}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	di2 := NewGraph(Directed, &GraphIDOption{"test"})
	if got, want := flatten(di2.String()), `digraph "test" {}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	di3 := NewGraph(Directed, &GraphIDOption{"-"})
	if di3.id == "-" {
		t.Error("got dash id instead of randomly generated one")
	}
	if got, want := flatten(di3.String()), fmt.Sprintf(`digraph "%s" {}`, di3.id); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestStrict(t *testing.T) {
	// test strict directed
	{
		graph := NewGraph(StrictDirected)
		if got, want := flatten(graph.String()), `strict digraph "" {}`; got != want {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	}
	// test strict undirected
	{
		graph := NewGraph(StrictUndirected)
		if got, want := flatten(graph.String()), `strict graph "" {}`; got != want {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	}
}

func TestEmptyWithIDAndAttributes(t *testing.T) {
	di := NewGraph(Directed)
	di.SetAttribute(attributes.KeyStyle, attributes.NewString("filled"))
	di.SetAttribute(attributes.KeyColor, attributes.NewString("lightgrey"))
	if got, want := flatten(di.String()), `digraph "" {color="lightgrey";style="filled";}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestEmptyWithHTMLLabel(t *testing.T) {
	di := NewGraph(Directed)
	di.SetAttribute(attributes.KeyLabel, attributes.NewHTML("<B>Hi</B>"))
	if got, want := flatten(di.String()), `digraph "" {label=<<B>Hi</B>>;}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestEmptyWithLiteralValueLabel(t *testing.T) {
	di := NewGraph(Directed)
	di.SetAttribute(attributes.KeyLabel, attributes.NewLiteral(`"left-justified text\l"`))
	if got, want := flatten(di.String()), `digraph "" {label="left-justified text\l";}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestTwoConnectedNodes(t *testing.T) {
	di := NewGraph(Directed)
	n1 := di.Node("A")
	n2 := di.Node("B")
	di.Edge(n1, n2)
	if got, want := flatten(di.String()), fmt.Sprintf(`digraph "" {"%[1]s"[label="A"];"%[2]s"[label="B"];"%[1]s"->"%[2]s";}`, n1.id, n2.id); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestTwoConnectedNodesAcrossSubgraphs(t *testing.T) {
	di := NewGraph(Directed)
	n1 := di.Node("A")
	sub := di.Subgraph("my-sub")
	n2 := sub.Node("B")
	edge := di.Edge(n1, n2)
	edge.Label(attributes.NewString("cross-graph"))

	// test graph-level edge finding
	{
		want := []*Edge{edge}
		got := sub.FindEdges(n1, n2)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	}

	// test finding target edges
	{
		n3 := sub.Node("C")
		newEdge := di.Edge(n2, n3)
		want := []*Edge{newEdge}
		got := edge.EdgesTo(n3)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	}
}

func TestUndirectedTwoConnectedNodes(t *testing.T) {
	di := NewGraph(Undirected)
	n1 := di.Node("A")
	n2 := di.Node("B")
	di.Edge(n1, n2)
	if got, want := flatten(di.String()), fmt.Sprintf(`graph "" {"%[1]s"[label="A"];"%[2]s"[label="B"];"%[1]s"--"%[2]s";}`, n1.id, n2.id); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestGraph_FindEdges(t *testing.T) {
	di := NewGraph(Directed)
	n1 := di.Node("A")
	n2 := di.Node("B")
	want := []*Edge{di.Edge(n1, n2)}
	got := di.FindEdges(n1, n2)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TestGraph.FindEdges() = %v, want %v", got, want)
	}
}

func TestSubgraph(t *testing.T) {
	di := NewGraph(Directed)
	sub := di.Subgraph("test-id")
	sub.SetAttribute(attributes.KeyStyle, attributes.NewString("filled"))
	if got, want := flatten(di.String()), fmt.Sprintf(`digraph "" {subgraph "%s" {label="test-id";style="filled";}}`, sub.id); got != want {
		t.Errorf("got\n[%v] want\n[%v]", got, want)
	}
	sub.Label("new-label")
	if got, want := flatten(di.String()), fmt.Sprintf(`digraph "" {subgraph "%s" {label="new-label";style="filled";}}`, sub.id); got != want {
		t.Errorf("got\n[%v] want\n[%v]", got, want)
	}
	found, _ := di.FindSubgraph("test-id")
	if got, want := found, sub; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	subsub := sub.Subgraph("sub-test-id")
	found, _ = subsub.FindSubgraph("test-id")
	if got, want := found, sub; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}

	// test getting an existing subgraph
	gotGraph := sub.Subgraph("sub-test-id")
	if gotGraph != subsub {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", gotGraph, subsub)
	}
}

func TestSubgraphClusterOption(t *testing.T) {
	di := NewGraph(Directed)
	sub := di.Subgraph("s1", &ClusterOption{})
	if got, want := sub.id, "cluster_s1"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestNode(t *testing.T) {
	graph := NewGraph(Directed)
	node := graph.Node("")
	node.Label("test")
	node.Box()

	if node.ID() == "" {
		t.Error("got blank node id, expected a randonly generated")
		return
	}

	if got, want := flatten(graph.String()), fmt.Sprintf(`digraph "%s" {"%s"[label="test",shape="box"];}`, graph.id, node.ID()); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}

	// test extra node + finding inexistent edge
	{
		node2 := graph.Node("")
		node.EdgesTo(node2)

		want := []*Edge{}
		got := node.EdgesTo(node2)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	}
}

func TestEdgeLabel(t *testing.T) {
	di := NewGraph(Directed)
	n1 := di.Node("e1")
	n2 := di.Node("e2")
	n1.Edge(n2, "what")
	if got, want := flatten(di.String()), `digraph "" {"e1"[label="e1"];"e2"[label="e2"];"e1"->"e2"[label="what"];}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestSameRank(t *testing.T) {
	di := NewGraph(Directed)
	foo1 := di.Node("foo1")
	foo2 := di.Node("foo2")
	bar := di.Node("bar")
	foo1.Edge(foo2)
	foo1.Edge(bar)
	di.AddToSameRank("top-row", foo1, foo2)
	if got, want := flatten(di.String()), `digraph "" {"bar"[label="bar"];"foo1"[label="foo1"];"foo2"[label="foo2"];"foo1"->"foo2";"foo1"->"bar";{rank=same; "foo1";"foo2";};}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

// dot -Tpng cluster.dot > cluster.png && open cluster.png
func TestCluster(t *testing.T) {
	di := NewGraph(Directed)
	outside := di.Node("Outside")
	clusterA := di.Subgraph("Cluster A", &ClusterOption{})
	insideOne := clusterA.Node("one")
	insideTwo := clusterA.Node("two")
	clusterB := di.Subgraph("Cluster B", &ClusterOption{})
	insideThree := clusterB.Node("three")
	insideFour := clusterB.Node("four")
	outside.Edge(insideFour).Edge(insideOne).Edge(insideTwo).Edge(insideThree).Edge(outside)
	if err := ioutil.WriteFile("doc/cluster.dot", []byte(di.String()), os.ModePerm); err != nil {
		t.Errorf("unable to write dot file: %w", err)
	}
}

// remove tabs and newlines and spaces
func flatten(s string) string {
	return strings.Replace((strings.Replace(s, "\n", "", -1)), "\t", "", -1)
}

func TestDeleteLabel(t *testing.T) {
	g := NewGraph()
	n := g.Node("my-id")
	n.DeleteAttribute(attributes.KeyLabel)
	if got, want := flatten(g.String()), `digraph "" {"my-id";}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestGraph_FindNodeById_emptyGraph(t *testing.T) {
	di := NewGraph(Directed)

	_, found := di.FindNodeByID("F")

	if got, want := found, false; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestGraph_FindNodeById_multiNodeGraph(t *testing.T) {
	di := NewGraph(Directed)
	di.Node("A")
	di.Node("B")

	node, found := di.FindNodeByID("A")

	if got, want := node.ID(), "A"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}

	if got, want := found, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestGraph_FindNodeById_multiNodesInSubGraphs(t *testing.T) {
	di := NewGraph(Directed)
	di.Node("A")
	di.Node("B")
	sub := di.Subgraph("new subgraph")
	sub.Node("C")

	node, found := di.FindNodeByID("C")

	if got, want := node.ID(), "C"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}

	if got, want := found, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestGraph_FindNodes_multiNodesInSubGraphs(t *testing.T) {
	di := NewGraph(Directed)
	di.Node("A")
	di.Node("B")
	sub := di.Subgraph("new subgraph")
	sub.Node("C")

	nodes := di.FindNodes()

	if got, want := len(nodes), 3; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestLabelWithEscaping(t *testing.T) {
	di := NewGraph(Directed)
	n := di.Node("without-linefeed")
	n.SetAttribute(attributes.KeyLabel, attributes.NewLiteral(`"with \l linefeed"`))
	if got, want := flatten(di.String()), `digraph "" {"without-linefeed"[label="with \l linefeed"];}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestGraphNodeInitializer(t *testing.T) {
	di := NewGraph(Directed)
	di.NodeInitializer(func(n *Node) {
		n.SetAttribute(attributes.KeyLabel, attributes.NewString("test"))
	})
	n := di.Node("A")
	if got, want := n.GetAttribute(attributes.KeyLabel).(*attributes.String), attributes.NewString("test"); !reflect.DeepEqual(got, want) {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}

func TestGraphEdgeInitializer(t *testing.T) {
	di := NewGraph(Directed)
	di.EdgeInitializer(func(e *Edge) {
		e.SetAttribute(attributes.KeyLabel, attributes.NewString("test"))
	})
	e := di.Node("A").Edge(di.Node("B"))
	if got, want := e.GetAttribute(attributes.KeyLabel).(*attributes.String), attributes.NewString("test"); !reflect.DeepEqual(got, want) {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}

func TestGraphCreateNodeOnce(t *testing.T) {
	di := NewGraph(Undirected)
	n1 := di.Node("A")
	n2 := di.Node("A")
	if got, want := n1, n2; &n1 == &n2 {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}
