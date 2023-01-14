package sll

import "github.com/bdreece/gollections/pkg/iterator"

type sllIterator[TItem any] struct {
	current *node[TItem]
}

func (s *sll[TItem]) Iter() iterator.Iterator[TItem] {
	return &sllIterator[TItem]{s.first}
}

func (s *sllIterator[TItem]) Next() *TItem {
	tmp := s.current
	if tmp == nil {
		return nil
	}
	s.current = s.current.next
	return &tmp.item
}
