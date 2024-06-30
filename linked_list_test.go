package main_test

import (
	main "goadt"
	"testing"
)

func TestLinkedListAsQueue(t *testing.T) {
	testQueueIsFIFO(t, main.NewLinkedList[int](main.Equals))
}

func TestLinkedListAsStack(t *testing.T) {
	testStack(t, main.NewLinkedList[int](main.Equals))
}

func TestLinkedListAsDeque(t *testing.T) {
	testDequeIsFIFO(t, main.NewLinkedList[int](main.Equals))
}

func TestLinkedListAsList(t *testing.T) {
	testList(t, main.NewLinkedList[int](main.Equals))
}
