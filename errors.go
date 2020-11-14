package dot

import "errors"

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
