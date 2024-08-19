package heap

// Heap is a priority queue based on a binary heap.
type Heap[E any] interface {
	// Enqueue adds an element to the queue.
	Enqueue(element E)
	// Dequeue removes and returns the element with the highest priority.
	Dequeue() E
	// Peek returns the element with the highest priority.
	Peek() E
	// IsEmpty returns true if the queue is empty.
	IsEmpty() bool
	// Size returns the number of elements in the queue.
	Size() int
	// Clear removes all elements from the queue.
	Clear()
}
