package chapter4

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	tree := &BinaryTreeNode{
		5,
		&BinaryTreeNode{3, nil, nil},
		&BinaryTreeNode{8, &BinaryTreeNode{6, nil, nil}, nil},
	}
	cases := []struct {
		value int
		want  bool
	}{
		{5, true},
		{3, true},
		{6, true},
		{1, false},
		{9, false},
		{7, false},
	}
	for _, c := range cases {
		got := tree.Search(c.value)
		if got != c.want {
			t.Errorf("exampleTree.Search(%v) == %v, want %v",
				c.value, got, c.want)
		}
	}
}

func TestMinimalTree(t *testing.T) {
	cases := []struct {
		array []int
		want  *BinaryTreeNode
	}{
		{[]int{}, nil},
		{[]int{0}, &BinaryTreeNode{0, nil, nil}},
		{[]int{0, 1, 2, 3, 4},
			&BinaryTreeNode{
				2,
				&BinaryTreeNode{1,
					&BinaryTreeNode{0, nil, nil},
					nil,
				},
				&BinaryTreeNode{4,
					&BinaryTreeNode{3, nil, nil},
					nil,
				},
			},
		},
	}
	for _, c := range cases {
		got := MinimalTree(c.array)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MinimalTree(%v) == %v, want %v",
				c.array, got, c.want)
		}
	}
}
