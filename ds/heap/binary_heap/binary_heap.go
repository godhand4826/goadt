package binaryheap

import (
	adt "goadt"
	"goadt/fn"
	linkedlistQueue "goadt/queues/linkedlist"
)

type Node[E any] struct {
	value E
	left  *Node[E]
	right *Node[E]
}

// BinaryHeap is a priority queue based on a binary heap.
type BinaryHeap[E any] struct {
	compare fn.CompareFn[E]
	root    *Node[E]
	size    int
}

// New returns a new BinaryHeap.
func New[E any](compare fn.CompareFn[E]) *BinaryHeap[E] {
	return &BinaryHeap[E]{
		compare: compare,
	}
}

// Enqueue adds an element to the queue.
func (h *BinaryHeap[E]) Enqueue(element E) {
	newNode := &Node[E]{value: element}
	if h.root == nil {
		h.root = newNode
		h.size = 1
		return
	}

	// find the last node
	queue := linkedlistQueue.New[*Node[E]]()
	queue.Enqueue(h.root)
	for !queue.IsEmpty() {
		node := queue.Dequeue()
		if node.left == nil {
			node.left = newNode
			break
		}
		if node.right == nil {
			node.right = newNode
			break
		}
		queue.Enqueue(node.left)
		queue.Enqueue(node.right)
	}

	h.size++
	// heapify
	h.bubbleUp(newNode)
}

func (h *BinaryHeap[E]) bubbleUp(node *Node[E]) {
	if node == nil {
		return
	}
	parent := h.getParent(node)
	if parent == nil {
		return
	}
	if h.compare(node.value, parent.value) < 0 {
		node.value, parent.value = parent.value, node.value
		h.bubbleUp(parent)
	}
}

func (h *BinaryHeap[E]) getParent(node *Node[E]) *Node[E] {
	queue := linkedlistQueue.New[*Node[E]]()
	queue.Enqueue(h.root)
	for !queue.IsEmpty() {
		parent := queue.Dequeue()
		if parent.left == node || parent.right == node {
			return parent
		}
		if parent.left != nil {
			queue.Enqueue(parent.left)
		}
		if parent.right != nil {
			queue.Enqueue(parent.right)
		}
	}
	return nil
}

// Dequeue removes and returns the element with the highest priority.
func (h *BinaryHeap[E]) Dequeue() E {
	if h.IsEmpty() {
		panic(adt.ErrNoSuchElement)
	}
	// find the last node
	queue := linkedlistQueue.New[*Node[E]]()
	queue.Enqueue(h.root)
	var lastNode *Node[E]
	for !queue.IsEmpty() {
		lastNode = queue.Dequeue()
		if lastNode.left != nil {
			queue.Enqueue(lastNode.left)
		}
		if lastNode.right != nil {
			queue.Enqueue(lastNode.right)
		}
	}

	res := h.root.value
	h.root.value = lastNode.value
	// remove the last node
	queue.Enqueue(h.root)
	for !queue.IsEmpty() {
		node := queue.Dequeue()
		if node.left == lastNode {
			node.left = nil
			break
		}
		if node.right == lastNode {
			node.right = nil
			break
		}
		if node.left != nil {
			queue.Enqueue(node.left)
		}
		if node.right != nil {
			queue.Enqueue(node.right)
		}
	}
	h.size--
	// heapify
	h.bubbleDown(h.root)

	return res
}

func (h *BinaryHeap[E]) bubbleDown(node *Node[E]) {
	if node == nil {
		return
	}
	left := node.left
	right := node.right
	if left == nil && right == nil {
		return
	}
	if left != nil && h.compare(left.value, node.value) < 0 {
		node.value, left.value = left.value, node.value
		h.bubbleDown(left)
	}
	if right != nil && h.compare(right.value, node.value) < 0 {
		node.value, right.value = right.value, node.value
		h.bubbleDown(right)
	}
}

// Peek returns the element with the highest priority.
func (h *BinaryHeap[E]) Peek() E {
	if h.IsEmpty() {
		panic(adt.ErrNoSuchElement)
	}
	return h.root.value
}

// IsEmpty returns true if the queue is empty.
func (h *BinaryHeap[E]) IsEmpty() bool {
	return h.size == 0
}

// Size returns the number of elements in the queue.
func (h *BinaryHeap[E]) Size() int {
	return h.size
}

// Clear removes all elements from the queue.
func (h *BinaryHeap[E]) Clear() {
	h.root = nil
	h.size = 0
}
