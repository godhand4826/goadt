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
	queues.TestQueueIsLIFO(t, stackqueue.New(arraylist.New(fn.Equals[int])))
	queues.TestQueueIsLIFO(t, stackqueue.New(linkedlist.New(fn.Equals[int])))
}
