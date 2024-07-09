package linkedlist_test

import (
	"goadt/fn"
	"goadt/queues"
	"goadt/queues/linkedlist"
	"testing"
)

func TestHeap(t *testing.T) {
	queues.TestMinQueue(t, linkedlist.NewPriorityQueue(fn.Compare[int]))
}
