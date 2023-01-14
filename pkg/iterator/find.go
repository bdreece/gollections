package iterator

import "errors"

var ErrNotFound = errors.New("item not found")

type FindFunc[TItem any] func(TItem) bool

func Find[TItem any](iter Iterator[TItem], pred FindFunc[TItem]) (*TItem, error) {
	for val := iter.Next(); val != nil; val = iter.Next() {
		if pred(*val) {
			return val, nil
		}
	}
	return nil, ErrNotFound
}
