// MIT License
// Copyright (c) 2022 Brian Reece

package list

import "testing"

const (
	EXPECTED string = "expected %s: (%d), got: (%d)\n"
	ERROR    string = "experienced error: \"%s\"\n"
)

// TestNew asserts that the New function
// properly constructs a new List.
func TestNew(t *testing.T) {
	list := New[int]()
	if list.length != 0 {
		t.Errorf(EXPECTED, "len", 0, list.length)
	}
}

// TestPushBack asserts that the PushBack
// function properly appends the List with
// an item.
func TestPushBack(t *testing.T) {
	list := New[int]()
	numbers := []int{1, 2, 3, 4, 5}

	for i, number := range numbers {
		list.PushBack(number)
		if list.length != i+1 {
			t.Errorf(EXPECTED, "len", i+1, list.length)
		}
	}
}

// TestPopFront asserts that the PopFront
// function properly removes and returns
// an item from the front of the List.
func TestPopFront(t *testing.T) {
	list := New[int]()
	numbers := []int{1, 2, 3, 4, 5}

	for _, number := range numbers {
		list.PushBack(number)
	}

	for i, number := range numbers {
		if list.length != 5-i {
			t.Errorf(EXPECTED, "len", 5-i, list.length)
		}
		val, err := list.PopFront()
		if err != nil {
			t.Errorf(ERROR, err.Error())
		}
		if *val != number {
			t.Errorf(EXPECTED, "val", number, *val)
		}
	}
}
