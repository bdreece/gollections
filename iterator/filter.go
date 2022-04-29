// MIT License
// Copyright (c) 2022 Brian Reece

package iterator

// FilterPredicate is a function that returns true
// if the item should be included in the resulting
// iterator.
type FilterPredicate[T any] func(*T) bool

// Filter transforms an iterator by filtering out
// items based on a predicate function.
type Filter[T any] struct {
	iter Iterator[T]
	pred FilterPredicate[T]
}

// NewFilter constructs a new Filter over an iterator, using pred.
func NewFilter[T any](iter Iterator[T], pred FilterPredicate[T]) *Filter[T] {
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
