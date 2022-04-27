package list

import "testing"

const (
	EXPECTED string = "expected %s: (%d), got: (%d)\n"
)

func TestNew(t *testing.T) {
	list := New[int]()
	if list.length != 0 {
		t.Errorf(EXPECTED, "len", 0, list.length)
	}
}

func TestPush(t *testing.T) {
	list := New[int]()
	numbers := []int{1, 2, 3, 4, 5}

	for i, number := range numbers {
		list.Push(number)
		if list.length != i+1 {
			t.Errorf(EXPECTED, "len", i+1, list.length)
		}
	}
}

func TestPop(t *testing.T) {
	list := New[int]()
	numbers := []int{1, 2, 3, 4, 5}

	for _, number := range numbers {
		list.Push(number)
	}

	for i, number := range numbers {
		if list.length != 5-i {
			t.Errorf(EXPECTED, "len", 5-i, list.length)
		}
		val := list.Pop()
		if *val != number {
			t.Errorf(EXPECTED, "val", number, *val)
		}
	}
}
