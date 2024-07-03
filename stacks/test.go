package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestStack checks if the given stack follows the properties.
func TestStack(t *testing.T, stack Stack[int]) {
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
