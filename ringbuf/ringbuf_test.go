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

import "testing"

const (
	EXPECTED string = "expected %s: (%d), got: (%d)\n"
	ERROR    string = "experienced error: %s\n"
)

func checkFields(t *testing.T, b *RingBuf[int], length, capacity, head, tail int) {
	if b.length != length {
		t.Errorf(EXPECTED, "len", length, b.length)
	}

	if b.capacity != capacity {
		t.Errorf(EXPECTED, "cap", capacity, b.capacity)
	}

	if b.head != head {
		t.Errorf(EXPECTED, "head", head, b.head)
	}

	if b.tail != tail {
		t.Errorf(EXPECTED, "tail", tail, b.tail)
	}

}

func TestNew(t *testing.T) {
	ringbuf := New[int](5)
	checkFields(t, ringbuf, 0, 5, 0, 0)
}

func TestWrite(t *testing.T) {
	ringbuf := New[int](5)
	numbers := []int{1, 2, 3, 4, 5}

	for i, number := range numbers {
		checkFields(t, ringbuf, i, 5, 0, i)
		ringbuf.Write(number)
	}
}

func TestRead(t *testing.T) {
	ringbuf := New[int](5)
	numbers := []int{1, 2, 3, 4, 5}

	for _, number := range numbers {
		ringbuf.Write(number)
	}

	for i, number := range numbers {
		checkFields(t, ringbuf, 5-i, 5, i, 0)
		val, err := ringbuf.Read()
		if err != nil {
			t.Errorf(ERROR, err.Error())
		}
		if *val != number {
			t.Errorf(EXPECTED, "val", number, *val)
		}
	}
}

func TestPeek(t *testing.T) {
	ringbuf := New[int](5)
	numbers := []int{1, 2, 3, 4, 5}

	for _, number := range numbers {
		ringbuf.Write(number)
	}

	for range numbers {
		checkFields(t, ringbuf, 5, 5, 0, 0)
		val, err := ringbuf.Peek()
		if err != nil {
			t.Errorf(ERROR, err.Error())
		}

		if *val != 1 {
			t.Errorf(EXPECTED, "val", 1, *val)
		}
	}
}
