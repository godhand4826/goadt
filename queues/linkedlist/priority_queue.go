package linkedlist

import (
	adt "goadt"
	"goadt/ds/linkedlist"
	"goadt/fn"
	"goadt/queues"
)

var _ queues.Queue[int] = (*PriorityQueue[int])(nil)

type PriorityQueue[E any] struct {
	compareFn fn.CompareFn[E]
	list      *linkedlist.List[E]
}

func NewPriorityQueue[E any](
	compareFn fn.CompareFn[E],
) *PriorityQueue[E] {
	return &PriorityQueue[E]{
		compareFn: compareFn,
		list:      linkedlist.New[E](),
	}
}

// Enqueue implements queues.Queue.
func (l *PriorityQueue[E]) Enqueue(element E) {
	node := l.list.Find(func(value E) bool {
		return l.compareFn(value, element) > 0
	})
	if node == nil {
		l.list.Append(element)
	} else {
		node.Prepend(element)
	}
}

// Dequeue implements queues.Queue.
func (l *PriorityQueue[E]) Dequeue() E {
	node := l.list.First()
	if node == nil {
		panic(adt.ErrNoSuchElement)
	}
	node.Remove()
	return node.GetValue()
}

// Peek implements queues.Queue.
func (l *PriorityQueue[E]) Peek() E {
	node := l.list.First()
	if node == nil {
		panic(adt.ErrNoSuchElement)
	}
	return node.GetValue()
}

// IsEmpty implements queues.Queue.
func (l *PriorityQueue[E]) IsEmpty() bool {
	return l.list.IsEmpty()
}

// Size implements queues.Queue.
func (l *PriorityQueue[E]) Size() int {
	return l.list.Size()
}

// Clear implements queues.Queue.
func (l *PriorityQueue[E]) Clear() {
	l.list.Clear()
}
