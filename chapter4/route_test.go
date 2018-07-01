package chapter4

import "testing"

func TestRoute(t *testing.T) {
	graph := DirectedGraph{
		nodes: 7,
		adj: [][]int{
			[]int{1},
			[]int{2, 3},
			[]int{4, 6},
			[]int{},
			[]int{5},
			[]int{6},
			[]int{1},
		},
	}
	cases := []struct {
		a, b int
		want bool
	}{
		{0, 0, true},
		{1, 3, true},
		{0, 6, true},
		{6, 0, false},
	}
	for _, c := range cases {
		got := graph.Route(c.a, c.b)
		if got != c.want {
			t.Errorf("graph.Route(%v, %v) == %v, want %v",
				c.a, c.b, got, c.want)
		}
	}
}
