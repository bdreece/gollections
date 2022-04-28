// MIT License
// Copyright (c) 2022 Brian Reece

package ringbuf

type Iterator[T any] struct {
	*RingBuf[T]
}

func (iter *Iterator[T]) Next() (*T, error) {
	return iter.RingBuf.Read()
}
