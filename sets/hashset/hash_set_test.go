package hashset_test

import (
	"goadt/fn"
	"goadt/lists/arraylist"
	"goadt/sets"
	"goadt/sets/hashset"
	"testing"
)

func TestHashSet(t *testing.T) {
	sets.TestSet(t, hashset.New(
		fn.Equals,
		fn.Identity,
		arraylist.Factory[int]{},
	))

	sets.TestSet(t, hashset.New(
		fn.Equals,
		func(element int) int { return element % 2 },
		arraylist.Factory[int]{},
	))
}
