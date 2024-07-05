package arraylist

import (
	"goadt/fn"
	"goadt/lists"
)

var _ (lists.ListFactory[int]) = ArrayListFactory[int]{}

type ArrayListFactory[E any] struct{}

func (ArrayListFactory[E]) CreateList(equalFn fn.EqualFn[E]) lists.List[E] {
	return NewArrayList(equalFn)
}
