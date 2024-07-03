package deques

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDequeIsFIFO checks if the given FIFO deque follows the properties.
func TestDequeIsFIFO(t *testing.T, deque Deque[int]) {
	deque.Clear()

	var element, size int

	assert.True(t, deque.IsEmpty(), "Expected deque to be empty")
	assert.Equal(t, 0, deque.Size(), "Expected deque size to be zero")

	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushFront(0)

	size = deque.Size()
	assert.Equalf(t, 3, size, "Expected deque size to be 3, got %d", size)
	element = deque.Front()
	assert.Equal(t, 0, element, "Expected front to return 0, got %d", element)
	element = deque.Back()
	assert.Equal(t, 2, element, "Expected back to return 2, got %d", element)

	element = deque.PopFront()
	assert.Equal(t, 0, element, "Expected popFront to return 0, got %d", element)
	element = deque.Front()
	assert.Equal(t, 1, element, "Expected front to return 1, got %d", element)
	element = deque.Back()
	assert.Equal(t, 2, element, "Expected back to return 2, got %d", element)

	element = deque.PopBack()
	assert.Equal(t, 2, element, "Expected popBack to return 2, got %d", element)
	element = deque.Front()
	assert.Equal(t, 1, element, "Expected front to return 1, got %d", element)
	element = deque.Back()
	assert.Equal(t, 1, element, "Expected back to return 1, got %d", element)

	element = deque.PopFront()
	assert.Equal(t, 1, element, "Expected popFront to return 1, got %d", element)
	assert.True(t, deque.IsEmpty(), "Expected deque to be empty")
	assert.Equal(t, 0, deque.Size(), "Expected deque size to be zero")

	deque.PushBack(4)
	deque.PushBack(5)
	deque.PushBack(6)
	deque.Clear()
	assert.True(t, deque.IsEmpty(), "Expected deque to be empty")
	assert.Equal(t, 0, deque.Size(), "Expected deque size to be zero")

	assert.Panics(t, func() { deque.PopBack() }, "Expected popBack to panic when deque is empty")
	assert.Panics(t, func() { deque.PopFront() }, "Expected popFront to panic when deque is empty")
	assert.Panics(t, func() { deque.Front() }, "Expected front to panic when deque is empty")
	assert.Panics(t, func() { deque.Back() }, "Expected back to panic when deque is empty")
}
