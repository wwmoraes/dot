package dot

import (
	"github.com/wwmoraes/dot/generators"
)

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
			return ErrNilParent
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
			return ErrRootWithParent
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
			return ErrRootAsCluster
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
