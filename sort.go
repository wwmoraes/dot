package dot

import "sort"

func (g *graph) sortedNodesKeys() (keys []string) {
	for each := range g.nodes {
		keys = append(keys, each)
	}
	sort.StringSlice(keys).Sort()
	return
}
func (g *graph) sortedEdgesFromKeys() (keys []string) {
	for each := range g.edgesFrom {
		keys = append(keys, each)
	}
	sort.StringSlice(keys).Sort()
	return
}
func (g *graph) sortedSubgraphsKeys() (keys []string) {
	for each := range g.subgraphs {
		keys = append(keys, each)
	}
	sort.StringSlice(keys).Sort()
	return
}
