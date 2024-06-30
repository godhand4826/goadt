package main

import "slices"

var _ List[int] = (*ArrayList[int])(nil)
var _ Stack[int] = (*ArrayList[int])(nil)
var _ Queue[int] = (*ArrayList[int])(nil)
var _ Deque[int] = (*ArrayList[int])(nil)

// ArrayList wrapper for golang slice.
type ArrayList[E any] struct {
	elements []E
	equalFn  EqualFn[E]
}

func NewArrayList[E any](equalFn EqualFn[E]) *ArrayList[E] {
	return &ArrayList[E]{
		equalFn: equalFn,
	}
}

// IsEmpty implements List, Stack, Queue, Deque.
func (a *ArrayList[E]) IsEmpty() bool {
	return len(a.elements) == 0
}

// Size implements List, Stack, Queue, Deque.
func (a *ArrayList[E]) Size() int {
	return len(a.elements)
}

// Clear implements List, Stack, Queue, Deque.
func (a *ArrayList[E]) Clear() {
	a.elements = nil
}

// NewList implements List.
func (a *ArrayList[E]) NewList() List[E] {
	return NewArrayList(a.equalFn)
}

// Append implements List.
func (a *ArrayList[E]) Append(element E) {
	a.elements = append(a.elements, element)
}

// Prepend implements List.
func (a *ArrayList[E]) Prepend(element E) {
	a.elements = append([]E{element}, a.elements...)
}

// At implements List.
func (a *ArrayList[E]) At(index int) E {
	a.indexMustInRange(index)
	return a.elements[index]
}

// ReplaceAt implements List.
func (a *ArrayList[E]) ReplaceAt(index int, element E) E {
	a.indexMustInRange(index)
	value := a.elements[index]
	a.elements[index] = element
	return value
}

// InsertAt implements List.
func (a *ArrayList[E]) InsertAt(index int, element E) {
	a.indexMustInRange(index)
	a.elements = slices.Insert(a.elements, index, element)
}

// RemoveAt implements List.
func (a *ArrayList[E]) RemoveAt(index int) E {
	a.indexMustInRange(index)

	value := a.elements[index]
	a.elements = slices.Delete(a.elements, index, index+1)
	return value
}

// Slice implements List.
func (a *ArrayList[E]) Slice(start, end int) List[E] {
	start = max(start, 0)
	end = min(len(a.elements), end)

	list := NewArrayList(a.equalFn)
	if end-start <= 0 {
		return list
	}

	list.elements = a.elements[start:end]
	return list
}

// IndexOf implements List.
func (a *ArrayList[E]) IndexOf(element E) int {
	return slices.IndexFunc(a.elements, func(e E) bool {
		return a.equalFn(e, element)
	})
}

// Remove implements List.
func (a *ArrayList[E]) Remove(element E) bool {
	index := a.IndexOf(element)
	if index == -1 {
		return false
	}
	a.RemoveAt(index)
	return true
}

// Contains implements List.
func (a *ArrayList[E]) Contains(elements E) bool {
	return a.IndexOf(elements) != -1
}

// PushBack implements Deque.
func (a *ArrayList[E]) PushBack(element E) {
	a.Append(element)
}

// PushFront implements Deque.
func (a *ArrayList[E]) PushFront(element E) {
	a.Prepend(element)
}

// PushBack implements Deque.
func (a *ArrayList[E]) PopBack() E {
	if a.IsEmpty() {
		panic(ErrNoSuchElement)
	}
	element := a.elements[len(a.elements)-1]
	a.elements = a.elements[:len(a.elements)-1]
	return element
}

// PopFront implements Deque.
func (a *ArrayList[E]) PopFront() E {
	if a.IsEmpty() {
		panic(ErrNoSuchElement)
	}

	element := a.elements[0]
	a.elements = a.elements[1:]
	return element
}

// Back implements Deque.
func (a *ArrayList[E]) Back() E {
	if a.IsEmpty() {
		panic(ErrNoSuchElement)
	}

	return a.elements[len(a.elements)-1]
}

// Front implements Deque.
func (a *ArrayList[E]) Front() E {
	if a.IsEmpty() {
		panic(ErrNoSuchElement)
	}

	return a.elements[0]
}

// Enqueue implements Queue.
func (a *ArrayList[E]) Enqueue(element E) {
	a.PushBack(element)
}

// Dequeue implements Queue.
func (a *ArrayList[E]) Dequeue() E {
	return a.PopFront()
}

// Peek implements Queue.
func (a *ArrayList[E]) Peek() E {
	return a.Front()
}

// Back implements Stack.
func (a *ArrayList[E]) Top() E {
	return a.Back()
}

// Push implements Stack.
func (a *ArrayList[E]) Push(element E) {
	a.PushBack(element)
}

// Pop implements Stack.
func (a *ArrayList[E]) Pop() E {
	return a.PopBack()
}

func (a *ArrayList[E]) indexMustInRange(index int) {
	if !(index >= 0 && index < len(a.elements)) {
		panic(ErrIndexOutOfBound)
	}
}
