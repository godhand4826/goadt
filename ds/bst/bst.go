package bst

import "goadt/fn"

type Tree[E any] struct {
	root      *Node[E]
	compareFn fn.CompareFn[E]
	size      int
}

func New[E any](compareFn fn.CompareFn[E]) *Tree[E] {
	return &Tree[E]{
		compareFn: compareFn,
	}
}

func (b *Tree[E]) Insert(value E) {
	node := newNode(b, value)

	var y *Node[E]
	x := b.root
	for x != nil {
		y = x
		if b.compareFn(node.value, x.value) < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}
	node.parent = y
	if y == nil {
		b.root = node
	} else if b.compareFn(node.value, y.value) < 0 {
		y.left = node
	} else {
		y.right = node
	}

	b.size++
}

func (b *Tree[E]) Delete(value E) {
	node := b.Search(value)
	if node == nil {
		return
	}
	node.Delete()
}

func (b *Tree[E]) Search(value E) *Node[E] {
	return b.root.Search(value)
}

func (b *Tree[E]) Maximum() *Node[E] {
	return b.root.Maximum()
}

func (b *Tree[E]) Minimum() *Node[E] {
	return b.root.Minimum()
}

func (b *Tree[E]) Clear() {
	b.root = nil
	b.size = 0
}

func (b *Tree[E]) IsEmpty() bool {
	return b.size == 0
}

func (b *Tree[E]) Size() int {
	return b.size
}

type Node[E any] struct {
	tree   *Tree[E]
	value  E
	parent *Node[E]
	left   *Node[E]
	right  *Node[E]
}

func newNode[E any](tree *Tree[E], value E) *Node[E] {
	return &Node[E]{
		tree:  tree,
		value: value,
	}
}

func (x *Node[E]) Value() E {
	return x.value
}

func (x *Node[E]) Parent() *Node[E] {
	if x == nil {
		return nil
	}
	return x.parent
}

func (x *Node[E]) Left() *Node[E] {
	if x == nil {
		return nil
	}
	return x.left
}

func (x *Node[E]) Right() *Node[E] {
	if x == nil {
		return nil
	}
	return x.right
}

func (x *Node[E]) Delete() {
	if x == nil {
		return
	}

	if x.left == nil {
		x.shift(x.right)
	} else if x.right == nil {
		x.shift(x.left)
	} else {
		y := x.Successor()
		if y.parent != x {
			y.shift(y.right)
			y.right = x.right
			y.right.parent = y
		}
		x.shift(y)
		y.left = x.left
		y.left.parent = y
	}

	x.parent = nil
	x.left = nil
	x.right = nil

	x.tree.size--
}

func (x *Node[E]) shift(y *Node[E]) {
	if x.parent == nil {
		x.tree.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	if y != nil {
		y.parent = x.parent
	}
}

func (x *Node[E]) Search(value E) *Node[E] {
	if x == nil {
		return nil
	}
	fn := x.tree.compareFn
	for x != nil {
		d := fn(value, x.value)
		if d == 0 {
			break
		} else if d < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

func (x *Node[E]) Successor() *Node[E] {
	if x == nil {
		return nil
	}

	if x.right != nil {
		return x.right.Minimum()
	}

	y := x.parent
	for y != nil && x == y.right {
		x = y
		y = y.parent
	}
	return y
}

func (x *Node[E]) Predecessor() *Node[E] {
	if x == nil {
		return nil
	}

	if x.left != nil {
		return x.left.Maximum()
	}

	y := x.parent
	for y != nil && x == y.left {
		x = y
		y = y.parent
	}
	return y
}

func (x *Node[E]) Maximum() *Node[E] {
	if x == nil {
		return nil
	}

	for x.right != nil {
		x = x.right
	}
	return x
}

func (x *Node[E]) Minimum() *Node[E] {
	if x == nil {
		return nil
	}

	for x.left != nil {
		x = x.left
	}
	return x
}
