// MIT License
// Copyright (c) 2022 Brian Reece

package gollections

// Deque provides an interface for
// double-ended queue-like collections.
// A double-ended queue can be used in
// place of a queue, stack, or ring
// buffer depending on the use case.
type Deque[T any] interface {
	// PushFront prepends the deque
	// with an element, propagating
	// any errors from the underlying
	// data structure.
	PushFront(T) error

	// PopFront returns and removes the
	// element at the front of the deque,
	// propagating any errors from the
	// underlying data structure. Note
	// that the value returned is not
	// a pointer into the collection,
	// as the item will have been removed.
	PopFront() (*T, error)

	// PeekFront returns the first
	// element in the deque, if it
	// exists. Note that the returned
	// value may not necessarily be
	// a pointer into the collection;
	// see concrete implementations for
	// details.
	PeekFront() (*T, error)

	// PushBack appends the deque with
	// an element, propagating any
	// errors from the underlying data
	// structure.
	PushBack(T) error

	// PopFront returns and removes the
	// element at the back of the deque,
	// propagating any errors from the
	// underlying data structure. Note
	// that the value returned is not a
	// pointer into the collection,
	// as the item will have been removed.
	PopBack() (*T, error)

	// PeekBack returns the last
	// element in the deque, if it
	// exists. Note that the returned
	// value may not necessarily be
	// a pointer into the collection;
	// see concrete implementations for
	// details.
	PeekBack() (*T, error)
}
