package chapter4

import (
	"reflect"
	"testing"
)

func TestListOfDepths(t *testing.T) {
	tree := &BinaryTreeNode{
		5,
		&BinaryTreeNode{
			1,
			&BinaryTreeNode{
				2,
				&BinaryTreeNode{
					8,
					&BinaryTreeNode{6, nil, nil},
					&BinaryTreeNode{10, nil, nil},
				},
				nil,
			},
			&BinaryTreeNode{4, nil, nil},
		},
		&BinaryTreeNode{
			7,
			&BinaryTreeNode{9, nil, nil},
			&BinaryTreeNode{3, nil, nil},
		},
	}
	want := []ListNode{
		{5, nil},
		{7, &ListNode{1, nil}},
		{3, &ListNode{9, &ListNode{4, &ListNode{2, nil}}}},
		{8, nil},
		{10, &ListNode{6, nil}},
	}
	got := ListOfDepths(tree)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("ListOfDepths(%v) == %v, want %v", tree, got, want)
	}
}
