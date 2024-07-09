package linkedlist

import (
	adt "goadt"
	"goadt/deques"
	"goadt/ds/linkedlist"
)

var _ deques.Deque[int] = (*Deque[int])(nil)

type Deque[E any] struct {
	list *linkedlist.List[E]
}

func New[E any]() *Deque[E] {
	return &Deque[E]{
		list: linkedlist.New[E](),
	}
}

// PushBack implements deques.Deque.
func (l *Deque[E]) PushBack(element E) {
	l.list.Append(element)
}

// PushFront implements deques.Deque.
func (l *Deque[E]) PushFront(element E) {
	l.list.Prepend(element)
}

// PopBack implements deques.Deque.
func (l *Deque[E]) PopBack() E {
	node := l.list.Last()
	if node == nil {
		panic(adt.ErrNoSuchElement)
	}
	node.Remove()
	return node.GetValue()
}

// PopFront implements deques.Deque.
func (l *Deque[E]) PopFront() E {
	node := l.list.First()
	if node == nil {
		panic(adt.ErrNoSuchElement)
	}
	node.Remove()
	return node.GetValue()
}

// Back implements deques.Deque.
func (l *Deque[E]) Back() E {
	node := l.list.Last()
	if node == nil {
		panic(adt.ErrNoSuchElement)
	}
	return node.GetValue()
}

// Front implements deques.Deque.
func (l *Deque[E]) Front() E {
	node := l.list.First()
	if node == nil {
		panic(adt.ErrNoSuchElement)
	}
	return node.GetValue()
}

// IsEmpty implements deques.Deque.
func (l *Deque[E]) IsEmpty() bool {
	return l.list.IsEmpty()
}

// Size implements deques.Deque.
func (l *Deque[E]) Size() int {
	return l.list.Size()
}

// Clear implements deques.Deque.
func (l *Deque[E]) Clear() {
	l.list.Clear()
}
