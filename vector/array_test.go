// MIT License
// Copyright (c) 2022 Brian Reece

package vector

import "testing"

// TestGet asserts that the Get function
// properly returns a pointer to an item
// in a Vector.
func TestGet(t *testing.T) {
	vec, numbers := setup()
	m := len(numbers)

	for i, number := range numbers {
		n := len(*vec)

		if n != m {
			t.Errorf(EXPECTED, "len", m, n)
		}

		val, err := vec.Get(i)

		if err != nil {
			t.Errorf(ERROR, err.Error())
		}

		if *val != number {
			t.Errorf(EXPECTED, "val", number, *val)
		}
	}
}

// TestSet asserts that the Set function
// properly writes a new value to an item
// in a Vector
func TestSet(t *testing.T) {
	vec, numbers := setup()
	m := len(numbers)

	for i := range numbers {
		n := len(*vec)

		if n != m {
			t.Errorf(EXPECTED, "len", m, n)
		}

		vec.Set(i, 0)
		val, err := vec.Get(i)

		if err != nil {
			t.Errorf(ERROR, err.Error())
		}

		if *val != 0 {
			t.Errorf(EXPECTED, "val", 0, *val)
		}
	}
}

// TestIns asserts that the Ins function
// properly inserts an item into the Vector
// at the specified index.
func TestIns(t *testing.T) {
	vec, numbers := setup()

	for i := range numbers {
		n := len(*vec)
		m := len(numbers) + i

		if n != m {
			t.Errorf(EXPECTED, "len", m, n)
		}

		vec.Ins(i, 0)
		val, err := vec.Get(i)

		if err != nil {
			t.Errorf(ERROR, err.Error())
		}

		if *val != 0 {
			t.Errorf(EXPECTED, "val", 0, *val)
		}
	}
}

// TestDel asserts that the Del function
// properly removes and returns the element
// at the specified index.
func TestDel(t *testing.T) {
	vec, numbers := setup()

	for i, number := range numbers {
		n := len(*vec)
		m := len(numbers) - i

		if n != m {
			t.Errorf(EXPECTED, "len", m, n)
		}

		elem, err := vec.Del(0)

		if err != nil {
			t.Errorf(ERROR, err.Error())
		}

		if *elem != number {
			t.Errorf(EXPECTED, "val", number, *elem)
		}
	}
}
