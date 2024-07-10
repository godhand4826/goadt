package array_test

import (
	"goadt/fn"
	"goadt/lists"
	"goadt/lists/array"
	"testing"
)

func TestArrayAsList(t *testing.T) {
	lists.TestList(t, array.NewList(fn.Equals[int]))
}
