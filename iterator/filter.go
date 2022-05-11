// MIT License
// Copyright (c) 2022 Brian Reece

package iterator

// FilterPredicate is a function that returns true
// if the item should be included in the resulting
// iterator.
type FilterFunc[T any] func(*T) bool

// Filter transforms an iterator by filtering out
// items based on a predicate function.
type Filter[T any] struct {
	iter Iterator[T]
	pred FilterFunc[T]
}

// NewFilter constructs a new Filter over an iterator, using pred.
func NewFilter[T any](iter Iterator[T], pred FilterFunc[T]) *Filter[T] {
	return &Filter[T]{iter, pred}
}

// Next returns the next item in the Filter. Returns nil, error
// on collection error.
func (f *Filter[T]) Next() (*T, error) {
	for {
		item, err := f.iter.Next()
		if item == nil {
			return nil, nil
		}
		if !(f.pred)(item) {
			continue
		}
		if err != nil {
			return item, err
		}
		return item, nil
	}
}

// Any returns true if for any item in the iterator,
// pred returns true. Returns false, error upon collection
// error.
func Any[T any](iter Iterator[T], pred FilterFunc[T]) (bool, error) {
	for {
		item, err := iter.Next()
		if item == nil {
			break
		}
		if err != nil {
			return false, err
		}
		if (pred)(item) {
			return true, nil
		}
	}
	return false, nil
}

// All returns true if for all items in the iterator,
// pred returns true. Returns false, error upon collection
// error.
func All[T any](iter Iterator[T], pred FilterFunc[T]) (bool, error) {
	for {
		item, err := iter.Next()
		if item == nil {
			break
		}
		if err != nil {
			return true, err
		}
		if !(pred)(item) {
			return false, nil
		}
	}
	return true, nil
}

// None returns true if for all items in the iterator,
// pred returns false. Returns false, error upon collection
// error.
func None[T any](iter Iterator[T], pred FilterFunc[T]) (bool, error) {
	for {
		item, err := iter.Next()
		if item == nil {
			break
		}
		if err != nil {
			return true, err
		}
		if (pred)(item) {
			return false, err
		}
	}
	return true, nil
}
