package main_test

import (
	main "goadt"
	"testing"
)

func TestHeap(t *testing.T) {
	testMinQueue(t, main.NewHeap(main.Compare[int]))
}
