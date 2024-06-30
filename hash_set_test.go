package main_test

import (
	main "goadt"
	"testing"
)

func TestHashSet(t *testing.T) {
	testSet(t, main.NewHashSet(main.Equals, main.Identity, main.NewArrayList[int](main.Equals)))
	testSet(t, main.NewHashSet(main.Equals, func(element int) int {
		return element % 2
	}, main.NewArrayList[int](main.Equals)))
}
