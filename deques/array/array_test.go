package array_test

import (
	"goadt/deques"
	"goadt/deques/array"
	"testing"
)

func TestArrayAsDeque(t *testing.T) {
	deques.TestDequeIsFIFO(t, array.New[int]())
}
