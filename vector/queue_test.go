package vector

import "testing"

func TestEnqueue(t *testing.T) {
	vec := New[int]()
	_, numbers := setup()

	for i, number := range numbers {
		n := len(*vec)
		if n != i {
			t.Errorf(EXPECTED, "len", i, n)
		}
		vec.Enqueue(number)
		val := []int(*vec)[n]
		if val != number {
			t.Errorf(EXPECTED, "val", number, val)
		}
	}
}

func TestDequeue(t *testing.T) {
	vec, numbers := setup()

	for i, number := range numbers {
		n := len(*vec)
		m := len(numbers) - i

		if n != m {
			t.Errorf(EXPECTED, "len", m, n)
		}

		elem, err := vec.Dequeue()

		if err != nil {
			t.Errorf(ERROR, err.Error())
		}

		if *elem != number {
			t.Errorf(EXPECTED, "val", number, *elem)
		}
	}
}
