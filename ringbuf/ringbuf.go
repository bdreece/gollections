/**
 * MIT License
 *
 * Copyright (c) 2022 Brian Reece
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package ringbuf

import "errors"

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
		return nil, errors.New("read from empty ringbuf")
	}

	val := new(T)
	*val = b.data[b.head]

	b.head = (b.head + 1) % b.capacity
	b.length -= 1
	return val, nil
}

func (b RingBuf[T]) Peek() (*T, error) {
	if b.length <= 0 {
		return nil, errors.New("peek from empty ringbuf")
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
