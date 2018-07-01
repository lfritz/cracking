package chapter1

import (
	"strings"
	"unicode"
)

// IsPalindromePermutation returns true if s is a permutation of a palindrome.
// Solution to exercise 1.4.
func IsPalindromePermutation(s string) bool {
	input := make([]rune, 0, len(s))
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) {
			input = append(input, r)
		}
	}

	counts := make(map[rune]int)
	for _, r := range input {
		counts[r]++
	}

	foundOddCount := false
	for _, c := range counts {
		if c%2 == 1 {
			if foundOddCount {
				return false
			}
			foundOddCount = true
		}
	}
	return true
}
