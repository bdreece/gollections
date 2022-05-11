// MIT License
// Copyright (c) 2022 Brian Reece

package iterator

import "github.com/bdreece/gollections/errors"

type MapFunc[T, U any] func(*T) U

// Map transforms an iterator over one type into an iterator
// over another type, converting each item via pred. Returns
// Iterator[U], error on collection error.
type Map[T, U any] struct {
	iter Iterator[T]
	pred MapFunc[T, U]
}

func NewMap[T, U any](iter Iterator[T], pred MapFunc[T, U]) *Map[T, U] {
	return &Map[T, U]{iter, pred}
}

func (m *Map[T, U]) Next() (*U, error) {
	item, err := m.iter.Next()
	if item == nil {
		return nil, errors.Empty{}
	}
	if err != nil {
		return nil, err
	}
	mapped := (m.pred)(item)
	return &mapped, nil
}
