package dot

import (
	"errors"

	"github.com/wwmoraes/dot/generators"
)

// ErrSubgraphWithoutParent means a subgraph cannot be created without a parent
var ErrSubgraphWithoutParent = errors.New("cannot create subgraph without parent")

// ErrNonSubgraphWithParent means a non-subgraph cannot be created with a parent
var ErrNonSubgraphWithParent = errors.New("cannot create [di]graph with parent")

// ErrGraphWithoutGenerator means a graph cannot be created without a generator
var ErrGraphWithoutGenerator = errors.New("cannot create a graph without an ID generator")

// ErrRootCluster means a cluster graph cannot be created at root level
var ErrRootCluster = errors.New("cannot create a root cluster graph")

// ErrInvalidParent means an invalid parent was provided
var ErrInvalidParent = errors.New("cannot create subgraph with invalid or nil parent")

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

// WithID sets the graph id, or generates one if id is "-"
func WithID(id string) GraphOptionFn {
	return func(options GraphOptions) error {
		if id == "-" {
			if options.Generator() == nil {
				return ErrGraphWithoutGenerator
			}

			id = options.Generator().String()
		}

		options.SetID(id)

		return nil
	}
}

// WithParent sets the graph parent
func WithParent(graph Graph) GraphOptionFn {
	return func(options GraphOptions) error {
		if graph == nil {
			return ErrInvalidParent
		}

		options.SetParent(graph)
		options.SetType(GraphTypeSub)

		return nil
	}
}

// WithType sets the graph type
func WithType(graphType GraphType) GraphOptionFn {
	return func(options GraphOptions) error {
		if graphType == GraphTypeSub {
			if options.Parent() == nil {
				return ErrSubgraphWithoutParent
			}

			options.SetType(graphType)

			return nil
		}

		if options.Parent() != nil {
			return ErrNonSubgraphWithParent
		}

		options.SetType(graphType)

		return nil
	}
}

// WithStrict sets the graph as strict
func WithStrict() GraphOptionFn {
	return func(options GraphOptions) error {
		if options.Type() != GraphTypeSub {
			options.SetStrict(true)
		}

		return nil
	}
}

// WithCluster sets the graph as a cluster
func WithCluster() GraphOptionFn {
	return func(options GraphOptions) error {
		if options.Type() != GraphTypeSub {
			return ErrRootCluster
		}

		options.SetCluster(true)

		return nil
	}
}

// WithGenerator sets the graph id generator
func WithGenerator(generator generators.IDGenerator) GraphOptionFn {
	return func(options GraphOptions) error {
		options.SetGenerator(generator)

		return nil
	}
}

// WithNodeInitializer sets the graph node initializer function
func WithNodeInitializer(nodeInitializerFn NodeInitializerFn) GraphOptionFn {
	return func(options GraphOptions) error {
		options.SetNodeInitializer(nodeInitializerFn)

		return nil
	}
}

// WithEdgeInitializer sets the graph edge initializer function
func WithEdgeInitializer(edgeInitializerFn EdgeInitializerFn) GraphOptionFn {
	return func(options GraphOptions) error {
		options.SetEdgeInitializer(edgeInitializerFn)

		return nil
	}
}
