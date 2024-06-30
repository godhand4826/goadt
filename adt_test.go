package main_test

import (
	"testing"

	main "goadt"

	"github.com/stretchr/testify/assert"
)

// testStack checks if the given stack follows the properties.
func testStack(t *testing.T, stack main.Stack[int]) {
	stack.Clear()

	var element, size int

	assert.True(t, stack.IsEmpty(), "Expected stack to be empty")
	assert.Equal(t, 0, stack.Size(), "Expected stack size to be zero")

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	size = stack.Size()
	assert.Equalf(t, 3, size, "Expected stack size to be 3, got %d", size)

	element = stack.Top()
	assert.Equal(t, 3, element, "Expected peek to return 3, got %d", element)

	element = stack.Pop()
	assert.Equalf(t, 3, element, "Expected pop to return 3, got %d", element)

	element = stack.Pop()
	assert.Equalf(t, 2, element, "Expected pop to return 2, got %d", element)

	element = stack.Pop()
	assert.Equalf(t, 1, element, "Expected pop to return 1, got %d", element)

	assert.True(t, stack.IsEmpty(), "Expected stack to be empty")
	assert.Equal(t, 0, stack.Size(), "Expected stack size to be zero")

	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Clear()
	assert.True(t, stack.IsEmpty(), "Expected stack to be empty")
	assert.Equal(t, 0, stack.Size(), "Expected stack size to be zero")

	assert.Panics(t, func() { stack.Pop() }, "Expected pop to panic when stack is empty")
	assert.Panics(t, func() { stack.Top() }, "Expected top to panic when stack is empty")
}

