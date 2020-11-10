package attributes

// GraphType graphviz graph types
type GraphType string

var (
	// GraphTypeUndirected node edges will have no direction, i.e. a layout engine
	// can freely organize the graph and draw glyphless edges between the nodes
	GraphTypeUndirected GraphType = "graph"
	// GraphTypeDirected node edges will have direction, i.e. a layout engine will
	// organize the graph and draw glyph edges denoting its orientation
	GraphTypeDirected GraphType = "digraph"
	// GraphTypeSub can be used to group objects, serve as a context for attributes
	// or draw its contents as a independent cluster, on supported layout engines
	GraphTypeSub GraphType = "subgraph"
)
