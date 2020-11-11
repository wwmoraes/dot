package attributes

// ClusterMode mode used for handling clusters
type ClusterMode string

const (
	// ClusterModeLocal gives a special treatment to a subgraph whose name begins
	// with "cluster":
	//
	// The subgraph is laid out separately, and then integrated as a unit into its
	// parent graph, with a bounding rectangle drawn about it. If the cluster has
	// a label parameter, this label is displayed within the rectangle.
	ClusterModeLocal ClusterMode = "local"
	// ClusterModeGlobal turns off any special cluster processing
	ClusterModeGlobal ClusterMode = "global"
	// ClusterModeNone turns off any special cluster processing
	ClusterModeNone ClusterMode = "none"
)
