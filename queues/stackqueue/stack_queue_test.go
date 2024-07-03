package stackqueue_test

import (
	"goadt/fn"
	"goadt/lists/arraylist"
	"goadt/lists/linkedlist"
	"goadt/queues"
	"goadt/queues/stackqueue"
	"testing"
)

func TestStackAsLIFOQueue(t *testing.T) {
	queues.TestQueueIsLIFO(t, stackqueue.NewStackQueue(arraylist.NewArrayList(fn.Equals[int])))
	queues.TestQueueIsLIFO(t, stackqueue.NewStackQueue(linkedlist.NewLinkedList(fn.Equals[int])))
}
