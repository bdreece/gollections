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
		list.PushFront(number)
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

// TestFront asserts that the Front function
// properly returns a pointer to the front of
// the list.
func TestFront(t *testing.T) {
	list, _ := setup()
	val, err := list.Front()
	if err != nil {
		t.Errorf(ERROR, err.Error())
	}
	if *val != 5 {
		t.Errorf(EXPECTED, "val", 5, *val)
	}
}

// TestBack asserts that the Back function
// properly returns a pointer to the back
// of the list.
func TestBack(t *testing.T) {
	list, _ := setup()
	val, err := list.Back()
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
		val, err := list.Front()
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

// TestSet asserts that the Set function
// properly sets the item at the specified
// index to the provided value.
func TestSet(t *testing.T) {
	list, numbers := setup()
	for i := range numbers {
		list.Set(i, 0)
		val, err := list.Get(i)
		if err != nil {
			t.Errorf(ERROR, err.Error())
		}
		if *val != 0 {
			t.Errorf(EXPECTED, "val", 0, *val)
		}
	}
}

// TestGet asserts that the Get function
// properly returns a pointer to the item
// at the specified index.
func TestGet(t *testing.T) {
	list, numbers := setup()
	for i := range numbers {
		val, err := list.Get(i)
		if err != nil {
			t.Errorf(ERROR, err.Error())
		}
		if *val != numbers[4-i] {
			t.Errorf(EXPECTED, "val", numbers[4-i], *val)
		}
	}

}

// TestCollect asserts that the Collect
// function properly inserts all the items
// from an iterator into the list.
func TestCollect(t *testing.T) {

}

// TestAppend asserts that the Append
// function properly appends a variable
// number of items to the end of the list.
func TestAppend(t *testing.T) {

}

func TestExtend(t *testing.T) {

}
