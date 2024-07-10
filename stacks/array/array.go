package array

import (
	"goadt/ds/array"
	"goadt/stacks"
)

var _ stacks.Stack[int] = (*Stack[int])(nil)

type Stack[E any] struct {
	array *array.Array[E]
}

func New[E any]() *Stack[E] {
	return &Stack[E]{
		array: array.New[E](),
	}
}

// Pop implements stacks.Stack.
func (s *Stack[E]) Pop() E {
	return s.array.Remove(s.array.Size() - 1)
}

// Push implements stacks.Stack.
func (s *Stack[E]) Push(element E) {
	s.array.Append(element)
}

// Top implements stacks.Stack.
func (s *Stack[E]) Top() E {
	return s.array.At(s.array.Size() - 1)
}

// Size implements stacks.Stack.
func (s *Stack[E]) Size() int {
	return s.array.Size()
}

// IsEmpty implements stacks.Stack.
func (s *Stack[E]) IsEmpty() bool {
	return s.array.IsEmpty()
}

// Clear implements stacks.Stack.
func (s *Stack[E]) Clear() {
	s.array.Clear()
}
