package linkedlist_test

import (
	"goadt/stacks"
	"goadt/stacks/linkedlist"
	"testing"
)

func TestLinkedListAsStack(t *testing.T) {
	stacks.TestStack(t, linkedlist.NewStack[int]())
}
