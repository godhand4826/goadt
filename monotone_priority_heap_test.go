package main_test

import (
	"testing"

	main "goadt"
)

func TestHeapQueue(t *testing.T) {
	testQueueIsFIFO(t, main.NewHeapQueue[int]())
}

func TestHeapStack(t *testing.T) {
	testStack(t, main.NewHeapStack[int]())
}
