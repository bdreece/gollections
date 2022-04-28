// MIT License
// Copyright (c) 2022 Brian Reece

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
