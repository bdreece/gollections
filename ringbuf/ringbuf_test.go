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

// TestNew asserts that the New function
// properly constructs a RingBuf with the
// specified capacity.
func TestNew(t *testing.T) {
	ringbuf := New[int](5)
	checkFields(t, ringbuf, 0, 5, 0, 0)
}

// TestWrite asserts that the Write function
// properly writes an item into the RingBuf.
func TestWrite(t *testing.T) {
	ringbuf := New[int](5)
	numbers := []int{1, 2, 3, 4, 5}

	for i, number := range numbers {
		checkFields(t, ringbuf, i, 5, 0, i)
		ringbuf.Write(number)
	}
}

// TestRead asserts that the Read function
// properly reads an item from the RingBuf.
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

// TestPeek asserts that the Peek function
// properly reads an item from the RingBuf
// without advancing the head pointer.
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
