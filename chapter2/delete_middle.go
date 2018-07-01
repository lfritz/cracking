package chapter2

// DeleteMiddle removes a middle node (i.e. not the first or last) from a list,
// given only that node. Solution to exercise 2.3.
func DeleteMiddle(n *Node) {
	for n.next.next != nil {
		n.data = n.next.data
		n = n.next
	}
	n.data = n.next.data
	n.next = nil
}
