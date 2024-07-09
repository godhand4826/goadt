package linkedlist_test

import (
	"goadt/deques"
	"goadt/deques/linkedlist"
	"testing"
)

func TestLinkedListAsDeque(t *testing.T) {
	deques.TestDequeIsFIFO(t, linkedlist.New[int]())
}
