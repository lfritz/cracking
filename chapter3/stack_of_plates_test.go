package chapter3

import (
	"reflect"
	"testing"
)

func TestStackOfPlates(t *testing.T) {
	cases := []struct {
		capacity   int
		operations []int // -1 means Pop, otherwise Push the value
		want       []int // result of Pop operations
	}{
		{4, []int{1, 2, 3, 4, -1, -1, -1, -1}, []int{4, 3, 2, 1}},
		{2, []int{1, 2, 3, 4, -1, -1, -1, -1}, []int{4, 3, 2, 1}},
		{2, []int{1, 2, 3, -1, 4, -1, -1, -1}, []int{3, 4, 2, 1}},
	}
	for _, c := range cases {
		stack := NewStackOfPlates(c.capacity)
		got := make([]int, 0)
		for _, op := range c.operations {
			if op == -1 {
				got = append(got, stack.Pop())
			} else {
				stack.Push(op)
			}
		}
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("StackOfPlates yielded %v, expected %v",
				got, c.want)
		}
	}

}
