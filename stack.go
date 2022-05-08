// MIT License
// Copyright (c) 2022 Brian Reece

package gollections

// Stack provides an interface for
// stack-like collections (i.e. FILO).
type Stack[T any] interface {
    // Push inserts an element into
    // the stack, propagating any errors
    // from the underlying data structure.
    Push(T) error

    // Pop removes and returns an element
    // from the stack, propagating any
    // errors from the underlying data
    // structure.
    Pop() (*T, error)
}
