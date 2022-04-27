package vector

import "testing"

const (
	EXPECTED string = "expected %s: (%d), got: (%d)\n"
)

func TestNew(t *testing.T) {
	vec := New[int]()
	if len(*vec) != 0 {
		t.Errorf(EXPECTED, "len", 0, len(*vec))
	}
}

func TestPush(t *testing.T) {
	vec := New[int]()
	numbers := []int{1, 2, 3, 4, 5}

	for i, number := range numbers {
		if len(*vec) != i {
			t.Errorf(EXPECTED, "len", i, len(*vec))
		}
		vec.Push(number)
	}
}

func TestPop(t *testing.T) {
	vec := New[int]()
	numbers := []int{1, 2, 3, 4, 5}

	for _, number := range numbers {
		vec.Push(number)
	}

	for i, number := range numbers {
		if len(*vec) != 5-i {
			t.Errorf(EXPECTED, "len", 5-i, len(*vec))
		}
		val := vec.Pop()
		if *val != 6-number {
			t.Errorf(EXPECTED, "val", 6-number, *val)
		}
	}
}
