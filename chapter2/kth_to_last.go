package chapter2

// KthToLast1 returns the k-th-to-last element of the list. Recursive solution
// for exercise 2.2.
func KthToLast1(list *Node, k int) int {
	return get(list, length(list)-k)
}

// KthToLast2 returns the k-th-to-last element of the list. Iterative solution
// for exercise 2.2.
func KthToLast2(list *Node, k int) int {
	length := 0
	for n := list; n != nil; n = n.next {
		length++
	}
	index := length - k
	for i := 0; i < index; i++ {
		list = list.next
	}
	return list.data
}

// KthToLast3 returns the k-th-to-last element of the list. Solution with fast
// and slow pointers for exercise 2.2.
func KthToLast3(list *Node, k int) int {
	fast := list
	for i := 0; i < k; i++ {
		fast = fast.next
	}
	slow := list
	for fast != nil {
		fast = fast.next
		slow = slow.next
	}
	return slow.data
}
