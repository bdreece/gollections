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

import "fmt"

type ZipError struct {
	a error
	b error
}

func (z ZipError) Error() string {
	if z.a != nil && z.b != nil {
		return fmt.Sprintf("iter a: (%s), iter b: (%s)", z.a.Error(), z.b.Error())
	} else if z.a != nil {
		return fmt.Sprintf("iter a: (%s)", z.a.Error())
	} else if z.b != nil {
		return fmt.Sprintf("iter b: (%s)", z.b.Error())
	} else {
		return "unreachable! bug!"
	}
}

type ZipItem[T, U any] struct {
	A *T
	B *U
}

type Zip[T, U any] struct {
	a Iterator[T]
	b Iterator[U]
}

func NewZip[T any, U any](a Iterator[T], b Iterator[U]) *Zip[T, U] {
	return &Zip[T, U]{a, b}
}

func (z *Zip[T, U]) Next() (ZipItem[T, U], error) {
	a, err_a := z.a.Next()
	b, err_b := z.b.Next()

	item := ZipItem[T, U]{a, b}

	if err_a != nil || err_b != nil {
		return item, ZipError{
			a: err_a,
			b: err_b,
		}
	} else {
		return item, nil
	}
}
