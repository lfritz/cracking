package chapter4

import "testing"

var balancedTree = &BinaryTreeNode{5,
	&BinaryTreeNode{
		6,
		&BinaryTreeNode{7, nil, nil},
		nil},
	&BinaryTreeNode{
		8,
		nil,
		&BinaryTreeNode{9, nil, nil},
	},
}

var unbalancedTree = &BinaryTreeNode{5,
	&BinaryTreeNode{
		6,
		&BinaryTreeNode{7, nil, nil},
		nil},
	nil}

func TestBalance(t *testing.T) {
	cases := []struct {
		tree     *BinaryTreeNode
		height   int
		balanced bool
	}{
		{nil, 0, true},
		{&BinaryTreeNode{5, nil, nil}, 1, true},
		{balancedTree, 3, true},
		{unbalancedTree, 3, false},
	}
	for _, c := range cases {
		h, b := c.tree.balance()
		if h != c.height || b != c.balanced {
			t.Errorf("tree.balance() == %v, %v, want %v, %v",
				h, b, c.height, c.balanced)
		}
	}
}

func TestBalanced(t *testing.T) {
	cases := []struct {
		tree *BinaryTreeNode
		want bool
	}{
		{balancedTree, true},
		{unbalancedTree, false},
	}
	for _, c := range cases {
		got := c.tree.Balanced()
		if got != c.want {
			t.Errorf("tree.Balanced() == %v, want %v", got, c.want)
		}
	}
}
