// MIT License
// Copyright (c) 2022 Brian Reece

package vector

// Iterator provides an iterator over the
// items in a Vector.
type Iterator[T any] struct {
	*Vector[T]
	int
}

// Next returns the next item in the iterator.
// This method implements the iterator.Iterator
// interface. Returns nil, errors.IndexOutOfBounds
// after the last item.
func (iter *Iterator[T]) Next() (*T, error) {
	item, err := iter.Vector.Get(iter.int)
	iter.int++
	return item, err
}

// Prev returns the previous item in the iterator
// This method implements the iterator.Reverse
// interface. Returns nil, errors.IndexOutOfBounds
// after the first item.
func (iter *Iterator[T]) Prev() (*T, error) {
	item, err := iter.Vector.Get(iter.int)
	iter.int--
	return item, err
}
