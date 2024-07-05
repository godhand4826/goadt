package linkedlist

import (
	adt "goadt"
	"goadt/deques"
	"goadt/fn"
	"goadt/lists"
	"goadt/queues"
	"goadt/stacks"
)

var _ lists.List[int] = (*LinkedList[int])(nil)
var _ stacks.Stack[int] = (*LinkedList[int])(nil)
var _ queues.Queue[int] = (*LinkedList[int])(nil)
var _ deques.Deque[int] = (*LinkedList[int])(nil)

// LinkedList Doubly-linked list.
type LinkedList[E any] struct {
	head    *linkedListNode[E] // dummy head will always exist
	size    int
	equalFn fn.EqualFn[E]
}

func New[E any](equalFn fn.EqualFn[E]) *LinkedList[E] {
	return &LinkedList[E]{
		head:    newLinkedListNode(*new(E)),
		equalFn: equalFn,
	}
}

// IsEmpty implements List, Stack, Queue, Deque.
func (l *LinkedList[E]) IsEmpty() bool {
	return l.size == 0
}

// Size implements List, Stack, Queue, Deque.
func (l *LinkedList[E]) Size() int {
	return l.size
}

// Clear implements List, Stack, Queue, Deque.
func (l *LinkedList[E]) Clear() {
	l.head.remove()
	l.size = 0
}

// Append implements List.
func (l *LinkedList[E]) Append(element E) {
	l.PushBack(element)
}

// Prepend implements List.
func (l *LinkedList[E]) Prepend(element E) {
	l.PushFront(element)
}

// At implements List.
func (l *LinkedList[E]) At(index int) E {
	return l.at(index).value
}

// ReplaceAt implements List.
func (l *LinkedList[E]) ReplaceAt(index int, element E) E {
	node := l.at(index)
	value := node.value
	node.value = element
	return value
}

// InsertAt implements List.
func (l *LinkedList[E]) InsertAt(index int, element E) {
	node := newLinkedListNode(element)
	next := l.at(index)
	prev := next.prev

	next.prev = node
	node.next = next
	node.prev = prev
	prev.next = node

	l.size++
}

// RemoveAt implements List.
func (l *LinkedList[E]) RemoveAt(index int) E {
	node := l.at(index)
	value := node.value
	node.remove()
	l.size--
	return value
}

// Swap implement List.
func (l *LinkedList[E]) Swap(i, j int) {
	l.indexMustInRange(i)
	l.indexMustInRange(j)

	n1 := l.at(i)
	n2 := l.at(j)
	n1.value, n2.value = n2.value, n1.value
}

// Slice implements List.
func (l *LinkedList[E]) Slice(start, end int) lists.List[E] {
	start = max(start, 0)
	end = min(l.size, end)

	list := New(l.equalFn)
	node := l.head.next
	for i := 0; i < end; i++ {
		if i >= start {
			list.Append(node.value)
		}
		node = node.next
	}

	return list
}

// IndexOf implements List.
func (l *LinkedList[E]) IndexOf(element E) int {
	index, _ := l.find(element)
	return index
}

// Remove implements List.
func (l *LinkedList[E]) Remove(element E) bool {
	index, node := l.find(element)
	if index == -1 {
		return false
	}

	node.remove()
	l.size--
	return true
}

// Contains implements List.
func (l *LinkedList[E]) Contains(elements E) bool {
	return l.IndexOf(elements) != -1
}

// PushBack implements Deque.
func (l *LinkedList[E]) PushBack(element E) {
	head := l.head
	last := l.head.prev
	node := newLinkedListNode(element)

	last.next = node
	node.prev = last
	node.next = head
	head.prev = node

	l.size++
}

// PushFront implements Deque.
func (l *LinkedList[E]) PushFront(element E) {
	head := l.head
	first := l.head.next
	node := newLinkedListNode(element)

	head.next = node
	node.prev = head
	node.next = first
	first.prev = node

	l.size++
}

// PushBack implements Deque.
func (l *LinkedList[E]) PopBack() E {
	if l.IsEmpty() {
		panic(adt.ErrNoSuchElement)
	}

	node := l.head.prev
	node.remove()

	l.size--

	return node.value
}

// PopFront implements Deque.
func (l *LinkedList[E]) PopFront() E {
	if l.IsEmpty() {
		panic(adt.ErrNoSuchElement)
	}

	node := l.head.next
	node.remove()

	l.size--

	return node.value
}

// Back implements Deque.
func (l *LinkedList[E]) Back() E {
	if l.IsEmpty() {
		panic(adt.ErrNoSuchElement)
	}

	return l.head.prev.value
}

// Front implements Deque.
func (l *LinkedList[E]) Front() E {
	if l.IsEmpty() {
		panic(adt.ErrNoSuchElement)
	}

	return l.head.next.value
}

// Enqueue implements Queue.
func (l *LinkedList[E]) Enqueue(element E) {
	l.PushBack(element)
}

// Dequeue implements Queue.
func (l *LinkedList[E]) Dequeue() E {
	return l.PopFront()
}

// Peek implements Queue.
func (l *LinkedList[E]) Peek() E {
	return l.Front()
}

// Back implements Stack.
func (l *LinkedList[E]) Top() E {
	return l.Back()
}

// Push implements Stack.
func (l *LinkedList[E]) Push(element E) {
	l.PushBack(element)
}

// Pop implements Stack.
func (l *LinkedList[E]) Pop() E {
	return l.PopBack()
}

func (l *LinkedList[E]) at(index int) *linkedListNode[E] {
	l.indexMustInRange(index)

	node := l.head.next
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

func (l *LinkedList[E]) find(element E) (int, *linkedListNode[E]) {
	node := l.head.next
	for i := 0; i < l.size; i++ {
		if l.equalFn(element, node.value) {
			return i, node
		}

		node = node.next
	}

	return -1, nil
}

func (l *LinkedList[E]) indexMustInRange(index int) {
	if !(index >= 0 && index < l.size) {
		panic(adt.ErrIndexOutOfBound)
	}
}

type linkedListNode[E any] struct {
	value E
	next  *linkedListNode[E] // Defined as a pointer due to Go's restriction on recursive types
	prev  *linkedListNode[E]
}

func newLinkedListNode[T any](value T) *linkedListNode[T] {
	node := &linkedListNode[T]{
		value: value,
	}

	node.init()

	return node
}

func (node *linkedListNode[E]) init() {
	node.next = node
	node.prev = node
}

// remove self from list.
func (node *linkedListNode[E]) remove() {
	next := node.next
	prev := node.prev
	next.prev = prev
	prev.next = next

	node.init()
}
