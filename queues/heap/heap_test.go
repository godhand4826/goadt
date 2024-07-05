package heap_test

import (
	"goadt/fn"
	"goadt/queues"
	"goadt/queues/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	queues.TestMinQueue(t, heap.New(fn.Compare[int]))
}
