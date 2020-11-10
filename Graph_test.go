package dot

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/wwmoraes/dot/attributes"
)

// TestGraphBehavior tests all components with real use cases
func TestGraphBehavior(t *testing.T) {
	t.Run("default graph", func(t *testing.T) {
		graph := NewGraph(nil)

		graph.Node("n1").Edge(graph.Node("n2")).SetAttributeString(attributes.KeyLabel, "uses")

		expected := `digraph "" {"n1"[label="n1"];"n2"[label="n2"];"n1"->"n2"[label="uses"];}`

		if got, want := flatten(graph.String()), flatten(expected); got != want {
			t.Errorf("got [\n%v\n]want [\n%v\n]", got, want)
		}
	})
	t.Run("directed graph", func(t *testing.T) {
		graph := NewGraph(&GraphOptions{
			Type: attributes.GraphTypeDirected,
		})

		graph.Node("n1").Edge(graph.Node("n2")).SetAttributeString(attributes.KeyLabel, "uses")

		expected := `digraph "" {"n1"[label="n1"];"n2"[label="n2"];"n1"->"n2"[label="uses"];}`

		if got, want := flatten(graph.String()), flatten(expected); got != want {
			t.Errorf("got [\n%v\n]want [\n%v\n]", got, want)
		}
	})
	t.Run("undirected graph", func(t *testing.T) {
		graph := NewGraph(&GraphOptions{
			Type: attributes.GraphTypeUndirected,
		})

		graph.Node("n1").Edge(graph.Node("n2")).SetAttributeString(attributes.KeyLabel, "uses")

		expected := `graph "" {"n1"[label="n1"];"n2"[label="n2"];"n1"--"n2"[label="uses"];}`

		if got, want := flatten(graph.String()), flatten(expected); got != want {
			t.Errorf("got [\n%v\n]want [\n%v\n]", got, want)
		}
	})
	t.Run("subgraphs", func(t *testing.T) {
		graph := NewGraph(nil)
		subGraph := graph.Subgraph(nil)
		subSubGraph := subGraph.Subgraph(nil)

		n1 := graph.Node("n1")
		n1.Edge(subGraph.Node("n2")).Edge(subSubGraph.Node("n3")).Edge(n1)

		expected := `digraph "" {subgraph "" {subgraph "" {label="";"n3"[label="n3"];}label="";"n2"[label="n2"];}"n1"[label="n1"];"n1"->"n2";"n2"->"n3";"n3"->"n1";}`

		if got, want := flatten(graph.String()), flatten(expected); got != want {
			t.Errorf("got [\n%v\n]want [\n%v\n]", got, want)
		}
	})
}

