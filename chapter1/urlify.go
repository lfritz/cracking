package chapter1

// URLify replaces all spaces in s with "%20". There must be enough space at the
// end for the extra characters, and n must give the length of the actual
// string. Solution to exercise 1.4.
func URLify(s []rune, n int) {
	spaces := []int{}
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			spaces = append(spaces, i)
		}
	}
	spaces = append(spaces, n)

	for i := len(spaces) - 1; i > 0; i-- {
		move(s, spaces[i-1]+1, spaces[i], 2*i)
		insert(s, spaces[i-1]+2*(i-1), []rune{'%', '2', '0'})
	}
}

func move(s []rune, start, end, by int) {
	for i := end - 1; i >= start; i-- {
		s[i+by] = s[i]
	}
}

func insert(target []rune, where int, what []rune) {
	for i, r := range what {
		target[where+i] = r
	}
}
