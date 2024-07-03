package fn

type HashFn[E any, H comparable] func(E) H
