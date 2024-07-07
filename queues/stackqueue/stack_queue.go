package stackqueue

import (
	"goadt/queues"
	"goadt/stacks"
)

var _ (queues.Queue[int]) = (*Queue[int])(nil)

// Queue wraps a stack in to a LIFO queue.
type Queue[E any] struct {
	stacks.Stack[E]
}

func New[E any](s stacks.Stack[E]) queues.Queue[E] {
	return &Queue[E]{s}
}

func (s *Queue[E]) Enqueue(element E) {
	s.Push(element)
}

func (s *Queue[E]) Dequeue() E {
	return s.Pop()
}

func (s *Queue[E]) Peek() E {
	return s.Top()
}
