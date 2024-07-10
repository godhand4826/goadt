package array

import (
	"goadt/ds/array"
	"goadt/fn"
	"goadt/lists"
)

var _ lists.List[int] = (*ArrayList[int])(nil)

// ArrayList wrapper for golang slice.
type ArrayList[E any] struct {
	array   *array.Array[E]
	equalFn fn.EqualFn[E]
}

func NewList[E any](equalFn fn.EqualFn[E]) lists.List[E] {
	return New(equalFn)
}

func New[E any](equalFn fn.EqualFn[E]) *ArrayList[E] {
	return &ArrayList[E]{
		array:   array.New[E](),
		equalFn: equalFn,
	}
}

func (l *ArrayList[E]) IsEmpty() bool {
	return l.array.IsEmpty()
}

func (l *ArrayList[E]) Size() int {
	return l.array.Size()
}

func (l *ArrayList[E]) Clear() {
	l.array.Clear()
}

func (l *ArrayList[E]) Append(element E) {
	l.array.Append(element)
}

func (l *ArrayList[E]) Prepend(element E) {
	l.array.Prepend(element)
}

func (l *ArrayList[E]) At(index int) E {
	return l.array.At(index)
}

func (l *ArrayList[E]) ReplaceAt(index int, element E) E {
	value := l.array.At(index)
	l.array.Set(index, element)
	return value
}

func (l *ArrayList[E]) InsertAt(index int, element E) {
	l.array.Insert(index, element)
}

func (l *ArrayList[E]) RemoveAt(index int) E {
	return l.array.Remove(index)
}

func (l *ArrayList[E]) Swap(i, j int) {
	l.array.Swap(i, j)
}

func (l *ArrayList[E]) Slice(start, end int) lists.List[E] {
	start = max(start, 0)
	end = min(l.array.Size(), end)

	list := New(l.equalFn)
	list.array = l.array.Slice(start, end)
	return list
}

func (l *ArrayList[E]) IndexOf(element E) int {
	return l.array.IndexOf(func(value E) bool {
		return l.equalFn(value, element)
	})
}

func (l *ArrayList[E]) Remove(element E) bool {
	index := l.IndexOf(element)
	if index == -1 {
		return false
	}
	l.RemoveAt(index)
	return true
}

func (l *ArrayList[E]) Contains(elements E) bool {
	return l.IndexOf(elements) != -1
}
