package hashset_test

import (
	"goadt/fn"
	"goadt/lists"
	"goadt/lists/arraylist"
	"goadt/sets"
	"goadt/sets/hashset"
	"testing"
)

func TestHashSet(t *testing.T) {
	var listFactory lists.ListFactory[int] = arraylist.ArrayListFactory[int]{}
	sets.TestSet(t, hashset.NewHashSet(fn.Equals, fn.Identity, listFactory))

	var modTwo = func(element int) int {
		return element % 2
	}
	sets.TestSet(t, hashset.NewHashSet(fn.Equals, modTwo, listFactory))
}