// testQueueIsFIFO checks if the given FIFO queue follows the properties.
func testQueueIsFIFO(t *testing.T, queue main.Queue[int]) {
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

// testQueueIsLIFO checks if the given LIFO queue follows the properties.
func testQueueIsLIFO(t *testing.T, queue main.Queue[int]) {
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

// testDequeIsFIFO checks if the given FIFO deque follows the properties.
func testDequeIsFIFO(t *testing.T, deque main.Deque[int]) {
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

// testList check if the given list follows the properties.
func testList(t *testing.T, list main.List[int]) {
	list.Clear()
	assert.True(t, list.IsEmpty(), "Expected list to be empty")
	assert.Equal(t, 0, list.Size(), "Expected list size to be zero")

	list.Append(1)
	assert.Equal(t, 1, list.Size(), "Expected list size to be 1")
	assert.Equal(t, 1, list.At(0), "Expected list[0] to be 1")
	assert.True(t, list.Contains(1), "Expected list to contain 1")
	assert.False(t, list.Contains(2), "Expected list not to contain 2")

	list.Prepend(2)
	assert.Equal(t, 2, list.Size(), "Expected list size to be 2")
	assert.Equal(t, 2, list.At(0), "Expected list[0] to be 2")
	assert.Equal(t, 1, list.At(1), "Expected list[1] to be 1")
	assert.True(t, list.Contains(1), "Expected list to contain 1")
	assert.True(t, list.Contains(2), "Expected list to contain 2")

	assert.Equal(t, 0, list.IndexOf(2), "Expected index of 2 to be 0")
	assert.Equal(t, 1, list.IndexOf(1), "Expected index of 1 to be 1")
	assert.Equal(t, -1, list.IndexOf(3), "Expected index of 3 to be -1")

	list.ReplaceAt(1, 3)
	assert.Equal(t, 3, list.At(1), "Expected list[1] to be 3")
	assert.True(t, list.Contains(3), "Expected list to contain 3")
	assert.False(t, list.Contains(1), "Expected list not to contain 1")

	removedElement := list.RemoveAt(0)
	assert.Equal(t, 2, removedElement, "Expected removed element to be 2")
	assert.Equal(t, 1, list.Size(), "Expected list size to be 1")
	assert.Equal(t, 3, list.At(0), "Expected list[0] to be 3")

	list.Clear()
	assert.True(t, list.IsEmpty(), "Expected list to be empty after clear")
	assert.Equal(t, 0, list.Size(), "Expected list size to be zero after clear")

	list.Append(10)
	list.Append(20)
	list.Append(30)
	sliced := list.Slice(1, 1)
	assert.True(t, sliced.IsEmpty(), "Expected sliced list to be empty")
	assert.Equal(t, 0, sliced.Size(), "Expected sliced list size to be zero")

	sliced = list.Slice(1, 2)
	assert.Equal(t, 1, sliced.Size(), "Expected sliced list size to be 1")
	assert.Equal(t, sliced.At(0), list.At(1), "Expected sliced[0] equal to list[1]")

	sliced = list.Slice(-1, 2)
	assert.Equal(t, 2, sliced.Size(), "Expected sliced list size to be 2")
	assert.Equal(t, sliced.At(0), list.At(0), "Expected sliced[0] equal to list[0]")
	assert.Equal(t, sliced.At(1), list.At(1), "Expected sliced[1] equal to list[1]")

	sliced = list.Slice(0, 100)
	assert.Equal(t, 3, sliced.Size(), "Expected sliced list size to be 3")
	assert.Equal(t, sliced.At(0), list.At(0), "Expected sliced[0] equal to list[0]")
	assert.Equal(t, sliced.At(1), list.At(1), "Expected sliced[1] equal to list[1]")
	assert.Equal(t, sliced.At(2), list.At(2), "Expected sliced[2] equal to list[2]")

	list.Clear()
	list.Append(1)
	list.Append(2)
	list.Append(1)
	changed := list.Remove(1)
	assert.True(t, changed, "Expected list is changed")
	assert.Equal(t, 2, list.Size(), "Expected list size to be 2")
	assert.Equal(t, 2, list.At(0), "Expected list[0] equal to 2")
	assert.Equal(t, 1, list.At(1), "Expected list[1] equal to 1")

	changed = list.Remove(1)
	assert.True(t, changed, "Expected list is changed")
	assert.Equal(t, 1, list.Size(), "Expected list size to be 1")
	assert.Equal(t, 2, list.At(0), "Expected list[0] equal to 2")

	changed = list.Remove(3)
	assert.False(t, changed, "Expected list is not changed")
	assert.Equal(t, 1, list.Size(), "Expected list size to be 1")
	assert.Equal(t, 2, list.At(0), "Expected list[0] equal to 2")

	list.Clear()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.InsertAt(0, 4)
	list.InsertAt(1, 5)
	assert.Equal(t, 5, list.Size(), "Expected list size to be 5")
	assert.Equal(t, 4, list.At(0), "Expected list[0] equal to be 4")
	assert.Equal(t, 5, list.At(1), "Expected list[1] equal to be 5")
	assert.Equal(t, 1, list.At(2), "Expected list[2] equal to be 1")
	assert.Equal(t, 2, list.At(3), "Expected list[3] equal to be 2")
	assert.Equal(t, 3, list.At(4), "Expected list[4] equal to be 3")

	list.Clear()
	assert.Panics(t, func() { list.At(0) }, "Expected at to panic when list is empty")
}

// testSet checks if the given set follows the properties.
func testSet(t *testing.T, set main.Set[int]) {
	set.Clear()

	assert.True(t, set.IsEmpty(), "Expected set to be empty")
	assert.Equal(t, 0, set.Size(), "Expected set size to be zero")

	set.Add(1)
	assert.Equal(t, 1, set.Size(), "Expected set size to be 1")

	set.Add(2)
	assert.Equal(t, 2, set.Size(), "Expected set size to be 2")

	set.Add(2)
	assert.Equal(t, 2, set.Size(), "Expected set size to be 2")

	set.Add(10)
	set.Clear()
	assert.True(t, set.IsEmpty(), "Expected set to be empty")
	assert.Equal(t, 0, set.Size(), "Expected set size to be zero")

	set.Add(10)
	set.Add(20)
	set.Add(30)
	set.Clear()
	assert.True(t, set.IsEmpty(), "Expected set to be empty")
	assert.Equal(t, 0, set.Size(), "Expected set size to be zero")

	assert.False(t, set.Contains(0), "Expected set not contains 0")
	assert.False(t, set.Contains(10), "Expected set not contains 10")
	assert.False(t, set.Contains(20), "Expected set not contains 20")

	set.Add(10)
	assert.False(t, set.Contains(0), "Expected set not contains 0")
	assert.True(t, set.Contains(10), "Expected set contains 10")
	assert.False(t, set.Contains(20), "Expected set not contains 20")
	assert.Equal(t, 1, set.Size(), "Expected set size to be 1")

	set.Add(10)
	assert.False(t, set.Contains(0), "Expected set not contains 0")
	assert.True(t, set.Contains(10), "Expected set contains 10")
	assert.False(t, set.Contains(20), "Expected set not contains 20")
	assert.Equal(t, 1, set.Size(), "Expected set size to be 1")

	set.Add(20)
	assert.False(t, set.Contains(0), "Expected set not contains 0")
	assert.True(t, set.Contains(10), "Expected set contains 10")
	assert.True(t, set.Contains(20), "Expected set contains 20")
	assert.Equal(t, 2, set.Size(), "Expected set size to be 2")

	set.Remove(10)
	assert.False(t, set.Contains(0), "Expected set not contains 0")
	assert.False(t, set.Contains(10), "Expected set not contains 10")
	assert.True(t, set.Contains(20), "Expected set contains 20")
	assert.Equal(t, 1, set.Size(), "Expected set size to be 1")

	set.Remove(0)
	assert.False(t, set.Contains(0), "Expected set not contains 0")
	assert.False(t, set.Contains(10), "Expected set not contains 10")
	assert.True(t, set.Contains(20), "Expected set contains 20")
	assert.Equal(t, 1, set.Size(), "Expected set size to be 1")

	set.Remove(20)
	assert.False(t, set.Contains(0), "Expected set not contains 0")
	assert.False(t, set.Contains(10), "Expected set not contains 10")
	assert.False(t, set.Contains(20), "Expected set not contains 20")
	assert.Equal(t, 0, set.Size(), "Expected set size to be 0")
	assert.True(t, set.IsEmpty(), "Expected set to be empty")
}
