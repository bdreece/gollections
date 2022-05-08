package vector

import "testing"

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
