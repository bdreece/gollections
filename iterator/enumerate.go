// MIT License
// Copyright (c) 2022 Brian Reece

package iterator

import "fmt"

type EnumerateError struct {
	err   error
	index int
}

func (e EnumerateError) Error() string {
	return fmt.Sprintf("enumerate error, i: (%d), error: \"%s\"", e.index, e.err.Error())
}

type EnumerateItem[T any] struct {
	item  T
	index int
}

type Enumerate[T any] struct {
	iter  Iterator[T]
	index int
}

func NewEnumerate[T any](iter Iterator[T]) *Enumerate[T] {
	return &Enumerate[T]{
		iter:  iter,
		index: 0,
	}
}

func (e *Enumerate[T]) Next() (*EnumerateItem[T], error) {
	item, err := e.iter.Next()
	enum_item := EnumerateItem[T]{
		item:  *item,
		index: e.index,
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
