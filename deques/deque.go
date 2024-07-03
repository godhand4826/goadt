package deques

// Deque represents a double-ended queue.
// A Deque typically, but not necessarily, operates in a LIFO (last-in-first-out) manner at each end.
// Implementers are free to define the behavior of the "Push" and "Pop" methods,
// and are not constrained by the names "front" and "back".
type Deque[E any] interface {
	PushBack(E)    // Adds an element to the back of the deque
	PushFront(E)   // Adds an element to the front of the deque
	PopBack() E    // Removes and returns the back element from the deque
	PopFront() E   // Removes and returns the front element from the deque
	Back() E       // Returns the back element without removing it
	Front() E      // Returns the front element without removing it
	IsEmpty() bool // Checks if the deque is empty
	Size() int     // Returns the number of elements in the deque
	Clear()        // Removes all elements from the deque
}
