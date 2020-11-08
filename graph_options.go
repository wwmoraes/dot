package dot

type GraphOption interface {
	Apply(*Graph)
}

type ClusterOption struct{}

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
	Strict     = &GraphTypeOption{"strict"}
	Undirected = &GraphTypeOption{"graph"}
	Directed   = &GraphTypeOption{"digraph"}
	Sub        = &GraphTypeOption{"subgraph"}
)

type GraphTypeOption struct {
	Name string
}

func (o *GraphTypeOption) Apply(g *Graph) {
	g.graphType = o.Name
}
