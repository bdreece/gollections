// MIT License
// Copyright (c) 2022 Brian Reece

package iterator

type FilterPredicate[T any] func(*T) bool

type Filter[T any] struct {
	iter Iterator[T]
	pred FilterPredicate[T]
}

func NewFilter[T any](iter Iterator[T], pred FilterPredicate[T]) *Filter[T] {
	return &Filter[T]{iter, pred}
}

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
