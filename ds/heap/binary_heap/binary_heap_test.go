package binaryheap_test

import (
	binaryheap "goadt/ds/heap/binary_heap"
	"goadt/fn"
	"testing"
)

func Test_Heap(t *testing.T) {
	h := binaryheap.New[int](fn.Compare[int])
	if h.IsEmpty() != true {
		t.Errorf("IsEmpty() should return true")
	}

	h.Enqueue(10)
	h.Enqueue(1)
	if h.IsEmpty() != false {
		t.Errorf("IsEmpty() should return false")
	}

	h.Enqueue(2)
	h.Enqueue(3)
	h.Enqueue(4)
	h.Enqueue(5)
	// h.Enqueue(6)

	if h.Size() != 6 {
		t.Errorf("Size() should return 6")
	}

	if h.Peek() != 1 {
		t.Errorf("Peek() should return 1")
	}

	if h.Dequeue() != 1 {
		t.Errorf("Dequeue() should return 1")
	}

	if h.Size() != 5 {
		t.Errorf("Size() should return 5")
	}

	h.Clear()
	if h.IsEmpty() != true {
		t.Errorf("IsEmpty() should return true")
	}

	if h.Size() != 0 {
		t.Errorf("Size() should return 0")
	}
}
