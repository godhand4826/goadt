package linkedlist_test

import (
	"goadt/fn"
	"goadt/lists"
	"goadt/lists/linkedlist"
	"testing"
)

func TestLinkedListAsList(t *testing.T) {
	lists.TestList(t, linkedlist.NewList(fn.Equals[int]))
}
