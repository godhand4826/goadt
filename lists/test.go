package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestList check if the given list follows the properties.
func TestList(t *testing.T, list List[int]) {
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
