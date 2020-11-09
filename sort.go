package dot

import "sort"

func (g *GraphData) sortedNodesKeys() (keys []string) {
	for each := range g.nodes {
		keys = append(keys, each)
	}
	sort.StringSlice(keys).Sort()
	return
}
func (g *GraphData) sortedEdgesFromKeys() (keys []string) {
	for each := range g.edgesFrom {
		keys = append(keys, each)
	}
	sort.StringSlice(keys).Sort()
	return
}
func (g *GraphData) sortedSubgraphsKeys() (keys []string) {
	for each := range g.subgraphs {
		keys = append(keys, each)
	}
	sort.StringSlice(keys).Sort()
	return
}
