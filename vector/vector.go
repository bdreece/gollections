// MIT License
// Copyright (c) 2022 Brian Reece

package vector

import (
	"github.com/bdreece/gollections/errors"
	"github.com/bdreece/gollections/iterator"
)

// Vector is a slice of contiguous data.
type Vector[T any] []T

// New creates a new Vector
func New[T any]() *Vector[T] {
	return new(Vector[T])
}

// PushBack appends the Vector with an item.
// This method implements part of the
// gollections.Deque interface.
func (v *Vector[T]) PushBack(item T) {
	*v = append([]T(*v), item)
}

// PopBack removes and returns an item from the
// back of the Vector. Returns nil, errors.Empty
// if the Vector is empty. This method implements
// part of the gollections.Deque interface.
func (v *Vector[T]) PopBack() (*T, error) {
	n := len(*v)
	if n == 0 {
		return nil, errors.IndexOutOfBounds{Index: 0, Bounds: 0}
	}
	val := new(T)
	if n > 0 {
		*val = []T(*v)[n-1]
		*v = []T(*v)[:n-1]
	} else {
		*val = []T(*v)[0]
		*v = *New[T]()
	}
	return val, nil
}

// PeekBack returns a pointer to the last item
// in the vector. Returns nil, errors.Empty if
// the vector is empty. This method implements
// part of the gollections.Deque interface.
func (v Vector[T]) PeekBack() (*T, error) {
	if len(v) == 0 {
		return nil, errors.Empty{}
	}
	return &[]T(v)[len(v)-1], nil
}

// PushFront prepends the Vector with an item.
// This method implements part of the
// gollections.Deque interface.
func (v *Vector[T]) PushFront(item T) {
	newVec := []T{item}
	*v = append(newVec, []T(*v)...)
}

// PopFront removes and returns an item from the
// front of the Vector. Returns nil, errors.Empty
// if the Vector is empty. This method implements
// part of the gollections.Deque interface.
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
		*v = *New[T]()
	}
	return val, nil
}

// PeekFront returns a pointer to the first item
// in the vector. Returns nil, errors.Empty if
// the vector is empty. This method implements
// part of the gollections.Deque interface.
func (v Vector[T]) PeekFront() (*T, error) {
	if len(v) == 0 {
		return nil, errors.Empty{}
	}
	return &[]T(v)[0], nil
}

// Enqueue appends an element to the back
// of the vector. This method implements
// the gollections.Queue interface.
func (v *Vector[T]) Enqueue(elem T) {
	v.PushBack(elem)
}

// Dequeue returns and removes an element
// from the front of the vector. This method
// implements the gollections.Queue interface.
func (v *Vector[T]) Dequeue() (*T, error) {
	return v.PopFront()
}

// Push inserts an element at the front of the
// vector; simply an alias to PushFront. This
// method implements part of the gollections.Stack
// interface.
func (v *Vector[T]) Push(elem T) {
	v.PushFront(elem)
}

// Pop removes and returns an element from the
// front of the vector; simply an alias to
// PopFront. This method implements part of the
// gollections.Stack interface.
func (v *Vector[T]) Pop() (*T, error) {
	return v.PopFront()
}

// Peek returns a pointer to the element at
// the front of the vector; simply an alias
// to Front. This method implements part of
// the gollections.Stack interface.
func (v Vector[T]) Peek() (*T, error) {
	return v.PeekFront()
}

// Get returns a pointer to the item at index i.
// Returns nil, errors.IndexOutOfBounds if i < 0
// or i > len(v). This method implements part of
// the gollections.Array interface.
func (v Vector[T]) Get(i int) (*T, error) {
	n := len(v)
	if i > n {
		return nil, errors.IndexOutOfBounds{Index: i, Bounds: n}
	}
	return &[]T(v)[i], nil
}

// Set sets the item at index i to value.
// Returns errors.IndexOutOfBounds if i < 0
// or i > len(v). This method implements part of
// the gollections.Array interface
func (v *Vector[T]) Set(i int, value T) error {
	n := len(*v)
	if i >= n {
		return errors.IndexOutOfBounds{Index: i, Bounds: n}
	}
	[]T(*v)[i] = value
	return nil
}

// Ins inserts an item directly at index i. Returns
// errors.IndexOutOfBounds if i < 0 or i > len(v). This
// method implements part of the gollections.Array
// interface.
func (v *Vector[T]) Ins(i int, value T) error {
	var (
		before []T
		after  []T
	)
	n := len(*v)
	if i >= n {
		return errors.IndexOutOfBounds{Index: i, Bounds: n}
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

// Del removes and returns the element at
// the specified index, returning
// errors.IndexOutOfBounds if i < 0 or
// i > len(v). This method implements part
// of the gollections.Array interface.
func (v *Vector[T]) Del(i int) (*T, error) {
	n := len(*v)
	if i >= n {
		return nil, errors.IndexOutOfBounds{Index: i, Bounds: n}
	}
	item := []T(*v)[i]
	if i > 0 {
		before := []T(*v)[:i]
		after := []T(*v)[i+1:]
		*v = before
		*v = append(*v, after...)
	} else if n > 0 {
		*v = []T(*v)[1:]
	} else {
		*v = *New[T]()
	}
	return &item, nil
}

// IntoIterator returns an iterator over the items
// in the Vector. This method implements the
// IntoIterator interface.
func (v *Vector[T]) IntoIterator() iterator.Iterator[T] {
	return &Iterator[T]{v, 0}
}

// FromIterator appends the vector with the contents
// of an iterator. This method implements part of the
// FromIterator interface.
func (v *Vector[T]) FromIterator(iter iterator.Iterator[T]) error {
	return iterator.ForEach(
		iter,
		func(item *T) {
			v.PushBack(*item)
		},
	)
}

// Collect appends the vector with a variable number
// of items. This method implements the
// gollections.Collect interface.
func (v *Vector[T]) Collect(elems ...T) {
	for _, elem := range elems {
		v.PushBack(elem)
	}
}

// Clear sets all items in the vector to
// T's zero value.
func (v *Vector[T]) Clear() {
	if n := len(*v); n > 0 {
		for i := 0; i < n; i++ {
			[]T(*v)[i] = *new(T)
		}
	}
}

// Reserve allocates an additional amount of
// space in the vector for more elements. The
// values at these indices are zero-initialized.
func (v *Vector[T]) Reserve(additional int) {
	oldVec := []T(*v)
	*v = make(Vector[T], len(oldVec)+additional)
	for i, item := range oldVec {
		[]T(*v)[i] = item
	}
}

// Truncate reconstructs an empty vector.
func (v *Vector[T]) Truncate() {
	*v = *new(Vector[T])
}
