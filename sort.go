package dot

import "sort"

func (g *graphData) sortedNodesKeys() []string {
	keys := make([]string, 0, len(g.subgraphs))

	for each := range g.nodes {
		keys = append(keys, each)
	}

	sort.StringSlice(keys).Sort()

	return keys
}
func (g *graphData) sortedEdgesFromKeys() []string {
	keys := make([]string, 0, len(g.subgraphs))

	for each := range g.edgesFrom {
		keys = append(keys, each)
	}

	sort.StringSlice(keys).Sort()

	return keys
}
func (g *graphData) sortedSubgraphsKeys() []string {
	keys := make([]string, 0, len(g.subgraphs))

	for each := range g.subgraphs {
		keys = append(keys, each)
	}

	sort.StringSlice(keys).Sort()

	return keys
}
