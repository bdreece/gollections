package vector

import "testing"

// TestCollect asserts that the Collect
// function properly appends a variable
// number of elements onto the vector.
func TestCollect(t *testing.T) {
	vec := New[int]()
	_, numbers := setup()

	for i, number := range numbers {
		n := len(*vec)

		if n != i {
			t.Errorf(EXPECTED, "len", i, n)
		}

		vec.Collect(number)
		elem := []int(*vec)[n]
		if elem != number {
			t.Errorf(EXPECTED, "val", number, elem)
		}
	}
}
