package chapter2

import (
	"testing"
)

func TestRemoveDups(t *testing.T) {
	cases := []struct {
		input, want string
	}{
		{"", ""},
		{"1", "1"},
		{"12", "12"},
		{"11", "1"},
		{"2425", "245"},
		{"121551", "125"},
		{"12345", "12345"},
	}
	for _, c := range cases {
		list := stringToList(c.input)
		RemoveDups(list)
		got := listToString(list)
		if got != c.want {
			t.Errorf("RemoveDups(%+v) == %+v, want %+v",
				list, got, c.want)
		}
	}
}
