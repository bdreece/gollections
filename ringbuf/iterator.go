// MIT License
// Copyright (c) 2022 Brian Reece

package ringbuf

// Iterator provides an iterator over the
// items in a RingBuf
type Iterator[T any] struct {
	*RingBuf[T]
}

// Next returns the next item in the RingBuf.
// Returns nil, errors.Empty after the last
// item has been read.
func (iter *Iterator[T]) Next() (*T, error) {
	return iter.RingBuf.Read()
}
