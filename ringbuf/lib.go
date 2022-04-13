package ringbuf

import "errors"

type RingBuf[T any] struct {
  data   []T
  capacity   int
  length int
  head   int
  tail   int
}

func New[T any](capacity int) RingBuf[T] {
  return RingBuf[T]{
    data: make([]T, capacity),
    capacity: capacity,
    length: 0,
    head: 0,
    tail: 0,
  }
}

func (b *RingBuf[T]) Read() (any, error) {
  if b.length <= 0 {
    return nil, errors.New("read from empty ringbuf")
  }

  val := b.data[b.head]
  
  b.head = (b.head + 1) % b.capacity
  b.length -= 1
  return val, nil
}

func (b RingBuf[T]) Peek() (any, error) {
  if b.length <= 0 {
    return nil, errors.New("peek from empty ringbuf")
  }
  return b.data[b.head], nil
}

func (b *RingBuf[T]) Write(val T) {
  b.data[b.tail] = val
  b.tail = (b.tail + 1) % b.capacity
  b.length += 1
}

func (b *RingBuf[T]) Clear() {
  b.data = make([]T, b.capacity)
}
