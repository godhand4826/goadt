package linkedlist

import (
	adt "goadt"
	"goadt/ds/linkedlist"
	"goadt/queues"
)

var _ queues.Queue[int] = (*Queue[int])(nil)

type Queue[E any] struct {
	list *linkedlist.List[E]
}

func New[E any]() *Queue[E] {
	return &Queue[E]{
		list: linkedlist.New[E](),
	}
}

// Enqueue implements queues.Queue.
func (l *Queue[E]) Enqueue(element E) {
	l.list.Append(element)
}

// Dequeue implements queues.Queue.
func (l *Queue[E]) Dequeue() E {
	node := l.list.First()
	if node == nil {
		panic(adt.ErrNoSuchElement)
	}
	node.Remove()
	return node.GetValue()
}

// Peek implements queues.Queue.
func (l *Queue[E]) Peek() E {
	node := l.list.First()
	if node == nil {
		panic(adt.ErrNoSuchElement)
	}
	return node.GetValue()
}

// IsEmpty implements queues.Queue.
func (l *Queue[E]) IsEmpty() bool {
	return l.list.IsEmpty()
}

// Size implements queues.Queue.
func (l *Queue[E]) Size() int {
	return l.list.Size()
}

// Clear implements queues.Queue.
func (l *Queue[E]) Clear() {
	l.list.Clear()
}
