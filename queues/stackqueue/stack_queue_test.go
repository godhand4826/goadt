package stackqueue_test

import (
	"goadt/fn"
	"goadt/lists/arraylist"
	"goadt/queues"
	"goadt/queues/stackqueue"
	"goadt/stacks/linkedlist"
	"testing"
)

func TestStackAsLIFOQueue(t *testing.T) {
	queues.TestQueueIsLIFO(t, stackqueue.New(arraylist.New(fn.Equals[int])))
	queues.TestQueueIsLIFO(t, stackqueue.New(linkedlist.NewStack[int]()))
}
