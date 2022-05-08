// MIT License
// Copyright (c) 2022 Brian Reece

package gollections

import "github.com/bdreece/gollections/iterator"

// Extend describes a collection whose contents
// can be extended with the contents of an iterator.
type Extend[T any] interface {
    // Extend extends the contents of the collections
    // with the contents of an iterator. Returns any
    // error propagated by the iterator, besides when
    // the iterator is exhausted.
	Extend(iterator.Iterator[T]) error
}
