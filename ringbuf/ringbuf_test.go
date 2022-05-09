// MIT License
// Copyright (c) 2022 Brian Reece

package ringbuf

import "testing"

const (
	EXPECTED string = "expected %s: (%d), got: (%d)\n"
	ERROR    string = "experienced error: %s\n"
)

func setup() (*RingBuf[int], []int) {
	buf := New[int](5)
	numbers := []int{1, 2, 3, 4, 5}
	for _, a := range numbers {
		buf.Enqueue(a)
	}

	return buf, numbers
}

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

// TestClear asserts that the Clear function
// properly zeroes all items in the ring
// buffer.
func TestClear(t *testing.T) {
	ringbuf, _ := setup()
	ringbuf.Clear()
	val, err := ringbuf.Dequeue()
	for ; val != nil; val, err = ringbuf.Dequeue() {
		if err != nil {
			t.Errorf(ERROR, err.Error())
		}
		if *val != 0 {
			t.Errorf(EXPECTED, "val", 0, *val)
		}
	}
}

// TestPeek asserts that the Peek function
// properly reads an item from the RingBuf
// without advancing the head pointer.
func TestPeek(t *testing.T) {
	ringbuf, numbers := setup()

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
