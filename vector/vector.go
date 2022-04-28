// MIT License
// Copyright (c) 2022 Brian Reece

package vector

import "github.com/bdreece/gollections/errors"

type Vector[T any] []T

func New[T any]() *Vector[T] {
	return new(Vector[T])
}

func (v *Vector[T]) PushBack(value T) {
	*v = append([]T(*v), value)
}

func (v *Vector[T]) PushFront(value T) {
	newVec := []T{value}
	*v = append(newVec, []T(*v)...)
}

func (v *Vector[T]) PopFront() (*T, error) {
	n := len(*v)
	if n == 0 {
		return nil, errors.NewIndexOutOfBounds(0, 0)
	}

	val := new(T)
	*val = []T(*v)[0]
	if n > 0 {
		*v = []T(*v)[1:]
	} else {
		v = New[T]()
	}
	return val, nil
}

func (v *Vector[T]) PopBack() (*T, error) {
	n := len(*v)
	if n == 0 {
		return nil, errors.NewIndexOutOfBounds(0, 0)
	}
	val := new(T)
	if n > 0 {
		*val = []T(*v)[n-1]
		*v = []T(*v)[:n-1]
	} else {
		*val = []T(*v)[0]
		v = New[T]()
	}
	return val, nil
}

func (v Vector[T]) Get(i int) (*T, error) {
	n := len(v)
	if i > n {
		return nil, errors.NewIndexOutOfBounds(i, n)
	}
	return &[]T(v)[i], nil
}

func (v *Vector[T]) Set(i int, value T) error {
	n := len(*v)
	if i >= n {
		return errors.NewIndexOutOfBounds(i, n)
	}
	[]T(*v)[i] = value
	return nil
}

func (v *Vector[T]) InsertAfter(i int, value T) error {
	n := len(*v)
	if i >= n {
		return errors.NewIndexOutOfBounds(i, n)
	}
	before := []T(*v)[:i+1]
	after := []T(*v)[i+1:]
	*v = append(before, value)
	*v = append(*v, after...)
	return nil
}

func (v *Vector[T]) InsertBefore(i int, value T) error {
	var (
		before []T
		after  []T
	)
	n := len(*v)
	if i >= n {
		return errors.NewIndexOutOfBounds(i, n)
	}
	if i > 0 {
		before = []T(*v)[:i]
		after = []T(*v)[i:]
	} else {
		before = []T{}
		after = []T(*v)
	}
	*v = append(before, value)
	*v = append(*v, after...)
	return nil
}

func (v *Vector[T]) Collect(values ...T) {
	for _, value := range values {
		v.PushBack(value)
	}
}

func (v *Vector[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{v}
}

func (v *Vector[T]) ReverseIterator() *ReverseIterator[T] {
	return &ReverseIterator[T]{v}
}
