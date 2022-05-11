// MIT License
// Copyright (c) 2022 Brian Reece

package list

import (
	"github.com/bdreece/gollections/errors"
	"github.com/bdreece/gollections/iterator"
)

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

func newNode[T any](value T, prev, next *node[T]) *node[T] {
	return &node[T]{value, next, prev}
}

// List is the doubly-linked list data structure.
type List[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

// New constructs a new List
func New[T any]() *List[T] {
	return &List[T]{head: nil, tail: nil, length: 0}
}

// PushBack pushes an item onto the back
// of the doubly-linked list.
func (l *List[T]) PushBack(item T) {
	l.length += 1
	if l.head == nil {
		l.head = newNode(item, nil, nil)
		l.tail = l.head
		return
	}

	l.tail.next = newNode(item, l.tail, nil)
	l.tail = l.tail.next
}

// PopBack removes and returns an item from
// the back of the List. Returns nil, errors.Empty
// if the List is empty.
func (l *List[T]) PopBack() (*T, error) {
	if l.tail == nil {
		return nil, errors.Empty{}
	}
	value := new(T)
	*value = l.tail.value

	if l.tail.prev != nil {
		// Recede tail
		l.tail = l.tail.prev

		// Unlink removed tail
		l.tail.next = nil
	} else {
		// Truncate list
		l.head = nil
		l.tail = nil
	}

	l.length -= 1
	return value, nil
}

// PeekBack returns a pointer to the back
// of the List. Returns nil, errors.Empty
// if the List is empty.
func (l *List[T]) PeekBack() (*T, error) {
	if l.tail == nil {
		return nil, errors.Empty{}
	}

	return &l.tail.value, nil
}

// PushFront pushes an item onto the front
// of the double-linked list.
func (l *List[T]) PushFront(item T) {
	l.length += 1
	if l.head == nil {
		l.head = newNode(item, nil, nil)
		l.tail = l.head
		return
	}
	l.head.prev = newNode(item, nil, l.head)
	l.head = l.head.prev
}

// PopFront removes and returns an item from
// the front of the List. Returns nil, errors.Empty
// if the List is empty.
func (l *List[T]) PopFront() (*T, error) {
	if l.head == nil {
		return nil, errors.Empty{}
	}

	value := new(T)
	*value = l.head.value

	if l.head.next != nil {
		// Advance head
		l.head = l.head.next
		// Unlink deleted head
		l.head.prev = nil
	} else {
		// Truncate list
		l.head = nil
		l.tail = nil
	}

	l.length -= 1
	return value, nil
}

// PeekFront returns a pointer to the front
// of the List. Returns nil, errors.Empty
// if the List is empty.
func (l *List[T]) PeekFront() (*T, error) {
	if l.head == nil {
		return nil, errors.Empty{}
	}

	return &l.head.value, nil
}

// Enqueue appends an item to the end of
// the list.
func (l *List[T]) Enqueue(elem T) {
	l.PushBack(elem)
}

// Dequeue removes and returns the first
// element in the list.
func (l *List[T]) Dequeue() (*T, error) {
	return l.PopFront()
}

// Push prepends the list with an element.
func (l *List[T]) Push(elem T) {
	l.PushFront(elem)
}

// Pop removes and returns the first
// element in the list.
func (l *List[T]) Pop() (*T, error) {
	return l.PopFront()
}

// Peek returns the first element in
// the list without removing it.
func (l *List[T]) Peek() (*T, error) {
	return l.PeekFront()
}

// Get returns a pointer to an item in the
// List located at index. Returns nil, errors.Empty
// if the List is empty. Returns nil,
// errors.IndexOutOfBounds if index < 0 or index > len(l).
func (l *List[T]) Get(index int) (*T, error) {
	if l.head == nil {
		return nil, errors.Empty{}
	}

	walk := l.head
	for i := 0; i < index; i++ {
		if walk.next == nil {
			return nil, errors.IndexOutOfBounds{Index: index, Bounds: i}
		}
		walk = walk.next
	}

	return &walk.value, nil
}

// Set sets the item located at index in
// the List. returns errors.IndexOutOfBounds
// if index < 0 or index > len(l). Returns
// errors.Empty if index > 0 and List is empty.
func (l *List[T]) Set(index int, item T) error {
	if l.head == nil {
		if index == 0 {
			l.head = newNode(item, nil, nil)
			return nil
		}
		return errors.Empty{}
	}

	walk := l.head
	for i := 0; i < index; i++ {
		if walk.next == nil {
			return errors.IndexOutOfBounds{Index: index, Bounds: i}
		}
		walk = walk.next
	}

	walk.value = item
	return nil
}

// Ins inserts an element into the list at the specified
// index. Returns errors.IndexOutOfBounds if index < 0
// or index > len(l).
func (l *List[T]) Ins(index int, elem T) error {
	if l.head == nil {
		return errors.Empty{}
	}

	walk := l.head
	for i := 0; i < index; i++ {
		if walk.next == nil {
			return errors.IndexOutOfBounds{Index: index, Bounds: i}
		}
		walk = walk.next
	}

	item := newNode[T](elem, walk.prev, walk)
	item.prev.next = item
	walk.prev = item
	return nil
}

// Del removes and returns an element from the list at the
// specified index. Returns errors.IndexOutOfBounds if
// index < 0 or index > len(l).
func (l *List[T]) Del(index int) (*T, error) {
	if l.head == nil {
		return nil, errors.Empty{}
	}

	walk := l.head
	for i := 0; i < index; i++ {
		if walk.next == nil {
			return nil, errors.IndexOutOfBounds{Index: index, Bounds: i}
		}
		walk = walk.next
	}
	value := new(T)
	*value = walk.value

	// Unlink neighbors
	walk.prev.next = nil
	walk.next.prev = nil

	return value, nil
}

// Collect appends a variable number of items to the list.
func (l *List[T]) Collect(items ...T) {
	for _, item := range items {
		l.PushBack(item)
	}
}

// IntoIterator returns an iterator over the items in the list.
func (l *List[T]) IntoIterator() iterator.Iterator[T] {
	return &Iterator[T]{l.head}
}

// FromIterator collects the items from the iterator into the
// list.
func (l *List[T]) FromIterator(iter iterator.Iterator[T]) error {
	if err := iterator.ForEach(iter, func(item *T) {
		l.PushBack(*item)
	}); err != nil {
		return err
	}
	return nil
}
