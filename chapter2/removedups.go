package chapter2

// RemoveDups removes duplicates from an unsorted list. Solution to exercise 2.1.
func RemoveDups(list *Node) {
	seen := make(map[int]bool)
	var previous *Node
	for current := list; current != nil; current = current.next {
		if seen[current.data] {
			previous.next = current.next
			continue
		}
		seen[current.data] = true
		previous = current
	}
}
