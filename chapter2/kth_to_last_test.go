package chapter2

import (
	"math/rand"
	"testing"
)

type testCase struct {
	list string
	k    int
	want int
}

var testCases = []testCase{
	{"1", 1, '1'},
	{"12345", 1, '5'},
	{"12345", 2, '4'},
	{"12345", 5, '1'},
}

func TestKthToLast1(t *testing.T) {
	for _, c := range testCases {
		got := KthToLast1(stringToList(c.list), c.k)
		if got != c.want {
			t.Errorf("KthToLast1(%v, %v) == %v, want %v",
				c.list, c.k, got, c.want)
		}
	}
}

func TestKthToLast2(t *testing.T) {
	for _, c := range testCases {
		got := KthToLast2(stringToList(c.list), c.k)
		if got != c.want {
			t.Errorf("KthToLast2(%v, %v) == %v, want %v",
				c.list, c.k, got, c.want)
		}
	}
}

func TestKthToLast3(t *testing.T) {
	for _, c := range testCases {
		got := KthToLast3(stringToList(c.list), c.k)
		if got != c.want {
			t.Errorf("KthToLast3(%v, %v) == %v, want %v",
				c.list, c.k, got, c.want)
		}
	}
}

const n = 1000

func randomList(length int) *Node {
	var list *Node
	for i := 0; i < length; i++ {
		list = &Node{rand.Int(), list}
	}
	return list
}

func BenchmarkKthToLast1(b *testing.B) {
	list := randomList(n)
	indexes := rand.Perm(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KthToLast1(list, indexes[i%n]+1)
	}
}

func BenchmarkKthToLast2(b *testing.B) {
	list := randomList(n)
	indexes := rand.Perm(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KthToLast2(list, indexes[i%n]+1)
	}
}

func BenchmarkKthToLast3(b *testing.B) {
	list := randomList(n)
	indexes := rand.Perm(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KthToLast3(list, indexes[i%n]+1)
	}
}
