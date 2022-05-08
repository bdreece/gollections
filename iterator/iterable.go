// MIT License
// Copyright (c) 2022 Brian Reece

package iterator

type IntoIterator[T any] interface {
	IntoIterator() Iterator[T]
}

type FromIterator[T any] interface {
	FromIterator(Iterator[T]) error
	Append(T) error
}

type Iterable[T any] interface {
	IntoIterator[T]
	FromIterator[T]
}
