// MIT License
// Copyright (c) 2022 Brian Reece

package iterator

// Collection provides a common interface for
// all iterable collections.
type Collection[T any] interface {
	Iterator() Iterator[T]
	Collect(Iterator[T]) error
	Append(...T)
	Extend(Collection[T]) error
}
