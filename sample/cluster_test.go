package main

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/wwmoraes/dot"
	"github.com/wwmoraes/dot/attributes"
	"github.com/wwmoraes/dot/formatters"
)

const dotFileName string = "cluster.dot"

const expectedOutput string = `
digraph {
  subgraph "cluster_A" {
    graph [label="Cluster A"];
    "one";
    "two";
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

func flatten(s string, t *testing.T) string {
	t.Helper()
	return strings.Replace((strings.Replace(s, "\n", "", -1)), "  ", "", -1)
}

func TestSample_Cluster(t *testing.T) {
	if err := os.Remove(dotFileName); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			t.Fatal(err)
		}
	}

	main()

	fileInfo, err := os.Stat(dotFileName)
	if err != nil {
		t.Fatal(err)
	}

	if !fileInfo.Mode().IsRegular() {
		t.Fatalf("%s is not a regular file", dotFileName)
	}

	fileData, err := ioutil.ReadFile(dotFileName)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := flatten(string(fileData), t), flatten(expectedOutput, t); got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
	}
}

func BenchmarkSample_Write(b *testing.B) {
	b.Run("Direct", func(b *testing.B) {
		graph := setupGraph(b)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			graph.Write(ioutil.Discard)
		}
	})
	b.Run("IndentedWriter", func(b *testing.B) {
		graph := setupGraph(b)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			graph.Write(ioutil.Discard)
		}
	})
	b.Run("PrettyWriter", func(b *testing.B) {
		graph := setupGraph(b)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			graph.Write(formatters.NewPrettyWriter(ioutil.Discard))
		}
	})
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
