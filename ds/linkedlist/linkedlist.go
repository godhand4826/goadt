package linkedlist

import (
	adt "goadt"
)

// List Doubly-linked list.
type List[E any] struct {
	head *Node[E] // dummy head will always exist
	size int
}

func New[E any]() *List[E] {
	list := &List[E]{}
	list.head = newNode(list, *new(E))
	return list
}

func (l *List[E]) First() *Node[E] {
	if l.IsEmpty() {
		return nil
	}
	return l.head.next
}

func (l *List[E]) Last() *Node[E] {
	if l.IsEmpty() {
		return nil
	}
	return l.head.prev
}

func (l *List[E]) Find(fn func(value E) bool) *Node[E] {
	node := l.head.next
	for i := 0; i < l.size; i++ {
		if fn(node.value) {
			return node
		}
		node = node.next
	}
	return nil
}

func (l *List[E]) At(index int) *Node[E] {
	l.indexMustInRange(index)

	node := l.head.next
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

func (l *List[E]) Append(element E) {
	l.head.Prepend(element)
}

func (l *List[E]) Prepend(element E) {
	l.head.Append(element)
}

func (l *List[E]) Size() int {
	return l.size
}

func (l *List[E]) IsEmpty() bool {
	return l.size == 0
}

func (l *List[E]) Clear() {
	l.head.Remove()
	l.size = 0
}

func (l *List[E]) indexMustInRange(index int) {
	if !(index >= 0 && index < l.size) {
		panic(adt.ErrIndexOutOfBound)
	}
}

type Node[E any] struct {
	list  *List[E]
	value E
	next  *Node[E] // Defined as a pointer due to Go's restriction on recursive types
	prev  *Node[E]
}

func newNode[E any](list *List[E], value E) *Node[E] {
	node := &Node[E]{
		list:  list,
		value: value,
	}

	node.init()

	return node
}

func (n *Node[E]) GetValue() E {
	return n.value
}

func (n *Node[E]) SetValue(value E) {
	n.value = value
}

func (n *Node[E]) Next() *Node[E] {
	if n.next == n.list.head {
		return nil
	}
	return n.next
}

func (n *Node[E]) Prev() *Node[E] {
	if n.prev == n.list.head {
		return nil
	}
	return n.prev
}

func (n *Node[E]) Remove() {
	next := n.next
	prev := n.prev
	next.prev = prev
	prev.next = next

	n.init()

	n.list.size--
}

func (n *Node[E]) Append(value E) {
	node := newNode(n.list, value)
	next := n.next

	n.next = node
	node.prev = n
	node.next = next
	next.prev = node

	n.list.size++
}

func (n *Node[E]) Prepend(value E) {
	node := newNode(n.list, value)
	prev := n.prev

	prev.next = node
	node.prev = prev
	node.next = n
	n.prev = node

	n.list.size++
}

func (n *Node[E]) init() {
	n.next = n
	n.prev = n
}
