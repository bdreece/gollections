// MIT License
// Copyright (c) 2022 Brian Reece

package vector

import "testing"

const (
	EXPECTED string = "expected %s: (%d), got: (%d)\n"
	ERROR    string = "experienced error: %s\n"
)

func setup() (*Vector[int], []int) {
	vec := New[int]()
	numbers := []int{1, 2, 3, 4, 5}
	for _, number := range numbers {
		vec.PushBack(number)
	}
	return vec, numbers
}

// TestNew asserts that the New function
// properly constructs an empty Vector.
func TestNew(t *testing.T) {
	vec := New[int]()

	if len(*vec) != 0 {
		t.Errorf(EXPECTED, "len", 0, len(*vec))
	}
}

// TestFront asserts that the Front function
// properly returns the first element in the
// vector.
func TestFront(t *testing.T) {
	vec, numbers := setup()
	val, err := vec.Front()
	if err != nil {
		t.Errorf(ERROR, err.Error())
	}
	if *val != numbers[0] {
		t.Errorf(EXPECTED, "val", numbers[0], *val)
	}
}

// TestBack asserts that the Back function
// properly returns the last element in the
// vector.
func TestBack(t *testing.T) {
	vec, numbers := setup()
	val, err := vec.Back()
	if err != nil {
		t.Errorf(ERROR, err.Error())
	}
	if *val != numbers[len(numbers)-1] {
		t.Errorf(EXPECTED, "val", numbers[len(numbers)-1], *val)
	}
}

// TestPushBack asserts that the PushBack
// function properly appends a Vector with
// items.
func TestPushBack(t *testing.T) {
	vec := New[int]()
	numbers := []int{1, 2, 3, 4, 5}

	for i, number := range numbers {
		n := len(*vec)

		if n != i {
			t.Errorf(EXPECTED, "len", i, n)
		}

		vec.PushBack(number)
	}
}

// TestPushFront asserts that the PushFront
// function properly prepends a Vector with
// items.
func TestPushFront(t *testing.T) {
	vec := New[int]()
	numbers := []int{1, 2, 3, 4, 5}
	for i, number := range numbers {
		n := len(*vec)

		if n != i {
			t.Errorf(EXPECTED, "len", i, n)
		}

		vec.PushFront(number)
	}
}

// TestPopFront asserts that the PopFront
// function properly removes and returns
// the first item in the Vector
func TestPopFront(t *testing.T) {
	vec, numbers := setup()

	for i, number := range numbers {
		n := len(*vec)
		m := len(numbers) - i

		if n != m {
			t.Errorf(EXPECTED, "len", m, n)
		}

		val, err := vec.PopFront()

		if err != nil {
			t.Errorf(ERROR, err.Error())
		}

		if *val != number {
			t.Errorf(EXPECTED, "val", number, *val)
		}
	}
}

// TestPopBack asserts that the PopBack
// function properly removes and returns
// the last item in the Vector.
func TestPopBack(t *testing.T) {
	vec, numbers := setup()

	for i, number := range numbers {
		n := len(*vec)
		m := len(numbers) - i

		if n != m {
			t.Errorf(EXPECTED, "len", m, n)
		}

		val, err := vec.PopBack()

		if err != nil {
			t.Errorf(ERROR, err.Error())
		}

		if *val != 6-number {
			t.Errorf(EXPECTED, "val", 6-number, *val)
		}
	}
}

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

// TestReserve asserts that the Reserve
// function properly allocates additional
// space in the vector.
func TestReserve(t *testing.T) {
	vec, _ := setup()
	vec.Reserve(5)
	if len(*vec) != 10 {
		t.Errorf(EXPECTED, "len", 10, len(*vec))
	}

	for i, number := range *vec {
		if i < 5 {
			continue
		}
		if number != 0 {
			t.Errorf(EXPECTED, "val", 0, number)
		}
	}

}

// TestExtend asserts that the Extend function
// properly appends all the values from another
// collection into the vector.
func TestExtend(t *testing.T) {
	vec1, numbers := setup()
	vec2, _ := setup()
	vec1.Extend(vec2)
	for i, number := range *vec1 {
		if number != numbers[i%len(numbers)] {
			t.Errorf(EXPECTED, "val", numbers[i%len(numbers)], number)
		}
	}
}

// TestClear asserts that the Clear function
// properly sets every element to the zero
// value of its type.
func TestClear(t *testing.T) {
	vec, _ := setup()
	vec.Clear()
	for _, val := range *vec {
		if val != 0 {
			t.Errorf(EXPECTED, "val", 0, val)
		}
	}
}

// TestInsertBefore asserts that the InsertBefore
// function properly inserts an item into the Vector
// before the specified index.
func TestInsertBefore(t *testing.T) {
	vec, numbers := setup()

	for i := range numbers {
		n := len(*vec)
		m := len(numbers) + i

		if n != m {
			t.Errorf(EXPECTED, "len", m, n)
		}

		vec.InsertBefore(i, 0)
		val, err := vec.Get(i)

		if err != nil {
			t.Errorf(ERROR, err.Error())
		}

		if *val != 0 {
			t.Errorf(EXPECTED, "val", 0, *val)
		}
	}
}

// TestInsertAfter asserts that the InsertAfter
// function properly inserts an item into the Vector
// after the specified index
func TestInsertAfter(t *testing.T) {
	vec, numbers := setup()

	for i := range numbers {
		n := len(*vec)
		m := len(numbers) + i

		if n != m {
			t.Errorf(EXPECTED, "len", m, n)
		}

		vec.InsertAfter(i, 0)
		val, err := vec.Get(i + 1)

		if err != nil {
			t.Errorf(ERROR, err.Error())
		}

		if *val != 0 {
			t.Errorf(EXPECTED, "val", 0, *val)
		}
	}
}
