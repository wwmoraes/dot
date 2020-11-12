package main

import (
	"log"
	"os"

	"github.com/wwmoraes/dot"
	"github.com/wwmoraes/dot/attributes"
	"github.com/wwmoraes/dot/formatters"
)

func main() {
	log.Println("creating graph instance...")
	rootGraph := dot.NewGraph(nil)

	log.Println("creating outside node...")
	outsideGraph := rootGraph.Node("Outside")

	log.Println("creating cluster subgraph A...")
	clusterA := rootGraph.Subgraph(&dot.GraphOptions{
		ID:      "A",
		Cluster: true,
	})
	clusterA.SetAttributeString(attributes.KeyLabel, "Cluster A")

	log.Println("creating node one...")
	insideOne := clusterA.Node("one")

	log.Println("creating node two...")
	insideTwo := clusterA.Node("two")

	log.Println("creating cluster subgraph B...")
	clusterB := rootGraph.Subgraph(&dot.GraphOptions{
		ID:      "B",
		Cluster: true,
	})
	clusterB.SetAttributeString(attributes.KeyLabel, "Cluster B")

	log.Println("creating node three...")
	insideThree := clusterB.Node("three")

	log.Println("creating node four...")
	insideFour := clusterB.Node("four")

	log.Println("creating edges...")
	outsideGraph.Edge(insideFour).Edge(insideOne).Edge(insideTwo).Edge(insideThree).Edge(outsideGraph)

	log.Println("trying to open cluster.dot file for write...")
	fd, err := os.Create("cluster.dot")
	if err == nil {
		log.Println("writing graph to cluster.dot...")
		rootGraph.Write(formatters.NewPrettyWriter(fd))
		log.Println("done!")
	}
}
