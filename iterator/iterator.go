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

type Iterator[T any] interface {
	Next() (*T, error)
}

func ForEach[T any](iter Iterator[T], fn func(*T)) error {
	for {
		item, err := iter.Next()
		if item == nil {
			break
		}
		if err != nil {
			return err
		}
		(fn)(item)
	}

	return nil
}

func Reduce[T any](iter Iterator[T], fn func(*T, T) *T) (*T, error) {
	total := new(T)
	for {
		item, err := iter.Next()
		if err != nil {
			return total, err
		}
		if item == nil {
			break
		}
		(fn)(total, *item)
	}
	return total, nil
}

func Fold[T any, U any](iter Iterator[T], init *U, fn func(*U, T) *U) (*U, error) {
	for {
		item, err := iter.Next()
		if err != nil {
			return init, err
		}
		if item == nil {
			break
		}
		(fn)(init, *item)
	}
	return init, nil
}
