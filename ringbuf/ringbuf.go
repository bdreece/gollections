// MIT License
// Copyright (c) 2022 Brian Reece

package ringbuf

import "github.com/bdreece/gollections/errors"

type RingBuf[T any] struct {
	data     []T
	capacity int
	length   int
	head     int
	tail     int
}

func New[T any](capacity int) *RingBuf[T] {
	return &RingBuf[T]{
		data:     make([]T, capacity),
		capacity: capacity,
		length:   0,
		head:     0,
		tail:     0,
	}
}

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

func (b RingBuf[T]) Peek() (*T, error) {
	if b.length <= 0 {
		return nil, errors.Empty{}
	}
	val := new(T)
	*val = b.data[b.head]
	return val, nil
}

func (b *RingBuf[T]) Write(val T) {
	b.data[b.tail] = val
	b.tail = (b.tail + 1) % b.capacity
	b.length += 1
}

func (b *RingBuf[T]) Clear() {
	b.data = make([]T, b.capacity)
}

func (b *RingBuf[T]) Collect(values ...T) {
	for _, value := range values {
		b.Write(value)
	}
}

func (b *RingBuf[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{b}
}
