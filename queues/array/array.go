package array

import (
	"goadt/ds/array"
	"goadt/queues"
)

var _ queues.Queue[int] = (*Queue[int])(nil)

type Queue[E any] struct {
	array *array.Array[E]
}

func New[E any]() *Queue[E] {
	return &Queue[E]{
		array: array.New[E](),
	}
}

// Enqueue implements queues.Queue.
func (l *Queue[E]) Enqueue(element E) {
	l.array.Append(element)
}

// Dequeue implements queues.Queue.
func (l *Queue[E]) Dequeue() E {
	return l.array.Remove(0)
}

// Peek implements queues.Queue.
func (l *Queue[E]) Peek() E {
	return l.array.At(0)
}

// IsEmpty implements queues.Queue.
func (l *Queue[E]) IsEmpty() bool {
	return l.array.IsEmpty()
}

// Size implements queues.Queue.
func (l *Queue[E]) Size() int {
	return l.array.Size()
}

// Clear implements queues.Queue.
func (l *Queue[E]) Clear() {
	l.array.Clear()
}
