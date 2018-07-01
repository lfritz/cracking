package chapter4

// A DBinaryTreeNode is a node in a doubly linked binary tree, i.e. one where
// each node has a pointer to its parent.
type DBinaryTreeNode struct {
	Value               int
	Parent, Left, Right *DBinaryTreeNode
}

func DBinaryTreeRoot(value int) *DBinaryTreeNode {
	return &DBinaryTreeNode{value, nil, nil, nil}
}

func (n *DBinaryTreeNode) Add(value int) {
	if value <= n.Value {
		if n.Left == nil {
			n.Left = &DBinaryTreeNode{value, n, nil, nil}
		} else {
			n.Left.Add(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &DBinaryTreeNode{value, n, nil, nil}
		} else {
			n.Right.Add(value)
		}
	}
}

func DBinarySearchTree(values ...int) *DBinaryTreeNode {
	root := DBinaryTreeRoot(values[0])
	for _, v := range values[1:len(values)] {
		root.Add(v)
	}
	return root
}

// Successor returns the in-order successor of a node in a binary search tree.
// Solution to exercise 4.6.
func (n *DBinaryTreeNode) Successor() *DBinaryTreeNode {
	if n.Right == nil {
		return n.successorUp()
	}
	return n.Right.leftMost()
}

func (n *DBinaryTreeNode) leftMost() *DBinaryTreeNode {
	if n.Left == nil {
		return n
	}
	return n.Left.leftMost()
}

func (n *DBinaryTreeNode) successorUp() *DBinaryTreeNode {
	if n.Parent == nil {
		return nil
	}
	if n.Parent.Left == n {
		return n.Parent
	}
	return n.Parent.successorUp()
}
