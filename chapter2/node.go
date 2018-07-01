package chapter2

// A Node is a node in a singly linked list.
type Node struct {
	data int
	next *Node
}

func get(list *Node, index int) int {
	if index == 0 {
		return list.data
	}
	return get(list.next, index-1)
}

func length(list *Node) int {
	if list == nil {
		return 0
	}
	return length(list.next) + 1
}
