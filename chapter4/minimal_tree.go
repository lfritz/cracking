package chapter4

// Search checks if value is present in the tree, assuming it's a search tree.
func (n *BinaryTreeNode) Search(value int) bool {
	if n == nil {
		return false
	}
	if n.Value == value {
		return true
	}
	if value < n.Value {
		return n.Left.Search(value)
	}
	return n.Right.Search(value)
}

// MinimalTree creates a binary search tree of minimal height, given a sorted
// array of unique integers.
func MinimalTree(values []int) *BinaryTreeNode {
	n := len(values)
	if n == 0 {
		return nil
	}
	root := n / 2
	return &BinaryTreeNode{
		values[root],
		MinimalTree(values[0:root]),
		MinimalTree(values[root+1 : n]),
	}
}
