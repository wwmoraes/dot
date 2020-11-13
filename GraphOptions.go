package dot

import (
	"fmt"

	"github.com/wwmoraes/dot/attributes"
	"github.com/wwmoraes/dot/generators"
)

// GraphOptions contains the parameters used for graph creation
type GraphOptions struct {
	parent Graph
	// Generator is used to create IDs for objects to prevent colision
	Generator generators.IDGenerator
	// ID immutable id
	ID string
	// Type the graph type (directed, undirected or sub)
	Type attributes.GraphType
	// Cluster forbids the creation of multi-edges i.e.:
	//
	// on directed graphs, only one one edge between a given pair of head and tail nodes is allowed
	//
	// on undirected graphs, only one edge between the same two nodes is allowed
	Strict bool
	// Cluster denotes if the graph is the special cluster subgraph, whose name
	// starts with "cluster_"
	Cluster bool
	// NodeInitializer applies defaults to newly created nodes
	NodeInitializer func(Node)
	// EdgeInitializer applies defaults to newly created edges
	EdgeInitializer func(StyledEdge)
}

// parseGraphOptions ensures the options are valid and initialized properly
func parseGraphOptions(options *GraphOptions) *GraphOptions {
	generator := generators.NewRandTimeIDGenerator(24)

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

	if options.Generator == nil {
		options.Generator = generators.NewRandTimeIDGenerator(24)
	}

	return options
}
