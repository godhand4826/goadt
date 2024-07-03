package lists

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
