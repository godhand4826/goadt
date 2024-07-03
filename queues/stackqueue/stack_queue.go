package stackqueue

import (
	"goadt/queues"
	"goadt/stacks"
)

var _ (queues.Queue[int]) = (*StackQueue[int])(nil)

// StackQueue wraps a stack in to a LIFO queue.
type StackQueue[E any] struct {
	stacks.Stack[E]
}

func NewStackQueue[E any](s stacks.Stack[E]) queues.Queue[E] {
	return &StackQueue[E]{s}
}

func (s *StackQueue[E]) Enqueue(element E) {
	s.Push(element)
}

func (s *StackQueue[E]) Dequeue() E {
	return s.Pop()
}

func (s *StackQueue[E]) Peek() E {
	return s.Top()
}
