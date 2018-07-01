package chapter2

// Partition partitions list so that all elements >= x come after all elements <
// x. Solution for exercise 2.4.
func Partition(list *Node, x int) *Node {
	leftHead := &Node{}
	rightHead := &Node{}
	leftTail := leftHead
	rightTail := rightHead
	for n := list; n != nil; n = n.next {
		if n.data < x {
			leftTail.next = n
			leftTail = n
		} else {
			rightTail.next = n
			rightTail = n
		}
	}
	rightTail.next = nil
	leftTail.next = rightHead.next
	return leftHead.next
}
