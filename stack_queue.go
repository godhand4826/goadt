package main

var _ (Queue[int]) = (*StackQueue[int])(nil)

// StackQueue wraps a stack in to a LIFO queue.
type StackQueue[E any] struct {
	Stack[E]
}

func NewStackQueue[E any](s Stack[E]) Queue[E] {
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
