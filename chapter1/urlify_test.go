package chapter1

import "testing"

func TestUrlify(t *testing.T) {
	cases := []struct {
		input  string
		length int
		want   string
	}{
		{"a bc d....", 6, "a%20bc%20d"},
		{"Mr John Smith....", 13, "Mr%20John%20Smith"},
	}
	for _, c := range cases {
		array := []rune(c.input)
		URLify(array, c.length)
		got := string(array)
		if got != c.want {
			t.Errorf("URLify(%s) => %s, want %s", c.input, got, c.want)
		}
	}
}
