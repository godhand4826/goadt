package linkedlist

import (
	adt "goadt"
	"goadt/ds/linkedlist"
	"goadt/stacks"
)

var _ stacks.Stack[int] = (*Stack[int])(nil)

type Stack[E any] struct {
	list *linkedlist.List[E]
}

func NewStack[E any]() stacks.Stack[E] {
	return New[E]()
}

func New[E any]() *Stack[E] {
	return &Stack[E]{
		list: linkedlist.New[E](),
	}
}

// Push implements stacks.Stack.
func (l *Stack[E]) Push(element E) {
	l.list.Append(element)
}

// Pop implements stacks.Stack.
func (l *Stack[E]) Pop() E {
	node := l.list.Last()
	if node == nil {
		panic(adt.ErrNoSuchElement)
	}
	node.Remove()
	return node.GetValue()
}

// Top implements stacks.Stack.
func (l *Stack[E]) Top() E {
	node := l.list.Last()
	if node == nil {
		panic(adt.ErrNoSuchElement)
	}
	return node.GetValue()
}

// IsEmpty implements stacks.Stack.
func (l *Stack[E]) IsEmpty() bool {
	return l.list.IsEmpty()
}

// Size implements stacks.Stack.
func (l *Stack[E]) Size() int {
	return l.list.Size()
}

// Clear implements stacks.Stack.
func (l *Stack[E]) Clear() {
	l.list.Clear()
}
