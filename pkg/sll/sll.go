package sll

import (
	"github.com/bdreece/gollections/pkg/collection"
	"github.com/bdreece/gollections/pkg/iterator"
	"github.com/bdreece/gollections/pkg/queue"
	"github.com/bdreece/gollections/pkg/stack"
)

// SLL provides the implementation of a singly-linked list
type SLL[TItem any] interface {
	stack.Stack[TItem]
	queue.Queue[TItem]
}

type node[TItem any] struct {
	item TItem
	next *node[TItem]
}

type sll[TItem any] struct {
	first *node[TItem]
}

// New creates a new SLL
func New[TItem any]() SLL[TItem] {
	return &sll[TItem]{nil}
}

// From creates a new SLL by concatenating the items
// of the given collection
func From[TItem any](c collection.Collection[TItem]) SLL[TItem] {
	s := New[TItem]()
	s.Concat(c)
	return s
}

func (s *sll[TItem]) Count() int {
	i := 0
	current := s.first
	for current != nil {
		current = current.next
		i++
	}

	return i
}

func (s *sll[TItem]) Concat(into iterator.IntoIterator[TItem]) collection.Collection[TItem] {
	return s.Collect(into.Iter())
}

func (s *sll[TItem]) Collect(iter iterator.Iterator[TItem]) collection.Collection[TItem] {
	iterator.ForEach(iter, func(item TItem) {
		s.Append(item)
	})
	return s
}

func (s *sll[TItem]) Append(item TItem) collection.Collection[TItem] {
	s.insertLast(item)
	return s
}

func (s *sll[TItem]) Push(item TItem) {
	s.insertFirst(item)
}

func (s *sll[TItem]) Pop() *TItem {
	return s.removeFirst()
}

func (s *sll[TItem]) Peek() *TItem {
	if s.first == nil {
		return nil
	}

	return &s.first.item
}

func (s *sll[TItem]) Enqueue(item TItem) {
	s.insertLast(item)
}

func (s *sll[TItem]) Dequeue() *TItem {
	return s.removeFirst()
}

func (s *sll[TItem]) insertFirst(item TItem) {
	tmp := s.first
	s.first = &node[TItem]{item, tmp}
}

func (s *sll[TItem]) insertLast(item TItem) {
	current := s.first
	if current == nil {
		s.first = &node[TItem]{item, nil}
		return
	}

	for current.next != nil {
		current = current.next
	}

	current.next = &node[TItem]{item, nil}
}

func (s *sll[TItem]) removeFirst() *TItem {
	tmp := s.first
	if tmp == nil {
		return nil
	}

	s.first = s.first.next
	return &tmp.item

}
