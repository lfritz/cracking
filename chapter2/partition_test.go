package chapter2

import (
	"testing"
)

func TestPartition(t *testing.T) {
	cases := []struct {
		list string
		x    rune
		want string
	}{
		{"", '5', ""},
		{"1", '5', "1"},
		{"12", '5', "12"},
		{"67", '5', "67"},
		{"1672", '5', "1267"},
		{"7313958", '5', "3137958"},
	}
	for _, c := range cases {
		got := listToString(Partition(stringToList(c.list), int(c.x)))
		if got != c.want {
			t.Errorf("Partition(%s, %d) == %s, want %s",
				c.list, c.x, got, c.want)
		}
	}
}
