---
title: GraphOptionFn
description: functor that mutates graph options
---

# `GraphOptionFn`

Implemented by functor values that set properties on a `GraphOptions` instance,
which are then used by `Graph`s `NewGraph`/`Graph.Subgraph` to configure new
[sub]graphs.

A `nil` error is expected as a return value if the functor has successfully set
the options it is intended to.

These functors are usually returned by `With*` functions that are free to accept
any parameters to be used within the functor, e.g.

```go
func WithMyOptions(id string) GraphOptionFn {
  return func(options GraphOptions) error {
    // use the given id
    options.SetID(id)
    // always create as strict
    options.SetStrict(true)
    // use undirected (graph) instead of the default directed (digraph)
    options.SetType(GraphTypeUndirected)

    return nil
  }
}
```

then it can be used as

```go
graph, _ := dot.New(WithMyOptions("some-id"))
```

## implementation

```go
// GraphOptionFn is a functor that mutates graph options
type GraphOptionFn func(GraphOptions) error
```

## source

```go
--8<-- "Graph.go"
```
