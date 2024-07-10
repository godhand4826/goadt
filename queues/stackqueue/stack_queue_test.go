package stackqueue_test

import (
	"goadt/queues"
	"goadt/queues/stackqueue"
	"goadt/stacks/array"
	"goadt/stacks/linkedlist"
	"testing"
)

func TestStackAsLIFOQueue(t *testing.T) {
	queues.TestQueueIsLIFO(t, stackqueue.New(array.New[int]()))
	queues.TestQueueIsLIFO(t, stackqueue.New(linkedlist.New[int]()))
}
