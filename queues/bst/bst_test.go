package bst_test

import (
	"goadt/fn"
	"goadt/queues"
	"goadt/queues/bst"
	"testing"
)

func TestBSTQueue(t *testing.T) {
	queues.TestMinQueue(t, bst.New(fn.Compare[int]))
}
