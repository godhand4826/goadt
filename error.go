package adt

import "errors"

var (
	ErrNoSuchElement   = errors.New("no such element")
	ErrIndexOutOfBound = errors.New("index out of bound")
)
