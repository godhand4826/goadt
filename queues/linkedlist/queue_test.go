package linkedlist_test

import (
	"goadt/queues"
	"goadt/queues/linkedlist"
	"testing"
)

func TestLinkedListAsQueue(t *testing.T) {
	queues.TestQueueIsFIFO(t, linkedlist.New[int]())
}
