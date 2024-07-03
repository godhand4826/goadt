package queues

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
