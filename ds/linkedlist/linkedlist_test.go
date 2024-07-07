package linkedlist_test

import (
	"goadt/ds/linkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	list := linkedlist.New[int]()

	assert.Nil(t, list.First())
	assert.Nil(t, list.Last())
	assert.Equal(t, 0, list.Size())
	assert.True(t, list.IsEmpty())

	list.Append(2)
	list.Append(3)
	list.Prepend(1)
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, 1, list.At(0).GetValue())
	assert.Equal(t, 2, list.At(1).GetValue())
	assert.Equal(t, 3, list.At(2).GetValue())

	node := list.First()
	assert.Equal(t, 1, node.GetValue())
	node = node.Next()
	assert.Equal(t, 2, node.GetValue())
	node = node.Next()
	assert.Equal(t, 3, node.GetValue())
	assert.Nil(t, node.Next())
	node = node.Prev()
	assert.Equal(t, 2, node.GetValue())
	node = node.Prev()
	assert.Equal(t, 1, node.GetValue())
	assert.Nil(t, node.Prev())

	node = list.Last()
	assert.Equal(t, 3, node.GetValue())
	node = node.Prev()
	assert.Equal(t, 2, node.GetValue())
	node = node.Prev()
	assert.Equal(t, 1, node.GetValue())
	assert.Nil(t, node.Prev())
	node = node.Next()
	assert.Equal(t, 2, node.GetValue())
	node = node.Next()
	assert.Equal(t, 3, node.GetValue())
	assert.Nil(t, node.Next())

	node = list.Find(func(value int) bool { return value%2 == 0 })
	assert.Equal(t, 2, node.GetValue())
	node = list.Find(func(value int) bool { return value == 1000 })
	assert.Nil(t, node)

	node = list.First()
	node.SetValue(10)
	assert.Equal(t, 10, node.GetValue())

	list.Clear()
	assert.Nil(t, list.First())
	assert.Nil(t, list.Last())
	assert.Equal(t, 0, list.Size())
	assert.True(t, list.IsEmpty())

	assert.Panics(t, func() { list.At(0) }, "Expected at to panic when list is empty")
}
