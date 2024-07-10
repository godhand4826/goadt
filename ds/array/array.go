package array

import (
	"slices"
)

type Array[E any] struct {
	elements []E
}

func New[E any]() *Array[E] {
	return &Array[E]{}
}

func (a *Array[E]) Append(elements ...E) {
	a.elements = append(a.elements, elements...)
}

func (a *Array[E]) Prepend(elements ...E) {
	a.elements = append(elements, a.elements...)
}

func (a *Array[E]) At(index int) E {
	return a.elements[index]
}

func (a *Array[E]) Set(index int, element E) {
	a.elements[index] = element
}

func (a *Array[E]) IndexOf(fn func(value E) bool) int {
	return slices.IndexFunc(a.elements, fn)
}

func (a *Array[E]) Insert(index int, elements ...E) {
	a.elements = slices.Insert(a.elements, index, elements...)
}

func (a *Array[E]) Remove(index int) E {
	value := a.elements[index]
	a.elements = slices.Delete(a.elements, index, index+1)
	return value
}

func (a *Array[E]) Swap(i, j int) {
	a.elements[i], a.elements[j] = a.elements[j], a.elements[i]
}

func (a *Array[E]) Slice(i, j int) *Array[E] {
	elements := make([]E, j-i)
	copy(elements, a.elements[i:j])

	return &Array[E]{
		elements: elements,
	}
}

func (a *Array[E]) IsEmpty() bool {
	return len(a.elements) == 0
}

func (a *Array[E]) Size() int {
	return len(a.elements)
}

func (a *Array[E]) Clear() {
	a.elements = nil
}
