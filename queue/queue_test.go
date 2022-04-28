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

package queue

import "testing"

const (
	EXPECTED_LEN string = "expected len: (%d), got: (%d)\n"
	EXPECTED_VAL string = "expected val: (%d), got: (%d)\n"
)

func TestNewQueue(t *testing.T) {
	queue := New[int]()
	if len(*queue) != 0 {
		t.Errorf(EXPECTED_LEN, 0, len(*queue))
	}
}

func TestEnqueue(t *testing.T) {
	queue := New[int]()
	numbers := []int{1, 2, 3, 4, 5}
	for i, number := range numbers {
		if len(*queue) != i {
			t.Errorf(EXPECTED_LEN, i, len(*queue))
		}
		queue.Enqueue(number)
	}
	if len(*queue) != 5 {
		t.Errorf(EXPECTED_LEN, 5, len(*queue))
	}
}

func TestDequeue(t *testing.T) {
	queue := New[int]()
	numbers := []int{1, 2, 3, 4, 5}
	for i, number := range numbers {
		if len(*queue) != i {
			t.Errorf(EXPECTED_LEN, i, len(*queue))
		}
		queue.Enqueue(number)
	}

	for _, number := range numbers {
		val := queue.Dequeue()
		if *val != number {
			t.Errorf(EXPECTED_VAL, number, *val)
		}
	}
}

func TestPeek(t *testing.T) {
	queue := New[int]()
	numbers := []int{1, 2, 3, 4, 5}

	for i, number := range numbers {
		if i != len(*queue) {
			t.Errorf(EXPECTED_LEN, i, len(*queue))
		}
		queue.Enqueue(number)
	}

	for range numbers {
		val := queue.Peek()
		if *val != 1 {
			t.Errorf(EXPECTED_VAL, 1, *val)
		}
	}
}
