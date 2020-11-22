package dot

import (
	"fmt"
	"strings"

	"github.com/wwmoraes/dot/generators"
)

// graphOptionsData contains the parameters used for graph creation
type graphOptionsData struct {
	parent Graph
	// Generator is used to create IDs for objects to prevent colision
	generator generators.IDGenerator
	// ID immutable id
	id string
	// Type the graph type (directed, undirected or sub)
	graphType GraphType
	// Cluster forbids the creation of multi-edges i.e.:
	//
	// on directed graphs, only one one edge between a given pair of head and tail nodes is allowed
	//
	// on undirected graphs, only one edge between the same two nodes is allowed
	strict bool
	// Cluster denotes if the graph is the special cluster subgraph, whose name
	// starts with "cluster_"
	cluster bool
	// NodeInitializer applies defaults to newly created nodes
	nodeInitializer NodeInitializerFn
	// EdgeInitializer applies defaults to newly created edges
	edgeInitializer EdgeInitializerFn
}

// NewGraphOptions creates a new GraphOptions value that has all values
// initialized properly
func NewGraphOptions(optionsFn ...GraphOptionFn) (options GraphOptions, err error) {
	options = &graphOptionsData{
		generator: generators.NewRandTimeIDGenerator(24),
		graphType: GraphTypeDirected,
	}

	for _, graphOptionFn := range optionsFn {
		err = graphOptionFn(options)
		if err != nil {
			return nil, err
		}
	}

	return options, err
}

// SetID changes the id
func (options *graphOptionsData) SetID(id string) {
	options.id = id

	hasClusterPreffix := strings.HasPrefix(options.id, "cluster_")

	if options.cluster && !hasClusterPreffix {
		options.id = fmt.Sprintf("cluster_%s", options.id)
	} else if !options.cluster && hasClusterPreffix {
		options.id = strings.TrimPrefix(options.id, "cluster_")
	}
}

// ID returns the id
func (options *graphOptionsData) ID() string {
	return options.id
}

// SetParent changes the parent graph
func (options *graphOptionsData) SetParent(graph Graph) {
	options.parent = graph
	options.graphType = GraphTypeSub
}

// Parent return the parent graph
func (options *graphOptionsData) Parent() Graph {
	return options.parent
}

// SetType changes the graph type
func (options *graphOptionsData) SetType(graphType GraphType) {
	options.graphType = graphType
}

// Type returns the graph type
func (options *graphOptionsData) Type() GraphType {
	return options.graphType
}

// SetStrict changes the strict flag
func (options *graphOptionsData) SetStrict(strict bool) {
	options.strict = strict
}

// Strict returns the strict flag
func (options *graphOptionsData) Strict() bool {
	return options.strict
}

// SetGenerator changes the ID generator
func (options *graphOptionsData) SetGenerator(generator generators.IDGenerator) {
	options.generator = generator
}

// Generator return the ID generator
func (options *graphOptionsData) Generator() generators.IDGenerator {
	return options.generator
}

// SetNodeInitializer changes the node initializer function
func (options *graphOptionsData) SetNodeInitializer(nodeInitializer NodeInitializerFn) {
	options.nodeInitializer = nodeInitializer
}

// NodeInitializer returns the node initializer function
func (options *graphOptionsData) NodeInitializer() NodeInitializerFn {
	return options.nodeInitializer
}

// SetEdgeInitializer changes the node initializer function
func (options *graphOptionsData) SetEdgeInitializer(edgeInitializer EdgeInitializerFn) {
	options.edgeInitializer = edgeInitializer
}

// EdgeInitializer returns the node initializer function
func (options *graphOptionsData) EdgeInitializer() EdgeInitializerFn {
	return options.edgeInitializer
}

// SetCluster changes the cluster flag
func (options *graphOptionsData) SetCluster(cluster bool) {
	options.cluster = cluster

	hasClusterPreffix := strings.HasPrefix(options.id, "cluster_")

	if options.cluster && !hasClusterPreffix {
		options.id = fmt.Sprintf("cluster_%s", options.id)
	}
}

// Cluster returns if the graph is a cluster
func (options *graphOptionsData) Cluster() bool {
	return options.cluster
}
