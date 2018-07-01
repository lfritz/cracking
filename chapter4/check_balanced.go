package chapter4

// Balanced checks if the tree is balanced. Here, balanced means the heights of
// the subtrees of any node never differ by more than one. Solution to exercise
// 4.4.
func (tree *BinaryTreeNode) Balanced() bool {
	_, balanced := tree.balance()
	return balanced
}

func (tree *BinaryTreeNode) balance() (height int, balanced bool) {
	if tree == nil {
		return 0, true
	}
	leftHeight, leftBalanced := tree.Left.balance()
	rightHeight, rightBalanced := tree.Right.balance()
	balanced = leftBalanced && rightBalanced &&
		leftHeight <= rightHeight+1 &&
		rightHeight <= leftHeight+1
	height = leftHeight + 1
	if rightHeight > leftHeight {
		height = rightHeight + 1
	}
	return
}
