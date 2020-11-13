package main

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/wwmoraes/dot"
	"github.com/wwmoraes/dot/attributes"
	"github.com/wwmoraes/dot/dottest"
	"github.com/wwmoraes/dot/formatters"
)

const dotFileName string = "sample.dot"

const expectedOutput string = `
digraph {
  subgraph "cluster_A" {
    graph [label="Cluster A"];
    "one";
    "two";
    {rank=same;"one";"two";}
  }
  subgraph "cluster_B" {
    graph [label="Cluster B"];
    "four";
    "three";
  }
  "Outside";
  "Outside"->"four";
  "four"->"one";
  "one"->"two";
  "three"->"Outside";
  "two"->"three";
}
`

func TestSample(t *testing.T) {
	if err := os.Remove(dotFileName); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			t.Fatal(err)
		}
	}

	main()

	fileInfo, err := os.Stat("sample.dot")
	if err != nil {
		t.Fatal(err)
	}

	if !fileInfo.Mode().IsRegular() {
		t.Fatalf("%s is not a regular file", dotFileName)
	}

	fileData, err := ioutil.ReadFile("sample.dot")
	if err != nil {
		t.Fatal(err)
	}

	if got, want := dottest.Flatten(t, string(fileData)), dottest.Flatten(t, expectedOutput); got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
	}
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
			name: "IndentedWriter",
			args: args{
				writer: formatters.NewIndentedWriter(ioutil.Discard),
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

			graph := setupGraph(b)

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

func setupGraph(b *testing.B) dot.Graph {
	b.Helper()

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
