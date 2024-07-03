package fn

import "cmp"

type CompareFn[E any] func(E, E) int

func Compare[E cmp.Ordered](a E, b E) int {
	return cmp.Compare(a, b)
}
