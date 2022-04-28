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

package iterator

type EnumerateItem[T any] struct {
	item  T
	index int
}

type Enumerate[T any] struct {
	iter  Iterator[T]
	index int
}

func NewEnumerate[T any](iter Iterator[T]) Enumerate[T] {
	return Enumerate[T]{
		iter:  iter,
		index: 0,
	}
}

func (e *Enumerate[T]) Next() (*EnumerateItem[T], error) {
	item, err := e.iter.Next()
	enum_item := EnumerateItem[T]{
		item:  *item,
		index: e.index,
	}
	e.index += 1
	if err != nil {
		return &enum_item, err
	}
	return &enum_item, nil
}
