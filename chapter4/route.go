package chapter4

// DirectedGraph represents a directed graph using adjacency lists.
type DirectedGraph struct {
	nodes int
	adj   [][]int
}

// Route finds out if there is a path from a to b in the graph. Solution to
// exercise 4.1.
func (g *DirectedGraph) Route(a, b int) bool {
	seen := make([]bool, g.nodes)
	return routeHelper(g, seen, a, b)
}

func routeHelper(g *DirectedGraph, seen []bool, a, b int) bool {
	if a == b {
		return true
	}
	seen[a] = true
	for _, n := range g.adj[a] {
		if !seen[n] && routeHelper(g, seen, n, b) {
			return true
		}
	}
	return false
}
