package main

import "container/heap"

var _ (Queue[int]) = (*Heap[int])(nil)

type Heap[E any] struct {
	compareFn CompareFn[E]
	arrayList *ArrayList[E]
}

func NewHeap[E any](
	compareFn CompareFn[E],
) Queue[E] {
	return &Heap[E]{
		compareFn: compareFn,
		arrayList: NewArrayList[E](nil),
	}
}

// Enqueue implements Queue.
func (h *Heap[E]) Enqueue(element E) {
	heap.Push(h, element)
}

// Dequeue implements Queue.
func (h *Heap[E]) Dequeue() E {
	return heap.Pop(h).(E)
}

// Peek implements Queue.
func (h *Heap[E]) Peek() E {
	if h.IsEmpty() {
		panic(ErrNoSuchElement)
	}
	return h.arrayList.At(0)
}

// IsEmpty implements Queue.
func (h *Heap[E]) IsEmpty() bool {
	return h.Size() == 0
}

// Size implements Queue.
func (h *Heap[E]) Size() int {
	return h.arrayList.Size()
}

// Clear implements Queue.
func (h *Heap[E]) Clear() {
	h.arrayList.Clear()
}

// Implement heap.Interface.
var _ heap.Interface = (*Heap[int])(nil)

// Len implements heap.Interface.
func (h *Heap[V]) Len() int {
	return h.arrayList.Size()
}

// Less implements heap.Interface.
func (h *Heap[V]) Less(i int, j int) bool {
	return h.compareFn(h.arrayList.At(i), h.arrayList.At(j)) < 0
}

// Pop implements heap.Interface.
func (h *Heap[V]) Pop() any {
	return h.arrayList.Pop()
}

// Push implements heap.Interface.
func (h *Heap[V]) Push(x any) {
	h.arrayList.Push(x.(V))
}

// Swap implements heap.Interface.
func (h *Heap[V]) Swap(i int, j int) {
	h.arrayList.Swap(i, j)
}
