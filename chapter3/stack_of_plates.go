package chapter3

// A StackOfPlates is a stack data structure that's composed of a set of stacks
// of fixed capacity. Solution for exercise 3.3.
type StackOfPlates struct {
	capacity int
	stacks   [][]int
}

// NewStackOfPlates returns an empty stack.
func NewStackOfPlates(capacity int) *StackOfPlates {
	return &StackOfPlates{capacity, make([][]int, 0)}
}

// Push pushes an element onto the stack.
func (s *StackOfPlates) Push(value int) {
	current := len(s.stacks) - 1
	if current == -1 || len(s.stacks[current]) == s.capacity {
		s.stacks = append(s.stacks, make([]int, 0, s.capacity))
		current++
	}
	s.stacks[current] = append(s.stacks[current], value)
}

// Pop removes the stack's top element and returns it.
func (s *StackOfPlates) Pop() int {
	current := len(s.stacks) - 1
	currentLength := len(s.stacks[current])
	value := s.stacks[current][currentLength-1]
	if currentLength == 1 {
		s.stacks = s.stacks[0:current]
	} else {
		s.stacks[current] = s.stacks[current][0 : currentLength-1]
	}
	return value
}
