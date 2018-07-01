package chapter1

import "testing"

func TestIsPalindromePermutation(t *testing.T) {
	cases := []struct {
		s    string
		want bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"tactcoa", true},   // palindrome: tacocat
		{"tact coa!", true}, // ignore spaces and punctuation
		{"Tact Coa", true},  // ignore case
		{"tactkoa", false},
		{"atactcoa", false},
	}
	for _, c := range cases {
		got := IsPalindromePermutation(c.s)
		if got != c.want {
			t.Errorf("IsPalindromePermutation(%v) == %v, want %v",
				c.s, got, c.want)
		}
	}
}
