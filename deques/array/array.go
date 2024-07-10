package array

import (
	"goadt/deques"
	"goadt/ds/array"
)

var _ deques.Deque[int] = (*Deque[int])(nil)

type Deque[E any] struct {
	array *array.Array[E]
}

func New[E any]() *Deque[E] {
	return &Deque[E]{
		array: array.New[E](),
	}
}

// PushBack implements deques.Deque.
func (d *Deque[E]) PushBack(element E) {
	d.array.Append(element)
}

// PushFront implements deques.Deque.
func (d *Deque[E]) PushFront(element E) {
	d.array.Prepend(element)
}

// PopBack implements deques.Deque.
func (d *Deque[E]) PopBack() E {
	return d.array.Remove(d.array.Size() - 1)
}

// PopFront implements deques.Deque.
func (d *Deque[E]) PopFront() E {
	return d.array.Remove(0)
}

// Back implements deques.Deque.
func (d *Deque[E]) Back() E {
	return d.array.At(d.Size() - 1)
}

// Front implements deques.Deque.
func (d *Deque[E]) Front() E {
	return d.array.At(0)
}

// IsEmpty implements deques.Deque.
func (d *Deque[E]) IsEmpty() bool {
	return d.array.IsEmpty()
}

// Size implements deques.Deque.
func (d *Deque[E]) Size() int {
	return d.array.Size()
}

// Clear implements deques.Deque.
func (d *Deque[E]) Clear() {
	d.array.Clear()
}
