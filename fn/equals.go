package fn

type EqualFn[E any] func(E, E) bool

func Equals[E comparable](a E, b E) bool {
	return a == b
}
