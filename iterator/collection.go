// MIT License
// Copyright (c) 2022 Brian Reece

package iterator

// Collection provides a common interface for
// all iterable collections.
type Collection[T any] interface {
	Iterator() Iterator[T]
	Collect(...T)
}

// Collect returns a collection filled with the
// items remaining in the iterator. Returns coll,
// error on collection error.
func Collect[T any, C Collection[T]](iter Iterator[T]) (*C, error) {
	var (
		err error
	)
	coll := new(C)
	for {
		item, err := iter.Next()
		if err != nil || item == nil {
			break
		}
		(*coll).Collect(*item)
	}
	return coll, err
}
