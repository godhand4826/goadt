package array_test

import (
	"goadt/queues"
	"goadt/queues/array"
	"testing"
)

func TestArrayAsQueue(t *testing.T) {
	queues.TestQueueIsFIFO(t, array.New[int]())
}
