package chapter2

func stringToList(s string) *Node {
	var list *Node
	for i := range s {
		data := s[len(s)-i-1]
		list = &Node{int(data), list}
	}
	return list
}

func listToString(list *Node) string {
	s := ""
	for n := list; n != nil; n = n.next {
		s = s + string(rune(n.data))
	}
	return s
}

func getNode(list *Node, index int) *Node {
	for i := 0; i < index; i++ {
		list = list.next
	}
	return list
}
