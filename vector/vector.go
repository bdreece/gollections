// MIT License
// Copyright (c) 2022 Brian Reece

package vector

import "github.com/bdreece/gollections/errors"

// Vector is a slice of contiguous data.
type Vector[T any] []T

// New creates a new Vector
func New[T any]() *Vector[T] {
	return new(Vector[T])
}

// PushBack appends the Vector with an item.
func (v *Vector[T]) PushBack(item T) {
	*v = append([]T(*v), item)
}

// PushFront prepends the Vector with an item.
func (v *Vector[T]) PushFront(item T) {
	newVec := []T{item}
	*v = append(newVec, []T(*v)...)
}

// PopFront removes and returns an item from the
// front of the Vector. Returns nil, errors.Empty
// if the Vector is empty.
func (v *Vector[T]) PopFront() (*T, error) {
	n := len(*v)
	if n == 0 {
		return nil, errors.Empty{}
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

// PopBack removes and returns an item from the
// back of the Vector. Returns nil, errors.Empty
// if the Vector is empty.
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

// Get returns a pointer to the item at index i.
// Returns nil, errors.IndexOutOfBounds if i < 0 or i > len(v).
func (v Vector[T]) Get(i int) (*T, error) {
	n := len(v)
	if i > n {
		return nil, errors.NewIndexOutOfBounds(i, n)
	}
	return &[]T(v)[i], nil
}

// Set sets the item at index i to value.
// Returns errors.IndexOutOfBounds if i < 0 or i > len(v).
func (v *Vector[T]) Set(i int, value T) error {
	n := len(*v)
	if i >= n {
		return errors.NewIndexOutOfBounds(i, n)
	}
	[]T(*v)[i] = value
	return nil
}

// InsertAfter inserts an item directly after index i.
// Returns errors.IndexOutOfBounds if i < 0 or i > len(v).
func (v *Vector[T]) InsertAfter(i int, item T) error {
	n := len(*v)
	if i >= n {
		return errors.NewIndexOutOfBounds(i, n)
	}
	before := []T(*v)[:i+1]
	after := []T(*v)[i+1:]
	*v = append(before, item)
	*v = append(*v, after...)
	return nil
}

// InsertBefore inserts an item directly before index i.
// Returns errors.IndexOutOfBounds if i < 0 or i > len(v).
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

// Collect inserts a variable number of items
// into the Vector. This method implements part
// of the Iterator interface.
func (v *Vector[T]) Collect(items ...T) {
	for _, value := range items {
		v.PushBack(value)
	}
}

// Iterator returns an iterator over the items
// in the Vector. This method implements part
// of the Iterator interface.
func (v *Vector[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{v, 0}
}
