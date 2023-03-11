package iterator

import "errors"

var ErrNotFound = errors.New("item not found")

// FindFunc represents the predicate function passed
// to Find. The first item that would result in this
// function returning true will be the returned from
// Find
type FindFunc[TItem any] func(TItem) bool

// Find returns the first item in the given iterator where the predicate
// function returns true
func Find[TItem any](iter Iterator[TItem], pred FindFunc[TItem]) (*TItem, error) {
	for val := iter.Next(); val != nil; val = iter.Next() {
		if pred(*val) {
			return val, nil
		}
	}
	return nil, ErrNotFound
}
