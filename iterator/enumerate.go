// MIT License
// Copyright (c) 2022 Brian Reece

package iterator

import "fmt"

// EnumerateError is the error type associated
// with an Enumerate.
type EnumerateError struct {
	err   error
	index int
}

// Error returns the error message of an EnumerateError.
func (e EnumerateError) Error() string {
	return fmt.Sprintf("enumerate error, i: (%d), error: \"%s\"", e.index, e.err.Error())
}

// EnumerateItem is the item over which an Enumerate iterates.
type EnumerateItem[T any] struct {
	Item  T
	Index int
}

// Enumerate is an iterator that provides the index of the current
// item being iterated over.
type Enumerate[T any] struct {
	iter  Iterator[T]
	index int
}

// NewEnumerate constructs a new Enumerate over an iterator.
func NewEnumerate[T any](iter Iterator[T]) *Enumerate[T] {
	return &Enumerate[T]{
		iter:  iter,
		index: 0,
	}
}

// Next returns the next item in the Enumerate. Returns
// item, EnumerateError on collection error. This method
// implements the Iterator interface.
func (e *Enumerate[T]) Next() (*EnumerateItem[T], error) {
	item, err := e.iter.Next()
	enum_item := EnumerateItem[T]{
		Item:  *item,
		Index: e.index,
	}
	e.index++
	if err != nil {
		return &enum_item, EnumerateError{
			index: e.index - 1,
			err:   err,
		}
	}
	return &enum_item, nil
}
