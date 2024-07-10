package array_test

import (
	"goadt/stacks"
	"goadt/stacks/array"
	"testing"
)

func TestArrayAsStack(t *testing.T) {
	stacks.TestStack(t, array.New[int]())
}
