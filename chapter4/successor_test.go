package chapter4

import "testing"

func TestSuccessor(t *testing.T) {
	root := DBinarySearchTree(8, 4, 2, 6, 10, 20, 15, 13, 17)
	cases := []struct {
		n    *DBinaryTreeNode
		want *DBinaryTreeNode
	}{
		{root.Left.Left, root.Left},
		{root.Left, root.Left.Right},
		{root.Left.Right, root},
		{root, root.Right},
		{root.Right, root.Right.Right.Left.Left},
	}
	for _, c := range cases {
		got := c.n.Successor()
		if got != c.want {
			t.Errorf("[%v].Successor() == [%v], want [%v]",
				c.n.Value, got.Value, c.want.Value)
		}
	}
}
