package chapter4

// ValidateBST checks if a binary tree is a search tree. Solution to exercise
// 4.5.
func (tree *BinaryTreeNode) ValidateBST() bool {
	return tree.checkBST(0, 0, false, false)
}

func (n *BinaryTreeNode) checkBST(min, max int, checkMin, checkMax bool) bool {
	if n == nil {
		return true
	}
	return !(checkMin && n.Value <= min) &&
		!(checkMax && n.Value > max) &&
		n.Left.checkBST(min, n.Value, checkMin, true) &&
		n.Right.checkBST(n.Value, max, true, checkMax)
}
