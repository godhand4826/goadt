package fn

func Identity[T any](v T) T {
	return v
}

// On takes two functions f and g, and returns a new function that
// transforms two inputs using g, then combines the transformed outputs using f.
// This is useful for working with comparison or equality functions, such as:
// - On(Compare, GetAge)
// - On(Equals, GetName)
func On[A, B, C any](
	f func(B, B) C, // Combiner function that takes two transformed inputs and returns a result
	g func(A) B, // Transformer function that transforms an input of type A to type B
) func(A, A) C {
	return func(a, b A) C {
		return f(g(a), g(b))
	}
}
