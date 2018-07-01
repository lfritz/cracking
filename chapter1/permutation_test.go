package chapter1

import "testing"

func TestPermutation(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	}{
		{"", "", true},
		{"a", "", false},
		{"", "a", false},
		{"a", "a", true},
		{"ab", "ba", true},
		{"ab", "ca", false},
		{"abcde", "abcde", true},
		{"abcde", "eabdc", true},
		{"abcde", "abfde", false},
	}
	for _, c := range cases {
		got := IsPermutation(c.a, c.b)
		if got != c.want {
			t.Errorf("IsPermutation(%v, %v) == %v, want %v",
				c.a, c.b, got, c.want)
		}
	}
}
