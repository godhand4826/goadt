package linkedlist_test

import (
	"goadt/deques"
	"goadt/fn"
	"goadt/lists"
	"goadt/lists/linkedlist"
	"goadt/queues"
	"goadt/stacks"
	"testing"
)

func TestLinkedListAsQueue(t *testing.T) {
	queues.TestQueueIsFIFO(t, linkedlist.NewLinkedList(fn.Equals[int]))
}

func TestLinkedListAsStack(t *testing.T) {
	stacks.TestStack(t, linkedlist.NewLinkedList(fn.Equals[int]))
}

func TestLinkedListAsDeque(t *testing.T) {
	deques.TestDequeIsFIFO(t, linkedlist.NewLinkedList(fn.Equals[int]))
}

func TestLinkedListAsList(t *testing.T) {
	lists.TestList(t, linkedlist.NewLinkedList(fn.Equals[int]))
}