// TestNewGraph tests NewGraph factory function with static outcome
func TestNewGraph(t *testing.T) {
	type args struct {
		options *GraphOptions
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty graph",
			args: args{
				options: nil,
			},
			want: `digraph "" {}`,
		},
		{
			name: "empty named graph",
			args: args{
				options: &GraphOptions{
					ID: "test",
				},
			},
			want: `digraph "test" {}`,
		},
		{
			name: "empty directed graph",
			args: args{
				options: &GraphOptions{
					Type: attributes.GraphTypeDirected,
				},
			},
			want: `digraph "" {}`,
		},
		{
			name: "empty named directed graph",
			args: args{
				options: &GraphOptions{
					ID:   "test",
					Type: attributes.GraphTypeDirected,
				},
			},
			want: `digraph "test" {}`,
		},
		{
			name: "empty undirected graph",
			args: args{
				options: &GraphOptions{
					Type: attributes.GraphTypeUndirected,
				},
			},
			want: `graph "" {}`,
		},
		{
			name: "empty named undirected graph",
			args: args{
				options: &GraphOptions{
					ID:   "test",
					Type: attributes.GraphTypeUndirected,
				},
			},
			want: `graph "test" {}`,
		},
		{
			name: "empty cluster graph",
			args: args{
				options: &GraphOptions{
					Cluster: true,
				},
			},
			want: `digraph "cluster_" {}`,
		},
		{
			name: "empty named cluster graph",
			args: args{
				options: &GraphOptions{
					ID:      "test",
					Cluster: true,
				},
			},
			want: `digraph "cluster_test" {}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flatten(NewGraph(tt.args.options).String()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_Initializers(t *testing.T) {
	t.Run("graph with node initializer", func(t *testing.T) {
		graph := NewGraph(&GraphOptions{
			NodeInitializer: func(nodeInstance Node) {
				nodeInstance.SetAttributeString(attributes.KeyClass, "test-class")
			},
		})

		graph.Node("n1")

		if got, want := flatten(graph.String()), `digraph "" {"n1"[class="test-class",label="n1"];}`; got != want {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
	t.Run("graph with edge initializer", func(t *testing.T) {
		graph := NewGraph(&GraphOptions{
			EdgeInitializer: func(edgeInstance StyledEdge) {
				edgeInstance.SetAttributeString(attributes.KeyClass, "test-class")
			},
		})

		graph.Node("n1").Edge(graph.Node("n2"))

		if got, want := flatten(graph.String()), `digraph "" {"n1"[label="n1"];"n2"[label="n2"];"n1"->"n2"[class="test-class"];}`; got != want {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
}

func TestGraph_FindSubgraph(t *testing.T) {
	t.Run("find existing subgraph from another subgraph", func(t *testing.T) {
		graph := NewGraph(&GraphOptions{
			ID: "root-graph",
		})
		sub1 := graph.Subgraph(&GraphOptions{ID: "subgraph-one"})
		sub2 := graph.Subgraph(&GraphOptions{ID: "subgraph-two"})

		got, found := sub1.FindSubgraph(sub2.ID())

		if !found {
			t.Error("subgraph not found as expected")
		}

		if want := sub2; !reflect.DeepEqual(got, want) {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	})
	t.Run("find no un-existant subgraph from another subgraph", func(t *testing.T) {
		graph := NewGraph(&GraphOptions{
			ID: "root-graph",
		})
		sub1 := graph.Subgraph(&GraphOptions{ID: "subgraph-one"})

		got, found := sub1.FindSubgraph("subgraph-two")

		if found {
			t.Error("subgraph was found, it wasn't expected")
		}

		if got != nil {
			t.Errorf("got [%v] want [%v]", got, nil)
		}
	})
}

// TestNewGraph_generatedID tests NewGraph factory option to generate unique IDs
func TestNewGraph_generatedID(t *testing.T) {
	tests := []struct {
		name    string
		options *GraphOptions
		want    func(graph Graph) string
	}{
		{
			name: "empty randomly-named graph",
			options: &GraphOptions{
				ID: "-",
			},
			want: func(graph Graph) string {
				return fmt.Sprintf(`digraph "%s" {}`, graph.ID())
			},
		},
		{
			name: "empty randomly-named directed graph",
			options: &GraphOptions{
				ID:   "-",
				Type: attributes.GraphTypeDirected,
			},
			want: func(graph Graph) string {
				return fmt.Sprintf(`digraph "%s" {}`, graph.ID())
			},
		},
		{
			name: "empty randomly-named undirected graph",
			options: &GraphOptions{
				ID:   "-",
				Type: attributes.GraphTypeUndirected,
			},
			want: func(graph Graph) string {
				return fmt.Sprintf(`graph "%s" {}`, graph.ID())
			},
		},
		{
			name: "empty randomly-named cluster directed graph",
			options: &GraphOptions{
				ID:      "-",
				Cluster: true,
			},
			want: func(graph Graph) string {
				return fmt.Sprintf(`digraph "%s" {}`, graph.ID())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.options.ID = "-"

			graph := NewGraph(tt.options)

			if graph.ID() == "" {
				t.Error("got empty ID instead of a random one")
			}

			if graph.ID() == "-" {
				t.Error("got dash ID instead of a random one")
			}

			want := tt.want(graph)

			if got := flatten(graph.String()); !reflect.DeepEqual(got, want) {
				t.Errorf("NewGraph() = %v, want %v", got, want)
			}
		})
	}
}

// TestNewGraph_invalid tests NewGraph factory invalid options
func TestNewGraph_invalid(t *testing.T) {
	tests := []struct {
		name    string
		options *GraphOptions
	}{
		{
			name: "subgraph without parent",
			options: &GraphOptions{
				Type: attributes.GraphTypeSub,
			},
		},
		{
			name: "non-subgraph with parent",
			options: &GraphOptions{
				parent: NewGraph(nil),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err == nil {
					t.Errorf("it should panic")
				}
			}()

			got := NewGraph(tt.options)

			if got != nil {
				t.Errorf("got [%v] want [%v]", got, nil)
			}
		})
	}
}

// TestGraph_Subgraph tests Graph.Subgraph factory
func TestGraph_Subgraph(t *testing.T) {
	tests := []struct {
		name    string
		options *GraphOptions
		want    string
	}{
		{
			name:    "empty anonymous subgraph",
			options: nil,
			want:    `digraph "" {subgraph "" {label="";}}`,
		},
		{
			name: "empty named subgraph",
			options: &GraphOptions{
				ID: "test-sub",
			},
			want: `digraph "" {subgraph "test-sub" {label="test-sub";}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := NewGraph(nil)
			graph.Subgraph(tt.options)
			if got := flatten(graph.String()); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}

	t.Run("empty randomly-named subgraph", func(t *testing.T) {
		graph := NewGraph(nil)
		subGraph := graph.Subgraph(&GraphOptions{
			ID: "-",
		})

		if subGraph.ID() == "" {
			t.Error("got empty ID instead of a random one")
		}

		if subGraph.ID() == "-" {
			t.Error("got dash ID instead of a random one")
		}

		if got, want := flatten(graph.String()), fmt.Sprintf(`digraph "" {subgraph "%[1]s" {label="%[1]s";}}`, subGraph.ID()); got != want {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	})
}

func TestEmpty(t *testing.T) {
	di := NewGraph(nil)
	if got, want := flatten(di.String()), `digraph "" {}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	di2 := NewGraph(&GraphOptions{
		ID: "test",
	})
	if got, want := flatten(di2.String()), `digraph "test" {}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	di3 := NewGraph(&GraphOptions{
		ID: "-",
	})
	if di3.ID() == "-" {
		t.Error("got dash id instead of randomly generated one")
	}
	if got, want := flatten(di3.String()), fmt.Sprintf(`digraph "%s" {}`, di3.ID()); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestStrict(t *testing.T) {
	// test strict directed
	{
		graph := NewGraph(&GraphOptions{
			Strict: true,
		})
		if got, want := flatten(graph.String()), `strict digraph "" {}`; got != want {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	}
	// test strict undirected
	{
		graph := NewGraph(&GraphOptions{
			Strict: true,
			Type:   attributes.GraphTypeUndirected,
		})
		if got, want := flatten(graph.String()), `strict graph "" {}`; got != want {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	}
}

func TestEmptyWithIDAndAttributes(t *testing.T) {
	di := NewGraph(nil)
	di.SetAttribute(attributes.KeyStyle, attributes.NewString("filled"))
	di.SetAttribute(attributes.KeyColor, attributes.NewString("lightgrey"))
	if got, want := flatten(di.String()), `digraph "" {color="lightgrey";style="filled";}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestEmptyWithHTMLLabel(t *testing.T) {
	di := NewGraph(nil)
	di.SetAttribute(attributes.KeyLabel, attributes.NewHTML("<B>Hi</B>"))
	if got, want := flatten(di.String()), `digraph "" {label=<<B>Hi</B>>;}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestEmptyWithLiteralValueLabel(t *testing.T) {
	di := NewGraph(nil)
	di.SetAttribute(attributes.KeyLabel, attributes.NewLiteral(`"left-justified text\l"`))
	if got, want := flatten(di.String()), `digraph "" {label="left-justified text\l";}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestTwoConnectedNodes(t *testing.T) {
	di := NewGraph(nil)
	n1 := di.Node("A")
	n2 := di.Node("B")
	di.Edge(n1, n2)
	if got, want := flatten(di.String()), fmt.Sprintf(`digraph "" {"%[1]s"[label="A"];"%[2]s"[label="B"];"%[1]s"->"%[2]s";}`, n1.ID(), n2.ID()); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestTwoConnectedNodesAcrossSubgraphs(t *testing.T) {
	di := NewGraph(nil)
	n1 := di.Node("A")
	sub := di.Subgraph(&GraphOptions{
		ID: "my-sub",
	})
	n2 := sub.Node("B")
	edge := di.Edge(n1, n2)
	edge.SetAttributeString(attributes.KeyLabel, "cross-graph")

	// test graph-level edge finding
	{
		want := []Edge{edge}
		got := di.FindEdges(n1, n2)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	}

	// test finding target edges
	{
		n3 := sub.Node("C")
		newEdge := di.Edge(n2, n3)
		want := []Edge{newEdge}
		got := edge.EdgesTo(n3)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	}
}

func TestUndirectedTwoConnectedNodes(t *testing.T) {
	di := NewGraph(&GraphOptions{
		Type: attributes.GraphTypeUndirected,
	})
	n1 := di.Node("A")
	n2 := di.Node("B")
	di.Edge(n1, n2)
	if got, want := flatten(di.String()), fmt.Sprintf(`graph "" {"%[1]s"[label="A"];"%[2]s"[label="B"];"%[1]s"--"%[2]s";}`, n1.ID(), n2.ID()); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestGraph_FindEdges(t *testing.T) {
	di := NewGraph(nil)
	n1 := di.Node("A")
	n2 := di.Node("B")
	want := []Edge{di.Edge(n1, n2)}
	got := di.FindEdges(n1, n2)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TestGraph.FindEdges() = %v, want %v", got, want)
	}
}

func TestSubgraph(t *testing.T) {
	di := NewGraph(nil)
	sub := di.Subgraph(&GraphOptions{
		ID: "test-id",
	})
	sub.SetAttributeString(attributes.KeyStyle, "filled")
	if got, want := flatten(di.String()), fmt.Sprintf(`digraph "" {subgraph "%s" {label="test-id";style="filled";}}`, sub.ID()); got != want {
		t.Errorf("got\n[%v] want\n[%v]", got, want)
	}
	sub.SetAttributeString(attributes.KeyLabel, "new-label")
	if got, want := flatten(di.String()), fmt.Sprintf(`digraph "" {subgraph "%s" {label="new-label";style="filled";}}`, sub.ID()); got != want {
		t.Errorf("got\n[%v] want\n[%v]", got, want)
	}
	foundGraph, _ := di.FindSubgraph("test-id")
	if got, want := foundGraph, sub; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
	subsub := sub.Subgraph(&GraphOptions{
		ID: "sub-test-id",
	})
	foundGraph, _ = di.FindSubgraph("test-id")
	if got, want := foundGraph, sub; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}

	// test getting an existing subgraph

	gotGraph, found := sub.FindSubgraph(subsub.ID())
	if !found {
		t.Errorf("%s not found, it was expected to be found", subsub.ID())
	}
	if gotGraph != subsub {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", gotGraph, subsub)
	}
}

func TestSubgraphClusterOption(t *testing.T) {
	di := NewGraph(nil)
	sub := di.Subgraph(&GraphOptions{
		ID:      "s1",
		Cluster: true,
	})
	if got, want := sub.ID(), "cluster_s1"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestNode(t *testing.T) {
	graph := NewGraph(nil)
	node := graph.Node("")
	node.SetAttributesString(attributes.MapString{
		attributes.KeyLabel: "test",
		attributes.KeyShape: "box",
	})

	if node.ID() == "" {
		t.Error("got blank node id, expected a randonly generated")
		return
	}

	if got, want := flatten(graph.String()), fmt.Sprintf(`digraph "%s" {"%s"[label="test",shape="box"];}`, graph.ID(), node.ID()); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}

	// test extra node + finding inexistent edge
	{
		node2 := graph.Node("")
		node.EdgesTo(node2)

		want := []Edge{}
		got := node.EdgesTo(node2)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	}
}

func TestEdgeLabel(t *testing.T) {
	di := NewGraph(nil)
	n1 := di.Node("e1")
	n2 := di.Node("e2")
	attr := attributes.NewAttributes()
	attr.SetAttributeString(attributes.KeyLabel, "what")
	n1.EdgeWithAttributes(n2, attr)
	if got, want := flatten(di.String()), `digraph "" {"e1"[label="e1"];"e2"[label="e2"];"e1"->"e2"[label="what"];}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestSameRank(t *testing.T) {
	di := NewGraph(nil)
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
	di := NewGraph(nil)
	outside := di.Node("Outside")
	clusterA := di.Subgraph(&GraphOptions{
		ID:      "Cluster A",
		Cluster: true,
	})
	clusterA.SetAttributeString(attributes.KeyLabel, "Cluster A")
	insideOne := clusterA.Node("one")
	insideTwo := clusterA.Node("two")
	clusterB := di.Subgraph(&GraphOptions{
		ID:      "Cluster B",
		Cluster: true,
	})
	clusterB.SetAttributeString(attributes.KeyLabel, "Cluster B")
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
	g := NewGraph(nil)
	n := g.Node("my-id")
	n.DeleteAttribute(attributes.KeyLabel)
	if got, want := flatten(g.String()), `digraph "" {"my-id";}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestGraph_FindNodeById_emptyGraph(t *testing.T) {
	di := NewGraph(nil)

	_, found := di.FindNodeByID("F")

	if got, want := found, false; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestGraph_FindNodeById_multiNodeGraph(t *testing.T) {
	di := NewGraph(nil)
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
	di := NewGraph(nil)
	di.Node("A")
	di.Node("B")
	sub := di.Subgraph(&GraphOptions{
		ID: "new subgraph",
	})
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
	di := NewGraph(nil)
	di.Node("A")
	di.Node("B")
	sub := di.Subgraph(&GraphOptions{
		ID: "new subgraph",
	})
	sub.Node("C")

	nodes := di.FindNodes()

	if got, want := len(nodes), 3; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestLabelWithEscaping(t *testing.T) {
	di := NewGraph(nil)
	n := di.Node("without-linefeed")
	n.SetAttribute(attributes.KeyLabel, attributes.NewLiteral(`"with \l linefeed"`))
	if got, want := flatten(di.String()), `digraph "" {"without-linefeed"[label="with \l linefeed"];}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestGraphNodeInitializer(t *testing.T) {
	di := NewGraph(&GraphOptions{
		NodeInitializer: func(n Node) {
			n.SetAttribute(attributes.KeyLabel, attributes.NewString("test"))
		},
	})
	n := di.Node("A")
	gotAttr, gotOk := n.GetAttribute(attributes.KeyLabel)
	if !gotOk {
		t.Error("attribute not found")
	}
	if got, want := gotAttr.(*attributes.String), attributes.NewString("test"); !reflect.DeepEqual(got, want) {
		t.Errorf("got [%v[1]:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestGraphEdgeInitializer(t *testing.T) {
	di := NewGraph(&GraphOptions{
		EdgeInitializer: func(e StyledEdge) {
			e.SetAttribute(attributes.KeyLabel, attributes.NewString("test"))
		},
	})
	e := di.Node("A").Edge(di.Node("B"))
	gotAttr, gotOk := e.GetAttribute(attributes.KeyLabel)
	if !gotOk {
		t.Error("attribute not found")
	}
	if got, want := gotAttr.(*attributes.String), attributes.NewString("test"); !reflect.DeepEqual(got, want) {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestGraphCreateNodeOnce(t *testing.T) {
	di := NewGraph(&GraphOptions{
		Type: attributes.GraphTypeUndirected,
	})
	n1 := di.Node("A")
	n2 := di.Node("A")
	if got, want := n1, n2; &n1 == &n2 {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}
