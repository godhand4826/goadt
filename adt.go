// nolint: lll
package main

import "errors"

var ErrNoSuchElement = errors.New("no such element")
var ErrIndexOutOfBound = errors.New("index out of bound")

// List is an abstraction of sequence of elements.
type List[E any] interface {
	NewList() List[E]       // Creates an empty list as a workaround for the factory pattern
	Append(E)               // Adds an element to the end of the list
	Prepend(E)              // Adds an element to the beginning of the list
	At(int) E               // Returns the element at the specified index
	ReplaceAt(int, E) E     // Replaces the element at the specified index with the given element, and returns the element previously at that position
	InsertAt(int, E)        // Inserts an element at the specified index
	RemoveAt(int) E         // Removes and returns the element at the specified index
	Slice(int, int) List[E] // Returns a new list containing elements from the start index (inclusive) to the end index (exclusive)
	IndexOf(E) int          // Returns the index of the first occurrence of the given element, or -1 if the element is not found
	Remove(E) bool          // Removes the first occurrence of the specified element from the list, returning true if the element was found and removed
	Contains(E) bool        // Checks if the list contains the specified element
	IsEmpty() bool          // Checks if the list is empty
	Size() int              // Returns the number of elements in the list
	Clear()                 // Removes all elements from the list
}

// Queue is an abstraction of the queuing concept.
// Queues typically, but do not necessarily, order elements in a FIFO (first-in-first-out) manner.
// Each Queue implementation may have its own priority properties. For example:
//   - FIFO (first-in-first-out) Queue: the priority of each inserted element is monotonically increasing
//   - LIFO (last-in-first-out) Queue: the priority of each inserted element is monotonically decreasing
//   - Priority Queue: priority determined by comparing elements or associated priority
type Queue[E any] interface {
	Enqueue(E)     // Adds an element to the queue
	Dequeue() E    // Removes and returns the front element from the queue
	Peek() E       // Returns the front element without removing it
	IsEmpty() bool // Checks if the queue is empty
	Size() int     // Returns the number of elements in the queue
	Clear()        // Removes all elements from the queue
}

// Stack is an abstraction of the stack data structure.
// A Stack should always follow the LIFO (last-in-first-out) principle.
type Stack[E any] interface {
	Push(E)        // Adds an element to the top of the stack
	Pop() E        // Removes and returns the top element from the stack
	Top() E        // Returns the top element without removing it
	IsEmpty() bool // Checks if the stack is empty
	Size() int     // Returns the number of elements in the stack
	Clear()        // Removes all elements from the stack
}

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

// Set is an abstraction of a mathematical set.
type Set[E any] interface {
	Add(E)           // Adds an element to the set
	Remove(E) bool   // Removes an element from the set
	Contains(E) bool // Checks if the set contains a specific element
	IsEmpty() bool   // Checks if the set is empty
	Size() int       // Returns the number of elements in the set
	Clear()          // Removes all elements from the set
}
