// MIT License
// Copyright (c) 2022 Brian Reece

package maps

import (
	"github.com/bdreece/gollections"
	"github.com/bdreece/gollections/iterator"
)

type Map[K comparable, V any] interface {
	gollections.Array[K, V]

	Exists(K) bool
	Keys() iterator.Iterator[K]
	Vals() iterator.Iterator[V]
}
