// MIT License
// Copyright (c) 2022 Brian Reece

package list

import "testing"

// TestPeekFront asserts that the Front function
// properly returns a pointer to the front of
// the list.
func TestPeekFront(t *testing.T) {
	list, _ := setup()
	val, err := list.PeekFront()
	if err != nil {
		t.Errorf(ERROR, err.Error())
	}
	if *val != 5 {
		t.Errorf(EXPECTED, "val", 5, *val)
	}
}

// TestPeekBack asserts that the Back function
// properly returns a pointer to the back
// of the list.
func TestPeekBack(t *testing.T) {
	list, _ := setup()
	val, err := list.PeekBack()
	if err != nil {
		t.Errorf(ERROR, err.Error())
	}
	if *val != 1 {
		t.Errorf(EXPECTED, "val", 1, *val)
	}
}

// TestPushFront asserts that the PushFront
// function properly prepends the List
// with the specified item.
func TestPushFront(t *testing.T) {
	list := New[int]()
	numbers := []int{1, 2, 3, 4, 5}
	for _, number := range numbers {
		list.PushFront(number)
		val, err := list.PeekFront()
		if err != nil {
			t.Errorf(ERROR, err.Error())
		}
		if *val != number {
			t.Errorf(EXPECTED, "val", number, *val)
		}
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

func TestPopBack(t *testing.T) {
	list, numbers := setup()
	for i, number := range numbers {
		if list.length != 5-i {
			t.Errorf(EXPECTED, "len", 5-i, list.length)
		}
		val, err := list.PopBack()
		if err != nil {
			t.Errorf(ERROR, err.Error())
		}
		if *val != number {
			t.Errorf(EXPECTED, "val", number, *val)
		}
	}
}
