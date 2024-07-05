package queues

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestQueueIsFIFO checks if the given FIFO queue follows the properties.
func TestQueueIsFIFO(t *testing.T, queue Queue[int]) {
	queue.Clear()

	var element, size int

	assert.True(t, queue.IsEmpty(), "Expected queue to be empty")
	assert.Equal(t, 0, queue.Size(), "Expected queue size to be zero")

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	size = queue.Size()
	assert.Equalf(t, 3, size, "Expected queue size to be 3, got %d", size)

	element = queue.Peek()
	assert.Equal(t, 1, element, "Expected peek to return 1, got %d", element)

	element = queue.Dequeue()
	assert.Equalf(t, 1, element, "Expected dequeue to return 1, got %d", element)

	element = queue.Dequeue()
	assert.Equalf(t, 2, element, "Expected dequeue to return 2, got %d", element)

	element = queue.Dequeue()
	assert.Equalf(t, 3, element, "Expected dequeue to return 3, got %d", element)

	assert.True(t, queue.IsEmpty(), "Expected queue to be empty")
	assert.Equal(t, 0, queue.Size(), "Expected queue size to be zero")

	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Enqueue(6)
	queue.Clear()
	assert.True(t, queue.IsEmpty(), "Expected queue to be empty")
	assert.Equal(t, 0, queue.Size(), "Expected queue size to be zero")

	assert.Panics(t, func() { queue.Dequeue() }, "Expected dequeue to panic when queue is empty")
	assert.Panics(t, func() { queue.Peek() }, "Expected peek to panic when queue is empty")
}

// TestQueueIsLIFO checks if the given LIFO queue follows the properties.
func TestQueueIsLIFO(t *testing.T, queue Queue[int]) {
	queue.Clear()

	var element, size int

	assert.True(t, queue.IsEmpty(), "Expected queue to be empty")
	assert.Equal(t, 0, queue.Size(), "Expected queue size to be zero")

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	size = queue.Size()
	assert.Equalf(t, 3, size, "Expected queue size to be 3, got %d", size)

	element = queue.Peek()
	assert.Equal(t, 3, element, "Expected peek to return 3, got %d", element)

	element = queue.Dequeue()
	assert.Equalf(t, 3, element, "Expected dequeue to return 3, got %d", element)

	element = queue.Dequeue()
	assert.Equalf(t, 2, element, "Expected dequeue to return 2, got %d", element)

	element = queue.Dequeue()
	assert.Equalf(t, 1, element, "Expected dequeue to return 1, got %d", element)

	assert.True(t, queue.IsEmpty(), "Expected queue to be empty")
	assert.Equal(t, 0, queue.Size(), "Expected queue size to be zero")

	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Enqueue(6)
	queue.Clear()
	assert.True(t, queue.IsEmpty(), "Expected queue to be empty")
	assert.Equal(t, 0, queue.Size(), "Expected queue size to be zero")

	assert.Panics(t, func() { queue.Dequeue() }, "Expected dequeue to panic when queue is empty")
	assert.Panics(t, func() { queue.Peek() }, "Expected peek to panic when queue is empty")
}

// TestMinQueue checks if the given queue follows the properties.
func TestMinQueue(t *testing.T, queue Queue[int]) {
	queue.Clear()

	assert.Equal(t, 0, queue.Size(), "Expected queue size to be 0")
	assert.True(t, queue.IsEmpty(), "Expected queue to be empty")

	queue.Enqueue(0)
	queue.Enqueue(1)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(2)
	assert.Equal(t, 5, queue.Size(), "Expected queue size to be 5")
	assert.False(t, queue.IsEmpty(), "Expected queue is not empty")

	for i := 0; i < 5; i++ {
		assert.Equalf(t, i, queue.Peek(), "Expected peek to be %d", i)
		value := queue.Dequeue()
		assert.Equalf(t, value, i, "Expected dequeue to be %d", i)
	}
	assert.Equal(t, 0, queue.Size(), "Expected queue size to be 0")
	assert.True(t, queue.IsEmpty(), "Expected queue to be empty")

	queue.Clear()
	assert.Panics(t, func() { queue.Peek() }, "Expected peek to panic when queue is empty")

	queue.Clear()
	const count = 1000
	for i := 0; i < count; i++ {
		queue.Enqueue(rand.Int()) //nolint: gosec
	}

	prev := math.MinInt
	for i := 0; i < count; i++ {
		value := queue.Dequeue()
		assert.True(t, prev <= value, "Expected dequeue elements are ascending")
		prev = value
	}
	assert.Equal(t, 0, queue.Size(), "Expected queue size to be 0")
	assert.True(t, queue.IsEmpty(), "Expected queue to be empty")
}
