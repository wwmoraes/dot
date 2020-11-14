package dot

import (
	"github.com/wwmoraes/dot/generators"
)

// GraphOptions is implemented by values used to initialize graphs
type GraphOptions interface {
	// SetID changes the id
	SetID(string)
	// ID returns the id
	ID() string
	// SetParent changes the parent graph
	SetParent(Graph)
	// Parent return the parent graph
	Parent() Graph
	// SetType changes the graph type
	SetType(GraphType)
	// Type returns the graph type
	Type() GraphType
	// SetStrict changes the strict flag
	SetStrict(bool)
	// Strict returns the strict flag
	Strict() bool
	// SetGenerator changes the ID generator
	SetGenerator(generators.IDGenerator)
	// Generator return the ID generator
	Generator() generators.IDGenerator
	// SetNodeInitializer changes the node initializer function
	SetNodeInitializer(NodeInitializerFn)
	// NodeInitializer returns the node initializer function
	NodeInitializer() NodeInitializerFn
	// SetEdgeInitializer changes the node initializer function
	SetEdgeInitializer(EdgeInitializerFn)
	// EdgeInitializer returns the node initializer function
	EdgeInitializer() EdgeInitializerFn
	// SetCluster changes the cluster flag
	SetCluster(bool)
}
