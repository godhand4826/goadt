package linkedlist

import (
	"goadt/ds/linkedlist"
	"goadt/fn"
	"goadt/lists"
)

var _ lists.List[int] = (*List[int])(nil)

type List[E any] struct {
	list    *linkedlist.List[E]
	equalFn fn.EqualFn[E]
}

func NewList[E any](equalFn fn.EqualFn[E]) lists.List[E] {
	return New(equalFn)
}

func New[E any](equalFn fn.EqualFn[E]) *List[E] {
	return &List[E]{
		list:    linkedlist.New[E](),
		equalFn: equalFn,
	}
}

// Append implements lists.List.
func (l *List[E]) Append(element E) {
	l.list.Append(element)
}

// At implements lists.List.
func (l *List[E]) At(index int) E {
	return l.list.At(index).GetValue()
}

// Clear implements lists.List.
func (l *List[E]) Clear() {
	l.list.Clear()
}

// Contains implements lists.List.
func (l *List[E]) Contains(element E) bool {
	return l.list.Find(func(value E) bool {
		return l.equalFn(element, value)
	}) != nil
}

// IndexOf implements lists.List.
func (l *List[E]) IndexOf(element E) int {
	node := l.list.First()
	for i := 0; i < l.Size(); i++ {
		if l.equalFn(node.GetValue(), element) {
			return i
		}
		node = node.Next()
	}
	return -1
}

// InsertAt implements lists.List.
func (l *List[E]) InsertAt(index int, element E) {
	l.list.At(index).Prepend(element)
}

// IsEmpty implements lists.List.
func (l *List[E]) IsEmpty() bool {
	return l.list.IsEmpty()
}

// Prepend implements lists.List.
func (l *List[E]) Prepend(element E) {
	l.list.Prepend(element)
}

// Remove implements lists.List.
func (l *List[E]) Remove(element E) bool {
	node := l.list.Find(func(value E) bool {
		return l.equalFn(value, element)
	})
	if node == nil {
		return false
	}
	node.Remove()
	return true
}

// RemoveAt implements lists.List.
func (l *List[E]) RemoveAt(index int) E {
	node := l.list.At(index)
	value := node.GetValue()
	node.Remove()
	return value
}

// ReplaceAt implements lists.List.
func (l *List[E]) ReplaceAt(index int, element E) E {
	node := l.list.At(index)
	value := node.GetValue()
	node.SetValue(element)
	return value
}

// Size implements lists.List.
func (l *List[E]) Size() int {
	return l.list.Size()
}

// Slice implements lists.List.
func (l *List[E]) Slice(start int, end int) lists.List[E] {
	start = max(start, 0)
	end = min(l.list.Size(), end)

	list := New(l.equalFn)
	node := l.list.First()
	for i := 0; i < end; i++ {
		if i >= start {
			list.Append(node.GetValue())
		}
		node = node.Next()
	}
	return list
}

// Swap implements lists.List.
func (l *List[E]) Swap(i int, j int) {
	n1 := l.list.At(i)
	n2 := l.list.At(j)

	v := n1.GetValue()
	n1.SetValue(n2.GetValue())
	n2.SetValue(v)
}
