package vector

import "testing"

// TestPeek asserts that the Peek
// function properly returns the first
// element in the vector.
func TestPeek(t *testing.T) {
	vec, numbers := setup()

	for range numbers {
		n := len(*vec)
		m := len(numbers)
		if n != m {
			t.Errorf(EXPECTED, "len", m, n)
		}
		elem, err := vec.Peek()
		if err != nil {
			t.Errorf(ERROR, err.Error())
		}
		if *elem != numbers[0] {
			t.Errorf(EXPECTED, "val", numbers[0], *elem)
		}
	}
}
