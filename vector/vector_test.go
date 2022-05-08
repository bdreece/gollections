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

func TestTruncate(t *testing.T) {
	vec, _ := setup()
	vec.Truncate()
	n := len(*vec)
	if n != 0 {
		t.Errorf(EXPECTED, "len", 0, n)
	}
}
