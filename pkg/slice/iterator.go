package slice

import (
	"github.com/bdreece/gollections/pkg/iterator"
)

type sliceIterator[TItem any] struct {
	_slice *slice[TItem]
	index  int
}

func (s *slice[TItem]) Iter() iterator.Iterator[TItem] {
	return &sliceIterator[TItem]{s, 0}
}

func (s *sliceIterator[TItem]) Next() *TItem {
	if s.index >= len(*s._slice) {
		return nil
	}
	val := (*s._slice)[s.index]
	s.index++
	return &val
}
