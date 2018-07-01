package chapter1

import "testing"

func TestIsUnique(t *testing.T) {
	cases := []struct {
		s    string
		want bool
	}{
		{"", true},
		{"a", true},
		{"abcde", true},
		{"aa", false},
		{"abcbe", false},
	}
	for _, c := range cases {
		got := IsUnique(c.s)
		if got != c.want {
			t.Errorf("IsUnique(%v) == %v, want %v", c.s, got, c.want)
		}
	}
}
