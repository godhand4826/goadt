package bst

import (
	adt "goadt"
	"goadt/ds/bst"
	"goadt/fn"
	"goadt/queues"
)

var _ queues.Queue[int] = (*Queue[int])(nil)

type Queue[E any] struct {
	tree *bst.Tree[E]
}

func New[E any](compareFn fn.CompareFn[E]) *Queue[E] {
	return &Queue[E]{
		tree: bst.New(compareFn),
	}
}

// Enqueue implements queues.Queue.
func (b *Queue[E]) Enqueue(element E) {
	b.tree.Insert(element)
}

// Dequeue implements queues.Queue.
func (b *Queue[E]) Dequeue() E {
	if b.tree.IsEmpty() {
		panic(adt.ErrNoSuchElement)
	}
	node := b.tree.Minimum()
	node.Delete()
	return node.Value()
}

// IsEmpty implements queues.Queue.
func (b *Queue[E]) IsEmpty() bool {
	return b.tree.IsEmpty()
}

// Peek implements queues.Queue.
func (b *Queue[E]) Peek() E {
	if b.IsEmpty() {
		panic(adt.ErrNoSuchElement)
	}
	return b.tree.Minimum().Value()
}

// Size implements queues.Queue.
func (b *Queue[E]) Size() int {
	return b.tree.Size()
}

// Clear implements queues.Queue.
func (b *Queue[E]) Clear() {
	b.tree.Clear()
}
