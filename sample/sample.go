package main

import (
	"io"
	"log"
	"os"

	"github.com/wwmoraes/dot"
	"github.com/wwmoraes/dot/attributes"
	"github.com/wwmoraes/dot/formatters"
)

func writeToWith(graph dot.Graph, fileName string, writerFunctors ...func(io.Writer) io.Writer) {
	log.Printf("trying to open %s file for write plain dot...\n", fileName)
	fd, err := os.Create(fileName)

	if err == nil {
		defer fd.Close()

		var writer io.Writer = fd
		for _, writerFunctor := range writerFunctors {
			writer = writerFunctor(writer)
		}

		log.Printf("writing graph to %s...\n", fileName)
		_, err := graph.WriteTo(writer)
		if err == nil {
			log.Printf("%s written successfully!\n", fileName)
		}
	}
}

func main() {
	log.Println("creating graph instance...")
	rootGraph, _ := dot.NewGraph()

	log.Println("creating outside node...")
	outsideGraph := rootGraph.Node("Outside")

	log.Println("creating cluster subgraph A...")
	clusterA, _ := rootGraph.Subgraph(
		dot.WithID("A"),
		dot.WithCluster(),
	)

	clusterA.SetAttributeString(attributes.KeyLabel, "Cluster A")

	log.Println("creating node one...")
	insideOne := clusterA.Node("one")

	log.Println("creating node two...")
	insideTwo := clusterA.Node("two")

	log.Println("creating cluster subgraph B...")
	clusterB, _ := rootGraph.Subgraph(
		dot.WithID("B"),
		dot.WithCluster(),
	)

	clusterB.SetAttributeString(attributes.KeyLabel, "Cluster B")

	log.Println("creating node three...")
	insideThree := clusterB.Node("three")

	log.Println("creating node four...")
	insideFour := clusterB.Node("four")

	log.Println("creating edges...")
	outsideGraph.Edge(insideFour).Edge(insideOne).Edge(insideTwo).Edge(insideThree).Edge(outsideGraph)

	writeToWith(rootGraph, "plain.dot")
	writeToWith(rootGraph, "pretty.dot", formatters.NewPrettyWriter)
}
