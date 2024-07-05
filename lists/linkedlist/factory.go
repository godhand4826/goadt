package linkedlist

import (
	"goadt/fn"
	"goadt/lists"
)

var _ (lists.ListFactory[int]) = LinkedListFactory[int]{}

type LinkedListFactory[E any] struct{}

func (LinkedListFactory[E]) CreateList(equalFn fn.EqualFn[E]) lists.List[E] {
	return NewLinkedList(equalFn)
}
