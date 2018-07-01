package chapter4

// A ListNode is a node in a singly linked list.
type ListNode struct {
	value int
	next  *ListNode
}

// ListOfDepths returns a list of all nodes at each depth. Solution to exercise
// 4.3.
func ListOfDepths(tree *BinaryTreeNode) []ListNode {
	result := make([]ListNode, 0)
	if tree == nil {
		return result
	}
	queue := []*BinaryTreeNode{tree}
	for len(queue) != 0 {
		nextQueue := make([]*BinaryTreeNode, 0)
		var list *ListNode
		for _, n := range queue {
			list = &ListNode{n.Value, list}
			if n.Left != nil {
				nextQueue = append(nextQueue, n.Left)
			}
			if n.Right != nil {
				nextQueue = append(nextQueue, n.Right)
			}
		}
		result = append(result, *list)
		queue = nextQueue
	}
	return result
}
