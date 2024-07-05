package arraylist_test

import (
	"goadt/deques"
	"goadt/fn"
	"goadt/lists"
	"goadt/lists/arraylist"
	"goadt/queues"
	"goadt/stacks"
	"testing"
)

func TestArrayListAsQueue(t *testing.T) {
	queues.TestQueueIsFIFO(t, arraylist.New(fn.Equals[int]))
}

func TestArrayListAsStack(t *testing.T) {
	stacks.TestStack(t, arraylist.New(fn.Equals[int]))
}

func TestArrayListAsDeque(t *testing.T) {
	deques.TestDequeIsFIFO(t, arraylist.New(fn.Equals[int]))
}

func TestArrayListAsList(t *testing.T) {
	lists.TestList(t, arraylist.New(fn.Equals[int]))
}
