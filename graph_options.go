package dot

// GraphOption functor to apply an option to a graph
type GraphOption interface {
	Apply(*Graph)
}

// ClusterOption preffixes a graph ID with cluster
type ClusterOption struct{}

// Apply applies options to a graph instance
func (o *ClusterOption) Apply(g *Graph) {
	g.beCluster()
}

// GraphIDOption sets the graph internal ID
type GraphIDOption struct {
	// ID graph ID
	ID string
}

// Apply applies options to a graph instance
func (o *GraphIDOption) Apply(g *Graph) {
	g.id = o.ID
}

var (
	// Undirected node edges will have no direction, i.e. the layout engine can
	// freely organize the graph and draw glyphless edges between the nodes
	Undirected = &GraphTypeOption{"graph", false}
	// StrictUndirected same as Undirected, but multi-edges are forbidden i.e.
	// only one edge between the same two nodes is allowed
	StrictUndirected = &GraphTypeOption{"graph", true}
	// Directed node edges will have direction, i.e. will be used by the layout
	// engine to organize the graph and draw glyph edges denoting its orientation
	Directed = &GraphTypeOption{"digraph", false}
	// StrictDirected same as Directed, but multi-edges are forbidden i.e. only
	// one edge between the same head and tail nodes is allowed
	StrictDirected = &GraphTypeOption{"digraph", true}
	// Sub can be used to group objects, serve as a context for attributes or draw
	// its contents as a independent cluster, on supported layout engines
	Sub = &GraphTypeOption{"subgraph", false}
)

// GraphTypeOption sets the graph type and strictness
type GraphTypeOption struct {
	Name   string
	Strict bool
}

// Apply applies options to a graph instance
func (o *GraphTypeOption) Apply(g *Graph) {
	g.strict = o.Strict
	g.graphType = o.Name
}
