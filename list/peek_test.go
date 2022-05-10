// MIT License
// Copyright (c) 2022 Brian Reece

package list

import "testing"

func TestPeek(t *testing.T) {
	list, numbers := setup()

	for _, number := range numbers {
		val, err := list.Peek()
		if err != nil {
			t.Errorf(ERROR, err.Error())
		}
		if *val != number {
			t.Errorf(EXPECTED, "val", number, *val)
		}
		list.Pop()
	}
}
