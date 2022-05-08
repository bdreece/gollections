package vector

import "testing"

func TestAppend(t *testing.T) {
	vec := New[int]()
	_, numbers := setup()

	for i, number := range numbers {
		n := len(*vec)

		if n != i {
			t.Errorf(EXPECTED, "len", i, n)
		}

		vec.Append(number)
		elem := []int(*vec)[n]
		if elem != number {
			t.Errorf(EXPECTED, "val", number, elem)
		}
	}
}
