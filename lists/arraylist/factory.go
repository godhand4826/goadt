package arraylist

import (
	"goadt/fn"
	"goadt/lists"
)

var _ (lists.Factory[int]) = Factory[int]{}

type Factory[E any] struct{}

func (Factory[E]) CreateList(equalFn fn.EqualFn[E]) lists.List[E] {
	return New(equalFn)
}
