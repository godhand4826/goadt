package monotone_test

import (
	"goadt/queues"
	"goadt/queues/monotone"
	"goadt/stacks"
	"testing"
)

func TestHeapQueue(t *testing.T) {
	queues.TestQueueIsFIFO(t, monotone.NewHeapQueue[int]())
}

func TestHeapStack(t *testing.T) {
	stacks.TestStack(t, monotone.NewHeapStack[int]())
}
