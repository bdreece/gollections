// MIT License
// Copyright (c) 2022 Brian Reece

package collection

import "github.com/bdreece/gollections/iterator"

// Collection provides a common interface for
// all iterable collections.
type Collection[T any] interface {
	Iterator() iterator.Iterator[T]
	Collect(iterator.Iterator[T]) error
	Append(...T)
	Extend(Collection[T]) error
}
