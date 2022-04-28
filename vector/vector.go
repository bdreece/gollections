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

package vector

import "fmt"

type IndexOutOfBoundsError struct {
	index  int
	bounds int
}

func (e *IndexOutOfBoundsError) Error() string {
	return fmt.Sprintf("index %d > bounds %d", e.index, e.bounds)
}

type Vector[T any] []T

func New[T any]() *Vector[T] {
	return new(Vector[T])
}

func (v *Vector[T]) PushBack(value T) {
	*v = append([]T(*v), value)
}

func (v *Vector[T]) PushFront(value T) {
	newVec := []T{value}
	*v = append(newVec, []T(*v)...)
}

func (v *Vector[T]) PopFront() (*T, error) {
	n := len(*v)
	if n == 0 {
		return nil, &IndexOutOfBoundsError{
			index:  0,
			bounds: 0,
		}
	}

	val := new(T)
	*val = []T(*v)[0]
	if n > 0 {
		*v = []T(*v)[1:]
	} else {
		v = New[T]()
	}
	return val, nil
}

func (v *Vector[T]) PopBack() (*T, error) {
	n := len(*v)
	if n == 0 {
		return nil, &IndexOutOfBoundsError{
			index:  0,
			bounds: 0,
		}
	}
	val := new(T)
	if n > 0 {
		*val = []T(*v)[n-1]
		*v = []T(*v)[:n-1]
	} else {
		*val = []T(*v)[0]
		v = New[T]()
	}
	return val, nil
}

func (v Vector[T]) Get(i int) (*T, error) {
	n := len(v)
	if i > n {
		return nil, &IndexOutOfBoundsError{
			index:  i,
			bounds: n,
		}
	}
	val := new(T)
	*val = []T(v)[i]
	return val, nil
}

func (v *Vector[T]) Set(i int, value T) error {
	n := len(*v)
	if i >= n {
		return &IndexOutOfBoundsError{
			index:  i,
			bounds: n,
		}
	}
	[]T(*v)[i] = value
	return nil
}

func (v *Vector[T]) InsertAfter(i int, value T) error {
	n := len(*v)
	if i >= n {
		return &IndexOutOfBoundsError{
			index:  i,
			bounds: n,
		}
	}
	before := []T(*v)[:i+1]
	after := []T(*v)[i+1:]
	*v = append(before, value)
	*v = append(*v, after...)
	return nil
}

func (v *Vector[T]) InsertBefore(i int, value T) error {
	var (
		before []T
		after  []T
	)
	n := len(*v)
	if i >= n {
		return &IndexOutOfBoundsError{
			index:  i,
			bounds: n,
		}
	}
	if i > 0 {
		before = []T(*v)[:i]
		after = []T(*v)[i:]
	} else {
		before = []T{}
		after = []T(*v)
	}
	*v = append(before, value)
	*v = append(*v, after...)
	return nil
}

func (v *Vector[T]) Fill(values ...T) {
	for _, value := range values {
		v.PushBack(value)
	}
}

func (v *Vector[T]) Iterator() Iterator[T] {
	return Iterator[T]{v}
}
