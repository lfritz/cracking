package chapter1

// IsPermutation returns true if b is a permutation of a. Solution to exercise
// 1.2.
func IsPermutation(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	counters := make(map[rune]int)
	for _, r := range a {
		counters[r]++
	}
	for _, r := range b {
		if counters[r] == 0 {
			return false
		}
		counters[r]--
	}
	return true
}
