package hashset_test

import (
	"goadt/fn"
	"goadt/lists/arraylist"
	"goadt/sets"
	"goadt/sets/hashset"
	"testing"
)

func TestHashSet(t *testing.T) {
	sets.TestSet(t, hashset.NewHashSet(fn.Equals, fn.Identity, arraylist.NewArrayList(fn.Equals[int])))
	sets.TestSet(t, hashset.NewHashSet(fn.Equals, func(element int) int {
		return element % 2
	}, arraylist.NewArrayList(fn.Equals[int])))
}
