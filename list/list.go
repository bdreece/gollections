// MIT License
// Copyright (c) 2022 Brian Reece

package list

import "github.com/bdreece/gollections/errors"

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

func newNode[T any](value T, next, prev *node[T]) *node[T] {
	return &node[T]{value, next, prev}
}

// List is the doubly-linked list data structure.
type List[T any] struct {
	head   *node[T]
	length int
}

// New constructs a new List
func New[T any]() *List[T] {
	return &List[T]{head: nil, length: 0}
}

// Front returns a pointer to the front
// of the List. Returns nil, errors.Empty
// if the List is empty.
func (l *List[T]) Front() (*T, error) {
	if l.head == nil {
		return nil, errors.Empty{}
	}

	return &l.head.value, nil
}

// Back returns a pointer to the back
// of the List. Returns nil, errors.Empty
// if the List is empty.
func (l *List[T]) Back() (*T, error) {
	if l.head == nil {
		return nil, errors.Empty{}
	}

	walk := l.head
	for walk.next != nil {
		walk = walk.next
	}

	return &walk.value, nil
}

// PushBack pushes an item onto the back
// of the doubly-linked list.
func (l *List[T]) PushBack(item T) {
	l.length += 1
	if l.head == nil {
		l.head = newNode(item, nil, nil)
		return
	}

	walk := l.head
	for walk.next != nil {
		walk = walk.next
	}

	walk.next = newNode(item, nil, walk)
}

// PushFront pushes an item onto the front
// of the double-linked list.
func (l *List[T]) PushFront(item T) {
	l.head = newNode(item, l.head, nil)
	l.head.next.prev = l.head
	l.length += 1
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
	l.head = l.head.next
	l.length -= 1
	return value, nil
}

// PopBack removes and returns an item from
// the back of the List. Returns nil, errors.Empty
// if the List is empty.
func (l *List[T]) PopBack() (*T, error) {
	if l.head == nil {
		return nil, errors.Empty{}
	}

	walk := l.head
	for walk.next != nil {
		walk = walk.next
	}

	value := new(T)
	*value = walk.value
	walk.prev.next = nil
	walk.prev = nil
	l.length -= 1
	return value, nil
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
			return nil, errors.NewIndexOutOfBounds(index, i)
		}
		walk = walk.next
	}

	val := &walk.value
	walk.prev.next = nil
	walk.prev = nil
	if walk.next != nil {
		walk.next.prev = nil
		walk.next = nil
	}
	return val, nil
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
			return errors.NewIndexOutOfBounds(index, i)
		}
		walk = walk.next
	}

	walk.value = item
	return nil
}
