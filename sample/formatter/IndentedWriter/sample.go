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

	log.Println("grouping nodes one and two...")
	clusterA.AddToSameRank(
		"test",
		clusterA.Node("one"),
		clusterA.Node("two"),
	)

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

	log.Println("trying to open sample.dot file for write...")
	fd, err := os.Create("sample.dot")
	if err == nil {
		log.Println("writing graph to sample.dot...")
		_, err := rootGraph.WriteTo(formatters.NewIndentedWriter(fd))
		if err != nil {
			log.Fatal(err)
		}
		log.Println("done!")
	}
}
