package main_test

import (
	"testing"

	main "goadt"
)

func TestStackAsLIFOQueue(t *testing.T) {
	testQueueIsLIFO(t, main.NewStackQueue(main.NewArrayList[int](main.Equals)))
	testQueueIsLIFO(t, main.NewStackQueue(main.NewLinkedList[int](main.Equals)))
}
