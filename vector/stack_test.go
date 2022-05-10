package vector

import "testing"

// TestPush asserts that the Push
// function properly prepends an
// element onto the vector.
func TestPush(t *testing.T) {
	vec := New[int]()
	_, numbers := setup()

	for _, number := range numbers {
		vec.Push(number)
		val := []int(*vec)[0]
		if val != number {
			t.Errorf(EXPECTED, "val", number, val)
		}
	}
}

// TestPop asserts that the Pop
// function properly removes and
// returns the first element in
// the vector.
func TestPop(t *testing.T) {
	vec, numbers := setup()

	for i, number := range numbers {
		n := len(*vec)
		m := len(numbers) - i
		if n != m {
			t.Errorf(EXPECTED, "len", m, n)
		}
		elem, err := vec.Pop()
		if err != nil {
			t.Errorf(ERROR, err.Error())
		}
		if *elem != number {
			t.Errorf(EXPECTED, "val", number, *elem)
		}
	}
}
