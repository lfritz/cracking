package chapter4

import "testing"

func tree(value int, left, right *BinaryTreeNode) *BinaryTreeNode {
	return &BinaryTreeNode{value, left, right}
}

func leaf(value int) *BinaryTreeNode {
	return &BinaryTreeNode{value, nil, nil}
}

func TestValidateBST(t *testing.T) {
	cases := []struct {
		tree *BinaryTreeNode
		want bool
	}{
		{
			tree(8, tree(4, leaf(2), leaf(6)),
				tree(10, nil, leaf(20))),
			true,
		},
		{
			tree(8, tree(4, leaf(2), leaf(12)),
				tree(10, nil, leaf(20))),
			false,
		},
		{
			tree(8, tree(4, leaf(2), leaf(6)),
				tree(10, leaf(7), nil)),
			false,
		},
		{
			tree(8, tree(4, leaf(4), leaf(6)),
				tree(10, nil, leaf(20))),
			true,
		},
		{
			tree(8, tree(4, leaf(2), leaf(6)),
				tree(10, leaf(8), nil)),
			false,
		},
	}
	for _, c := range cases {
		got := c.tree.ValidateBST()
		if got != c.want {
			t.Errorf("tree.ValidateBST() == %v, want %v",
				got, c.want)
		}
	}
}
