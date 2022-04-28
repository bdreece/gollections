// MIT License
// Copyright (c) 2022 Brian Reece

package ringbuf

import "github.com/bdreece/gollections/errors"

// RingBuf is the ring buffer data structure
type RingBuf[T any] struct {
	data     []T
	capacity int
	length   int
	head     int
	tail     int
}

// New constructs a new RingBuf with capacity
func New[T any](capacity int) *RingBuf[T] {
	return &RingBuf[T]{
		data:     make([]T, capacity),
		capacity: capacity,
		length:   0,
		head:     0,
		tail:     0,
	}
}

// Read reads an item from the RingBuf,
// advancing the head pointer. Returns
// nil, errors.Empty if ring buffer is
// empty.
func (b *RingBuf[T]) Read() (*T, error) {
	if b.length <= 0 {
		return nil, errors.Empty{}
	}

	val := new(T)
	*val = b.data[b.head]

	b.head = (b.head + 1) % b.capacity
	b.length -= 1
	return val, nil
}

// Peek reads an item from the RingBuf
// without advancing the head pointer,
// allowing the value to be read again.
// Returns nil, errors.Empty if ring
// buffer is empty.
func (b RingBuf[T]) Peek() (*T, error) {
	if b.length <= 0 {
		return nil, errors.Empty{}
	}
	val := new(T)
	*val = b.data[b.head]
	return val, nil
}

// Write writes an item into the RingBuf,
// advancing the tail pointer.
func (b *RingBuf[T]) Write(item T) {
	b.data[b.tail] = item
	b.tail = (b.tail + 1) % b.capacity
	b.length += 1
}

// Clear reconstructs the RingBuf in place,
// effectively zeroing all the items.
func (b *RingBuf[T]) Clear() {
	b.data = make([]T, b.capacity)
}

// Collect writes a variable number of items
// into the RingBuf. This method implements
// part of the Iterator interface.
func (b *RingBuf[T]) Collect(values ...T) {
	for _, value := range values {
		b.Write(value)
	}
}

// Iterator returns an iterator over the items
// in the RingBuf. This method implements part
// of the Iterator interface.
func (b *RingBuf[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{b}
}
