package main

var _ (Queue[int]) = (*Heap[int])(nil)

type Heap[E any] struct {
	compareFn CompareFn[E]
}

func NewHeap[E any](
	compareFn CompareFn[E],
) *Heap[E] {
	return &Heap[E]{
		compareFn: compareFn,
	}
}

// Clear implements Queue.
func (h *Heap[E]) Clear() {
	panic("unimplemented")
}

// Dequeue implements Queue.
func (h *Heap[E]) Dequeue() E {
	panic("unimplemented")
}

// Enqueue implements Queue.
func (h *Heap[E]) Enqueue(E) {
	panic("unimplemented")
}

// IsEmpty implements Queue.
func (h *Heap[E]) IsEmpty() bool {
	panic("unimplemented")
}

// Peek implements Queue.
func (h *Heap[E]) Peek() E {
	panic("unimplemented")
}

// Size implements Queue.
func (h *Heap[E]) Size() int {
	panic("unimplemented")
}
