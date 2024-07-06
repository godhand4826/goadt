package bst

import (
	adt "goadt"
	"goadt/ds/bst"
	"goadt/fn"
	"goadt/queues"
)

var _ queues.Queue[int] = (*Bst[int])(nil)

type Bst[E any] struct {
	tree *bst.Tree[E]
}

func New[E any](compareFn fn.CompareFn[E]) *Bst[E] {
	return &Bst[E]{
		tree: bst.New(compareFn),
	}
}

// Enqueue implements queues.Queue.
func (b *Bst[E]) Enqueue(element E) {
	b.tree.Insert(element)
}

// Dequeue implements queues.Queue.
func (b *Bst[E]) Dequeue() E {
	if b.tree.IsEmpty() {
		panic(adt.ErrNoSuchElement)
	}
	node := b.tree.Minimum()
	node.Delete()
	return node.Value()
}

// IsEmpty implements queues.Queue.
func (b *Bst[E]) IsEmpty() bool {
	return b.tree.IsEmpty()
}

// Peek implements queues.Queue.
func (b *Bst[E]) Peek() E {
	if b.IsEmpty() {
		panic(adt.ErrNoSuchElement)
	}
	return b.tree.Minimum().Value()
}

// Size implements queues.Queue.
func (b *Bst[E]) Size() int {
	return b.tree.Size()
}

// Clear implements queues.Queue.
func (b *Bst[E]) Clear() {
	b.tree.Clear()
}
