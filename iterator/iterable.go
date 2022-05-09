// MIT License
// Copyright (c) 2022 Brian Reece

package iterator

import "github.com/bdreece/gollections"

type IntoIterator[T any] interface {
	IntoIterator() Iterator[T]
}

type FromIterator[T any] interface {
	gollections.Collect[T]
	FromIterator(Iterator[T]) error
}

type Iterable[T any] interface {
	IntoIterator[T]
	FromIterator[T]
}
