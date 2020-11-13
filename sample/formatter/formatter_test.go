package formatter

import (
	"io"
	"io/ioutil"
	"testing"

	"github.com/wwmoraes/dot"
	"github.com/wwmoraes/dot/attributes"
	"github.com/wwmoraes/dot/formatters"
)

func SetupGraph(tb testing.TB) dot.Graph {
	tb.Helper()

	rootGraph := dot.NewGraph(nil)
	outsideGraph := rootGraph.Node("Outside")
	clusterA := rootGraph.Subgraph(&dot.GraphOptions{
		ID:      "A",
		Cluster: true,
	})
	clusterA.SetAttributeString(attributes.KeyLabel, "Cluster A")
	insideOne := clusterA.Node("one")
	insideTwo := clusterA.Node("two")
	clusterA.AddToSameRank(
		"test",
		clusterA.Node("one"),
		clusterA.Node("two"),
	)
	clusterB := rootGraph.Subgraph(&dot.GraphOptions{
		ID:      "B",
		Cluster: true,
	})
	clusterB.SetAttributeString(attributes.KeyLabel, "Cluster B")
	insideThree := clusterB.Node("three")
	insideFour := clusterB.Node("four")
	outsideGraph.Edge(insideFour).Edge(insideOne).Edge(insideTwo).Edge(insideThree).Edge(outsideGraph)

	return rootGraph
}

func BenchmarkSample_WriteTo(b *testing.B) {
	type args struct {
		writer io.Writer
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "direct",
			args: args{
				writer: ioutil.Discard,
			},
		},
		{
			name: "PrettyWriter",
			args: args{
				writer: formatters.NewPrettyWriter(ioutil.Discard),
			},
		},
	}

	for _, bb := range tests {
		b.Run(bb.name, func(b *testing.B) {
			var err error

			graph := SetupGraph(b)

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_, err = graph.WriteTo(bb.args.writer)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
