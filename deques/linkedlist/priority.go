package linkedlist

import (
	adt "goadt"
	"goadt/deques"
	"goadt/ds/linkedlist"
	"goadt/fn"
)

var _ deques.Deque[int] = (*Deque[int])(nil)

type PriorityDeque[E any] struct {
	compareFn fn.CompareFn[E]
	list      *linkedlist.List[E]
}

func NewPriorityDeque[E any](
	compareFn fn.CompareFn[E],
) *PriorityDeque[E] {
	return &PriorityDeque[E]{
		compareFn: compareFn,
		list:      linkedlist.New[E](),
	}
}

// PushBack implements deques.Deque.
func (l *PriorityDeque[E]) PushBack(element E) {
	node := l.list.FindLast(func(value E) bool {
		return l.compareFn(value, element) <= 0
	})
	if node == nil {
		l.list.Prepend(element)
	} else {
		node.Append(element)
	}
}

// PushFront implements deques.Deque.
func (l *PriorityDeque[E]) PushFront(element E) {
	node := l.list.Find(func(value E) bool {
		return l.compareFn(value, element) >= 0
	})
	if node == nil {
		l.list.Append(element)
	} else {
		node.Prepend(element)
	}
}

// PopBack implements deques.Deque.
func (l *PriorityDeque[E]) PopBack() E {
	node := l.list.Last()
	if node == nil {
		panic(adt.ErrNoSuchElement)
	}
	node.Remove()
	return node.GetValue()
}

// PopFront implements deques.Deque.
func (l *PriorityDeque[E]) PopFront() E {
	node := l.list.First()
	if node == nil {
		panic(adt.ErrNoSuchElement)
	}
	node.Remove()
	return node.GetValue()
}

// Back implements deques.Deque.
func (l *PriorityDeque[E]) Back() E {
	node := l.list.Last()
	if node == nil {
		panic(adt.ErrNoSuchElement)
	}
	return node.GetValue()
}

// Front implements deques.Deque.
func (l *PriorityDeque[E]) Front() E {
	node := l.list.First()
	if node == nil {
		panic(adt.ErrNoSuchElement)
	}
	return node.GetValue()
}

// IsEmpty implements deques.Deque.
func (l *PriorityDeque[E]) IsEmpty() bool {
	return l.list.IsEmpty()
}

// Size implements deques.Deque.
func (l *PriorityDeque[E]) Size() int {
	return l.list.Size()
}

// Clear implements deques.Deque.
func (l *PriorityDeque[E]) Clear() {
	l.list.Clear()
}
