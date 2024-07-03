package sets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSet checks if the given set follows the properties.
func TestSet(t *testing.T, set Set[int]) {
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
