package dot

import "errors"

// ErrSubgraphWithoutParent means a subgraph cannot be created without a parent
var ErrSubgraphWithoutParent = errors.New("cannot create subgraph without parent")

// ErrRootWithParent means a non-subgraph cannot be created with a parent
var ErrRootWithParent = errors.New("cannot create [di]graph with parent")

// ErrGraphWithoutGenerator means a graph cannot be created without a generator
var ErrGraphWithoutGenerator = errors.New("cannot create a graph without an ID generator")

// ErrRootAsCluster means a cluster graph cannot be created at root level
var ErrRootAsCluster = errors.New("cannot create a root cluster graph")

// ErrNilParent means a nil parent was provided
var ErrNilParent = errors.New("cannot create subgraph without a parent")
