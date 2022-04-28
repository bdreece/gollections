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

type List[T any] struct {
	head   *node[T]
	length int
}

func New[T any]() *List[T] {
	return &List[T]{head: nil, length: 0}
}

func (l *List[T]) Front() (*T, error) {
	if l.head == nil {
		return nil, errors.Empty{}
	}

	val := new(T)
	*val = l.head.value
	return val, nil
}

func (l *List[T]) Back() (*T, error) {
	if l.head == nil {
		return nil, errors.Empty{}
	}

	walk := l.head
	for walk.next != nil {
		walk = walk.next
	}

	val := new(T)
	*val = walk.value
	return val, nil
}

func (l *List[T]) PushBack(val T) {
	l.length += 1
	if l.head == nil {
		l.head = newNode(val, nil, nil)
		return
	}

	walk := l.head
	for walk.next != nil {
		walk = walk.next
	}

	walk.next = newNode(val, nil, walk)
}

func (l *List[T]) PushFront(val T) {
	l.head = newNode(val, l.head, nil)
	l.head.next.prev = l.head
	l.length += 1
}

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

	val := new(T)
	*val = walk.value
	walk.prev.next = nil
	walk.prev = nil
	if walk.next != nil {
		walk.next.prev = nil
		walk.next = nil
	}
	return val, nil
}
