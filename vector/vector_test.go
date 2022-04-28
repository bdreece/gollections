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

import "testing"

const (
	EXPECTED string = "expected %s: (%d), got: (%d)\n"
	ERROR    string = "experienced error: %s\n"
)

func setup() (*Vector[int], []int) {
	vec := New[int]()
	numbers := []int{1, 2, 3, 4, 5}
	for _, number := range numbers {
		vec.PushBack(number)
	}
	return vec, numbers
}

func TestNew(t *testing.T) {
	vec := New[int]()

	if len(*vec) != 0 {
		t.Errorf(EXPECTED, "len", 0, len(*vec))
	}
}

func TestPushBack(t *testing.T) {
	vec := New[int]()
	numbers := []int{1, 2, 3, 4, 5}

	for i, number := range numbers {
		n := len(*vec)

		if n != i {
			t.Errorf(EXPECTED, "len", i, n)
		}

		vec.PushBack(number)
	}
}

func TestPushFront(t *testing.T) {
	vec := New[int]()
	numbers := []int{1, 2, 3, 4, 5}
	for i, number := range numbers {
		n := len(*vec)

		if n != i {
			t.Errorf(EXPECTED, "len", i, n)
		}

		vec.PushFront(number)
	}
}

func TestPopFront(t *testing.T) {
	vec, numbers := setup()

	for i, number := range numbers {
		n := len(*vec)
		m := len(numbers) - i

		if n != m {
			t.Errorf(EXPECTED, "len", m, n)
		}

		val, err := vec.PopFront()

		if err != nil {
			t.Errorf(ERROR, err.Error())
		}

		if *val != number {
			t.Errorf(EXPECTED, "val", number, *val)
		}
	}
}

func TestPopBack(t *testing.T) {
	vec, numbers := setup()

	for i, number := range numbers {
		n := len(*vec)
		m := len(numbers) - i

		if n != m {
			t.Errorf(EXPECTED, "len", m, n)
		}

		val, err := vec.PopBack()

		if err != nil {
			t.Errorf(ERROR, err.Error())
		}

		if *val != 6-number {
			t.Errorf(EXPECTED, "val", 6-number, *val)
		}
	}
}

func TestGet(t *testing.T) {
	vec, numbers := setup()
	m := len(numbers)

	for i, number := range numbers {
		n := len(*vec)

		if n != m {
			t.Errorf(EXPECTED, "len", m, n)
		}

		val, err := vec.Get(i)

		if err != nil {
			t.Errorf(ERROR, err.Error())
		}

		if *val != number {
			t.Errorf(EXPECTED, "val", number, *val)
		}
	}
}

func TestSet(t *testing.T) {
	vec, numbers := setup()
	m := len(numbers)

	for i := range numbers {
		n := len(*vec)

		if n != m {
			t.Errorf(EXPECTED, "len", m, n)
		}

		vec.Set(i, 0)
		val, err := vec.Get(i)

		if err != nil {
			t.Errorf(ERROR, err.Error())
		}

		if *val != 0 {
			t.Errorf(EXPECTED, "val", 0, *val)
		}
	}
}

func TestInsertBefore(t *testing.T) {
	vec, numbers := setup()

	for i := range numbers {
		n := len(*vec)
		m := len(numbers) + i

		if n != m {
			t.Errorf(EXPECTED, "len", m, n)
		}

		vec.InsertBefore(i, 0)
		val, err := vec.Get(i)

		if err != nil {
			t.Errorf(ERROR, err.Error())
		}

		if *val != 0 {
			t.Errorf(EXPECTED, "val", 0, *val)
		}
	}
}

func TestInsertAfter(t *testing.T) {
	vec, numbers := setup()

	for i := range numbers {
		n := len(*vec)
		m := len(numbers) + i

		if n != m {
			t.Errorf(EXPECTED, "len", m, n)
		}

		vec.InsertAfter(i, 0)
		val, err := vec.Get(i + 1)

		if err != nil {
			t.Errorf(ERROR, err.Error())
		}

		if *val != 0 {
			t.Errorf(EXPECTED, "val", 0, *val)
		}
	}
}
