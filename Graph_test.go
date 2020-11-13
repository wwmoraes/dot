package dot

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path"
	"reflect"
	"testing"

	"github.com/wwmoraes/dot/attributes"
	"github.com/wwmoraes/dot/dottest"
)

// TestGraphBehavior tests all components with real use cases
func TestGraphBehavior(t *testing.T) {
	t.Run("default graph", func(t *testing.T) {
		graph := NewGraph(nil)

		graph.Node("n1").Edge(graph.Node("n2")).SetAttributeString(attributes.KeyLabel, "uses")

		expected := `digraph {"n1";"n2";"n1"->"n2"[label="uses"];}`

		if got, want := dottest.MustGetFlattenSerializableString(t, graph), dottest.Flatten(t, expected); got != want {
			t.Errorf("got [\n%v\n]want [\n%v\n]", got, want)
		}
	})
	t.Run("directed graph", func(t *testing.T) {
		graph := NewGraph(&GraphOptions{
			Type: GraphTypeDirected,
		})

		graph.Node("n1").Edge(graph.Node("n2")).SetAttributeString(attributes.KeyLabel, "uses")

		expected := `digraph {"n1";"n2";"n1"->"n2"[label="uses"];}`

		if got, want := dottest.MustGetFlattenSerializableString(t, graph), dottest.Flatten(t, expected); got != want {
			t.Errorf("got [\n%v\n]want [\n%v\n]", got, want)
		}
	})
	t.Run("undirected graph", func(t *testing.T) {
		graph := NewGraph(&GraphOptions{
			Type: GraphTypeUndirected,
		})

		graph.Node("n1").Edge(graph.Node("n2")).SetAttributeString(attributes.KeyLabel, "uses")

		expected := `graph {"n1";"n2";"n1"--"n2"[label="uses"];}`

		if got, want := dottest.MustGetFlattenSerializableString(t, graph), dottest.Flatten(t, expected); got != want {
			t.Errorf("got [\n%v\n]want [\n%v\n]", got, want)
		}
	})
	t.Run("subgraphs", func(t *testing.T) {
		graph := NewGraph(nil)
		subGraph := graph.Subgraph(nil)
		subSubGraph := subGraph.Subgraph(nil)

		n1 := graph.Node("n1")
		n1.Edge(subGraph.Node("n2")).Edge(subSubGraph.Node("n3")).Edge(n1)
		subSubGraph.Node("n4").Edge(subSubGraph.Node("n3"))

		expected := `digraph {subgraph {subgraph {"n3";"n4";"n4"->"n3";}"n2";}"n1";"n1"->"n2";"n2"->"n3";"n3"->"n1";}`

		if got, want := dottest.MustGetFlattenSerializableString(t, graph), dottest.Flatten(t, expected); got != want {
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
			want: `digraph {}`,
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
					Type: GraphTypeDirected,
				},
			},
			want: `digraph {}`,
		},
		{
			name: "empty named directed graph",
			args: args{
				options: &GraphOptions{
					ID:   "test",
					Type: GraphTypeDirected,
				},
			},
			want: `digraph "test" {}`,
		},
		{
			name: "empty undirected graph",
			args: args{
				options: &GraphOptions{
					Type: GraphTypeUndirected,
				},
			},
			want: `graph {}`,
		},
		{
			name: "empty named undirected graph",
			args: args{
				options: &GraphOptions{
					ID:   "test",
					Type: GraphTypeUndirected,
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
			graph := NewGraph(tt.args.options)
			if got := dottest.MustGetFlattenSerializableString(t, graph); !reflect.DeepEqual(got, tt.want) {
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

		if got, want := dottest.MustGetFlattenSerializableString(t, graph), `digraph {"n1"[class="test-class"];}`; got != want {
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

		if got, want := dottest.MustGetFlattenSerializableString(t, graph), `digraph {"n1";"n2";"n1"->"n2"[class="test-class"];}`; got != want {
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
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
	t.Run("find no un-existent subgraph from another subgraph", func(t *testing.T) {
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
				Type: GraphTypeDirected,
			},
			want: func(graph Graph) string {
				return fmt.Sprintf(`digraph "%s" {}`, graph.ID())
			},
		},
		{
			name: "empty randomly-named undirected graph",
			options: &GraphOptions{
				ID:   "-",
				Type: GraphTypeUndirected,
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

			if got := dottest.MustGetFlattenSerializableString(t, graph); !reflect.DeepEqual(got, want) {
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
				Type: GraphTypeSub,
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
			want:    `digraph {subgraph {}}`,
		},
		{
			name: "empty named subgraph",
			options: &GraphOptions{
				ID: "test-sub",
			},
			want: `digraph {subgraph "test-sub" {}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := NewGraph(nil)
			graph.Subgraph(tt.options)
			if got := dottest.MustGetFlattenSerializableString(t, graph); got != tt.want {
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

		if got, want := dottest.MustGetFlattenSerializableString(t, graph), fmt.Sprintf(`digraph {subgraph "%s" {}}`, subGraph.ID()); got != want {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
}

func TestEmpty(t *testing.T) {
	di := NewGraph(nil)
	if got, want := dottest.MustGetFlattenSerializableString(t, di), `digraph {}`; got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
	}
	di2 := NewGraph(&GraphOptions{
		ID: "test",
	})
	if got, want := dottest.MustGetFlattenSerializableString(t, di2), `digraph "test" {}`; got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
	}
	di3 := NewGraph(&GraphOptions{
		ID: "-",
	})
	if di3.ID() == "-" {
		t.Error("got dash id instead of randomly generated one")
	}
	if got, want := dottest.MustGetFlattenSerializableString(t, di3), fmt.Sprintf(`digraph "%s" {}`, di3.ID()); got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
	}
}

func TestStrict(t *testing.T) {
	// test strict directed
	{
		graph := NewGraph(&GraphOptions{
			Strict: true,
		})
		if got, want := dottest.MustGetFlattenSerializableString(t, graph), `strict digraph {}`; got != want {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	}
	// test strict undirected
	{
		graph := NewGraph(&GraphOptions{
			Strict: true,
			Type:   GraphTypeUndirected,
		})
		if got, want := dottest.MustGetFlattenSerializableString(t, graph), `strict graph {}`; got != want {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	}
}

func TestEmptyWithIDAndAttributes(t *testing.T) {
	di := NewGraph(nil)
	di.SetAttribute(attributes.KeyStyle, attributes.NewString("filled"))
	di.SetAttribute(attributes.KeyColor, attributes.NewString("lightgrey"))
	if got, want := dottest.MustGetFlattenSerializableString(t, di), `digraph {graph [color="lightgrey",style="filled"];}`; got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
	}
}

func TestEmptyWithHTMLLabel(t *testing.T) {
	di := NewGraph(nil)
	di.SetAttribute(attributes.KeyLabel, attributes.NewHTML("<B>Hi</B>"))
	if got, want := dottest.MustGetFlattenSerializableString(t, di), `digraph {graph [label=<<B>Hi</B>>];}`; got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
	}
}

func TestEmptyWithLiteralValueLabel(t *testing.T) {
	di := NewGraph(nil)
	di.SetAttribute(attributes.KeyLabel, attributes.NewLiteral(`"left-justified text\l"`))
	if got, want := dottest.MustGetFlattenSerializableString(t, di), `digraph {graph [label="left-justified text\l"];}`; got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
	}
}

func TestTwoConnectedNodes(t *testing.T) {
	di := NewGraph(nil)
	n1 := di.Node("A")
	n2 := di.Node("B")
	di.Edge(n1, n2)
	if got, want := dottest.MustGetFlattenSerializableString(t, di), fmt.Sprintf(`digraph {"%[1]s";"%[2]s";"%[1]s"->"%[2]s";}`, n1.ID(), n2.ID()); got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
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
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	}

	// test finding target edges
	{
		n3 := sub.Node("C")
		newEdge := di.Edge(n2, n3)
		want := []Edge{newEdge}
		got := edge.EdgesTo(n3)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	}
}

func TestUndirectedTwoConnectedNodes(t *testing.T) {
	di := NewGraph(&GraphOptions{
		Type: GraphTypeUndirected,
	})
	n1 := di.Node("A")
	n2 := di.Node("B")
	di.Edge(n1, n2)
	if got, want := dottest.MustGetFlattenSerializableString(t, di), fmt.Sprintf(`graph {"%[1]s";"%[2]s";"%[1]s"--"%[2]s";}`, n1.ID(), n2.ID()); got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
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
	if got, want := dottest.MustGetFlattenSerializableString(t, di), fmt.Sprintf(`digraph {subgraph "%s" {graph [style="filled"];}}`, sub.ID()); got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
	}
	sub.SetAttributeString(attributes.KeyLabel, "new-label")
	if got, want := dottest.MustGetFlattenSerializableString(t, di), fmt.Sprintf(`digraph {subgraph "%s" {graph [label="new-label",style="filled"];}}`, sub.ID()); got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
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
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
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

	if got, want := dottest.MustGetFlattenSerializableString(t, graph), fmt.Sprintf(`digraph {"%s"[label="test",shape="box"];}`, node.ID()); got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
	}

	// test extra node + finding inexistent edge
	{
		node2 := graph.Node("")
		node.EdgesTo(node2)

		want := []Edge{}
		got := node.EdgesTo(node2)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
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
	if got, want := dottest.MustGetFlattenSerializableString(t, di), `digraph {"e1";"e2";"e1"->"e2"[label="what"];}`; got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
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
	if got, want := dottest.MustGetFlattenSerializableString(t, di), `digraph {"bar";"foo1";"foo2";{rank=same;"foo1";"foo2";}"foo1"->"foo2";"foo1"->"bar";}`; got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
	}
}

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
	filePath := path.Join(t.TempDir(), "cluster.dot")
	if err := ioutil.WriteFile(filePath, []byte(dottest.MustGetSerializableString(t, di)), os.ModePerm); err != nil {
		t.Errorf("unable to write dot file: %w", err)
	}
}

func TestDeleteLabel(t *testing.T) {
	g := NewGraph(nil)
	n := g.Node("my-id")
	n.DeleteAttribute(attributes.KeyLabel)
	if got, want := dottest.MustGetFlattenSerializableString(t, g), `digraph {"my-id";}`; got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
	}
}

func TestGraph_FindNodeById_emptyGraph(t *testing.T) {
	di := NewGraph(nil)

	_, found := di.FindNodeByID("F")

	if got, want := found, false; got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
	}
}

func TestGraph_FindNodeById_multiNodeGraph(t *testing.T) {
	di := NewGraph(nil)
	di.Node("A")
	di.Node("B")

	node, found := di.FindNodeByID("A")

	if got, want := node.ID(), "A"; got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
	}

	if got, want := found, true; got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
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
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
	}

	if got, want := found, true; got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
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
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
	}
}

func TestLabelWithEscaping(t *testing.T) {
	di := NewGraph(nil)
	n := di.Node("without-linefeed")
	n.SetAttribute(attributes.KeyLabel, attributes.NewLiteral(`"with \l linefeed"`))
	if got, want := dottest.MustGetFlattenSerializableString(t, di), `digraph {"without-linefeed"[label="with \l linefeed"];}`; got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
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
		Type: GraphTypeUndirected,
	})
	n1 := di.Node("A")
	n2 := di.Node("A")
	if got, want := n1, n2; &n1 == &n2 {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestGraph_WriteTo(t *testing.T) {
	tests := []struct {
		name       string
		wantErr    error
		wantString string
	}{
		{
			name:       "zero data written",
			wantErr:    dottest.ErrLimit,
			wantString: "",
		},
		{
			name:       "partially written - strict preffix",
			wantErr:    dottest.ErrLimit,
			wantString: `strict `,
		},
		{
			name:       "partially written - graph type",
			wantErr:    dottest.ErrLimit,
			wantString: `strict digraph`,
		},
		{
			name:       "partially written - graph id",
			wantErr:    dottest.ErrLimit,
			wantString: `strict digraph "test-graph"`,
		},
		{
			name:       "partially written - global graph attributes group",
			wantErr:    dottest.ErrLimit,
			wantString: `strict digraph "test-graph" {`,
		},
		{
			name:       "partially written - global graph attributes group",
			wantErr:    dottest.ErrLimit,
			wantString: `strict digraph "test-graph" {graph `,
		},
		{
			name:       "partially written - global graph attributes",
			wantErr:    dottest.ErrLimit,
			wantString: `strict digraph "test-graph" {graph [label="test-graph"]`,
		},
		{
			name:       "partially written - global graph attributes separator",
			wantErr:    dottest.ErrLimit,
			wantString: `strict digraph "test-graph" {graph [label="test-graph"];`,
		},
		{
			name:       "partially written - subgraph start",
			wantErr:    dottest.ErrLimit,
			wantString: `strict digraph "test-graph" {graph [label="test-graph"];subgraph`,
		},
		{
			name:       "partially written - subgraph open block",
			wantErr:    dottest.ErrLimit,
			wantString: `strict digraph "test-graph" {graph [label="test-graph"];subgraph {`,
		},
		{
			name:       "partially written - subgraph first node ID",
			wantErr:    dottest.ErrLimit,
			wantString: `strict digraph "test-graph" {graph [label="test-graph"];subgraph {"n1"`,
		},
		{
			name:       "partially written - subgraph first node separator",
			wantErr:    dottest.ErrLimit,
			wantString: `strict digraph "test-graph" {graph [label="test-graph"];subgraph {"n1";`,
		},
		{
			name:       "partially written - subgraph second node ID",
			wantErr:    dottest.ErrLimit,
			wantString: `strict digraph "test-graph" {graph [label="test-graph"];subgraph {"n1";"n2"`,
		},
		{
			name:       "partially written - subgraph second node separator",
			wantErr:    dottest.ErrLimit,
			wantString: `strict digraph "test-graph" {graph [label="test-graph"];subgraph {"n1";"n2";`,
		},
		{
			name:       "partially written - subgraph anonymous same rank group open",
			wantErr:    dottest.ErrLimit,
			wantString: `strict digraph "test-graph" {graph [label="test-graph"];subgraph {"n1";"n2";{`,
		},
		{
			name:       "partially written - subgraph anonymous same rank group attribute",
			wantErr:    dottest.ErrLimit,
			wantString: `strict digraph "test-graph" {graph [label="test-graph"];subgraph {"n1";"n2";{rank=same;`,
		},
		{
			name:       "partially written - subgraph anonymous same rank group first node",
			wantErr:    dottest.ErrLimit,
			wantString: `strict digraph "test-graph" {graph [label="test-graph"];subgraph {"n1";"n2";{rank=same;"n1"`,
		},
		{
			name:       "partially written - subgraph anonymous same rank group node separator",
			wantErr:    dottest.ErrLimit,
			wantString: `strict digraph "test-graph" {graph [label="test-graph"];subgraph {"n1";"n2";{rank=same;"n1";`,
		},
		{
			name:       "partially written - subgraph anonymous same rank group second node",
			wantErr:    dottest.ErrLimit,
			wantString: `strict digraph "test-graph" {graph [label="test-graph"];subgraph {"n1";"n2";{rank=same;"n1";"n2"`,
		},
		{
			name:       "partially written - subgraph anonymous same rank group node separator",
			wantErr:    dottest.ErrLimit,
			wantString: `strict digraph "test-graph" {graph [label="test-graph"];subgraph {"n1";"n2";{rank=same;"n1";"n2";`,
		},
		{
			name:       "partially written - subgraph anonymous same rank group close",
			wantErr:    dottest.ErrLimit,
			wantString: `strict digraph "test-graph" {graph [label="test-graph"];subgraph {"n1";"n2";{rank=same;"n1";"n2";}`,
		},
		{
			name:       "fully written",
			wantString: `strict digraph "test-graph" {graph [label="test-graph"];subgraph {"n1";"n2";{rank=same;"n1";"n2";}"n1"->"n2"[label="subgraph-edge"];}"n1"[label="node1"];"n2"[label="node2"];"n1"->"n2"[label="graph-edge"];}`,
		},
	}

	graph := NewGraph(&GraphOptions{
		ID:     "test-graph",
		Strict: true,
	})
	graph.SetAttributeString("label", "test-graph")
	subGraph := graph.Subgraph(nil)
	subGraph.AddToSameRank("g1", subGraph.Node("n1"), subGraph.Node("n2"))
	subEdge := subGraph.Edge(subGraph.Node("n1"), subGraph.Node("n2"))
	subEdge.SetAttributeString("label", "subgraph-edge")
	edge := graph.Edge(graph.Node("n1"), graph.Node("n2"))
	graph.Node("n1").SetAttributeString("label", "node1")
	graph.Node("n2").SetAttributeString("label", "node2")
	edge.SetAttributeString("label", "graph-edge")

	for limit, tt := range tests {
		if limit == len(tests)-1 {
			limit = math.MaxInt32
		}
		t.Run(tt.name, func(t *testing.T) {
			wantN := int64(len(tt.wantString))
			dottest.TestByteWrite(t, graph, limit, tt.wantErr, wantN, tt.wantString)
		})
	}
}

func BenchmarkGraph_WriteTo(b *testing.B) {
	b.Run("empty graph to buffer", func(b *testing.B) {
		var err error
		var buf bytes.Buffer

		graph := NewGraph(nil)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, err = graph.WriteTo(&buf)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("empty graph to ioutil.Discard", func(b *testing.B) {
		var err error

		graph := NewGraph(nil)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, err = graph.WriteTo(ioutil.Discard)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
