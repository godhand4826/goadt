package sets

// Set is an abstraction of a mathematical set.
type Set[E any] interface {
	Add(E)           // Adds an element to the set
	Remove(E) bool   // Removes an element from the set
	Contains(E) bool // Checks if the set contains a specific element
	IsEmpty() bool   // Checks if the set is empty
	Size() int       // Returns the number of elements in the set
	Clear()          // Removes all elements from the set
}
