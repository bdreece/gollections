// MIT License
// Copyright (c) 2022 Brian Reece

package list

import "testing"

func TestPush(t *testing.T) {
	list := New[int]()
	_, numbers := setup()

	for _, number := range numbers {
		list.Push(number)
		val, err := list.PeekFront()
		if err != nil {
			t.Errorf(ERROR, err.Error())
		}
		if *val != number {
			t.Errorf(EXPECTED, "val", number, *val)
		}
	}
}

func TestPop(t *testing.T) {
	list, numbers := setup()

	for _, number := range numbers {
		val, err := list.Pop()
		if err != nil {
			t.Errorf(ERROR, err.Error())
		}
		if *val != number {
			t.Errorf(EXPECTED, "val", number, *val)
		}
	}
}
