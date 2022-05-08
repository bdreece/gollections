package vector

import "testing"

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
