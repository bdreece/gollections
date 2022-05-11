// MIT License
// Copyright (c) 2022 Brian Reece

package list

import "testing"

const (
	EXPECTED string = "expected %s: (%d), got: (%d)\n"
	ERROR    string = "experienced error: \"%s\"\n"
)

func setup() (*List[int], []int) {
	list := New[int]()
	numbers := []int{1, 2, 3, 4, 5}
	for _, number := range numbers {
		list.PushBack(number)
	}
	return list, numbers
}

// TestNew asserts that the New function
// properly constructs a new List.
func TestNew(t *testing.T) {
	list := New[int]()
	if list.length != 0 {
		t.Errorf(EXPECTED, "len", 0, list.length)
	}
}
