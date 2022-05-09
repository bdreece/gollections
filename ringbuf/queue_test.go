// MIT License
// Copyright (c) 2022 Brian Reece

package ringbuf

import "testing"

// TestEnqueue asserts that the Enqueue function
// properly writes an item into the RingBuf.
func TestEnqueue(t *testing.T) {
	ringbuf := New[int](5)
	numbers := []int{1, 2, 3, 4, 5}

	for i, number := range numbers {
		checkFields(t, ringbuf, i, 5, 0, i)
		ringbuf.Enqueue(number)
	}
}

// TestDequeue asserts that the Read function
// properly reads an item from the RingBuf.
func TestDequeue(t *testing.T) {
	ringbuf, numbers := setup()

	for i, number := range numbers {
		checkFields(t, ringbuf, 5-i, 5, i, 0)
		val, err := ringbuf.Dequeue()
		if err != nil {
			t.Errorf(ERROR, err.Error())
		}
		if *val != number {
			t.Errorf(EXPECTED, "val", number, *val)
		}
	}
}
