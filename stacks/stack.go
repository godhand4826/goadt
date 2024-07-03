package stacks

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
