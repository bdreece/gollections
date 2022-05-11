// MIT License
// Copyright (c) 2022 Brian Reece

package list

import "testing"

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
		if *val != numbers[i] {
			t.Errorf(EXPECTED, "val", numbers[i], *val)
		}
	}

}
