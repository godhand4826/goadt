package main

import "cmp"

type EqualFn[E any] func(E, E) bool

func Equals[E comparable](a E, b E) bool {
	return a == b
}

type CompareFn[E any] func(E, E) int

func Compare[E cmp.Ordered](a E, b E) int {
	return cmp.Compare(a, b)
}

type HashFn[E any, H comparable] func(E) H

func Identity[T any](v T) T {
	return v
}
