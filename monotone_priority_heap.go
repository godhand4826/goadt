package main

var _ (Queue[int]) = (*monotonePriorityHeap[int])(nil)
var _ (Stack[int]) = (*monotonePriorityHeap[int])(nil)

type monotonePriorityHeap[E any] struct {
	elements Queue[*PriorityItem[E]]
	priority int
}

func NewHeapQueue[E any]() Queue[E] {
	return newMonotonePriorityHeap(On(Compare, func(item *PriorityItem[E]) int {
		return item.priority
	}))
}

func NewHeapStack[E any]() Stack[E] {
	return newMonotonePriorityHeap(On(Compare, func(item *PriorityItem[E]) int {
		return -item.priority
	}))
}

func newMonotonePriorityHeap[E any](compareFn CompareFn[*PriorityItem[E]]) *monotonePriorityHeap[E] {
	return &monotonePriorityHeap[E]{
		elements: NewHeap(compareFn),
	}
}

// Enqueue implements Queue.
func (m *monotonePriorityHeap[E]) Enqueue(element E) {
	m.priority++
	m.elements.Enqueue(&PriorityItem[E]{
		value:    element,
		priority: m.priority,
	})
}

// Dequeue implements Queue.
func (m *monotonePriorityHeap[E]) Dequeue() E {
	return m.elements.Dequeue().value
}

// Peek implements Queue.
func (m *monotonePriorityHeap[E]) Peek() E {
	return m.elements.Peek().value
}

// IsEmpty implements Queue.
func (m *monotonePriorityHeap[E]) IsEmpty() bool {
	return m.elements.IsEmpty()
}

// Size implements Queue.
func (m *monotonePriorityHeap[E]) Size() int {
	return m.elements.Size()
}

// Clear implements Queue.
func (m *monotonePriorityHeap[E]) Clear() {
	m.elements.Clear()
}

// Pop implements Stack.
func (m *monotonePriorityHeap[E]) Pop() E {
	return m.Dequeue()
}

// Push implements Stack.
func (m *monotonePriorityHeap[E]) Push(element E) {
	m.Enqueue(element)
}

// Top implements Stack.
func (m *monotonePriorityHeap[E]) Top() E {
	return m.Peek()
}

type PriorityItem[E any] struct {
	value    E
	priority int
}
