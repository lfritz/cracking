package chapter1

// IsUnique determines if s has all unique characters. Solution to exercise 1.1.
func IsUnique(s string) bool {
	seen := make(map[rune]bool)
	for _, c := range s {
		if seen[c] {
			return false
		}
		seen[c] = true
	}
	return true
}
