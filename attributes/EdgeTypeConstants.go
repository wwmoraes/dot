package attributes

// EdgeType is a supported edge notation
type EdgeType string

const (
	// EdgeTypeDirected is the notation used on directed graph edges
	EdgeTypeDirected EdgeType = "->"
	// EdgeTypeUndirected is the notation used on undirected graph edges
	EdgeTypeUndirected EdgeType = "--"
)
