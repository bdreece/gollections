// MIT License
// Copyright (c) 2022 Brian Reece

package list

import "testing"

func TestEnqueue(t *testing.T) {
	list := New[int]()
	_, numbers := setup()

	for _, number := range numbers {
		list.Enqueue(number)
		got, err := list.PeekBack()
		if err != nil {
			t.Errorf(ERROR, err.Error())
		}
		if *got != number {
			t.Errorf(EXPECTED, "val", number, *got)
		}
	}
}

func TestDequeue(t *testing.T) {
	list, numbers := setup()

	for _, number := range numbers {
		val, err := list.Dequeue()
		if err != nil {
			t.Errorf(ERROR, err.Error())
		}
		if *val != number {
			t.Errorf(EXPECTED, "val", number, *val)
		}
	}
}
