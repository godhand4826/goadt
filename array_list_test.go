package main_test

import (
	"testing"

	main "goadt"
)

func TestArrayListAsQueue(t *testing.T) {
	testQueueIsFIFO(t, main.NewArrayList[int](main.Equals))
}

func TestArrayListAsStack(t *testing.T) {
	testStack(t, main.NewArrayList[int](main.Equals))
}

func TestArrayListAsDeque(t *testing.T) {
	testDequeIsFIFO(t, main.NewArrayList[int](main.Equals))
}

func TestArrayListAsList(t *testing.T) {
	testList(t, main.NewArrayList[int](main.Equals))
}
