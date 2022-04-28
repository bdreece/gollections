// MIT License
// Copyright (c) 2022 Brian Reece

package vector

type Iterator[T any] struct {
	*Vector[T]
}

func (iter *Iterator[T]) Next() (*T, error) {
	return iter.Vector.PopFront()
}

type ReverseIterator[T any] struct {
	*Vector[T]
}

func (iter *ReverseIterator[T]) Next() (*T, error) {
	return iter.Vector.PopBack()
}
