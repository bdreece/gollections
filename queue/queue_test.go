package queue

import "testing"

const (
	EXPECTED_LEN string = "expected len: (%d), got: (%d)\n"
	EXPECTED_VAL string = "expected val: (%d), got: (%d)\n"
)

func TestNewQueue(t *testing.T) {
	queue := New[int]()
	if len(*queue) != 0 {
		t.Errorf(EXPECTED_LEN, 0, len(*queue))
	}
}

func TestEnqueue(t *testing.T) {
	queue := New[int]()
	numbers := []int{1, 2, 3, 4, 5}
	for i, number := range numbers {
		if len(*queue) != i {
			t.Errorf(EXPECTED_LEN, i, len(*queue))
		}
		queue.Enqueue(number)
	}
	if len(*queue) != 5 {
		t.Errorf(EXPECTED_LEN, 5, len(*queue))
	}
}

func TestDequeue(t *testing.T) {
	queue := New[int]()
	numbers := []int{1, 2, 3, 4, 5}
	for i, number := range numbers {
		if len(*queue) != i {
			t.Errorf(EXPECTED_LEN, i, len(*queue))
		}
		queue.Enqueue(number)
	}

	for _, number := range numbers {
		val := queue.Dequeue()
		if *val != number {
			t.Errorf(EXPECTED_VAL, number, *val)
		}
	}
}

func TestPeek(t *testing.T) {
	queue := New[int]()
	numbers := []int{1, 2, 3, 4, 5}

	for i, number := range numbers {
		if i != len(*queue) {
			t.Errorf(EXPECTED_LEN, i, len(*queue))
		}
		queue.Enqueue(number)
	}

	for range numbers {
		val := queue.Peek()
		if *val != 1 {
			t.Errorf(EXPECTED_VAL, 1, *val)
		}
	}
}
