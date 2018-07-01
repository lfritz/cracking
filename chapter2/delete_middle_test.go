package chapter2

import "testing"

func TestDeleteMiddle(t *testing.T) {
	cases := []struct {
		input string
		index int
		want  string
	}{
		{"123", 1, "13"},
		{"12345", 1, "1345"},
		{"12345", 2, "1245"},
		{"12345", 3, "1235"},
	}
	for _, c := range cases {
		list := stringToList(c.input)
		DeleteMiddle(getNode(list, c.index))
		got := listToString(list)
		if got != c.want {
			t.Errorf("RemoveMiddle(%s[%d]) => %s, want %s",
				c.input, c.index, got, c.want)
		}
	}
}
