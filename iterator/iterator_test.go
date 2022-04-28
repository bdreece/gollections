// MIT License
// Copyright (c) 2022 Brian Reece

package iterator

import (
	"github.com/bdreece/gollections/vector"
	"testing"
)

const (
	EXPECTED string = "expected (%d), got (%d)\n"
	ERROR    string = "experienced error: \"%s\"\n"
)

func setup(t *testing.T) (*vector.Vector[int], []int) {
	t.Log("Setting up test")
	vec := vector.New[int]()
	t.Log("Created vector")
	numbers := []int{1, 2, 3, 4, 5}
	for _, number := range numbers {
		vec.PushBack(number)
		t.Logf("Pushed (%d) to back of vector\n", number)
	}
	return vec, numbers
}

func TestForEach(t *testing.T) {
	vec, numbers := setup(t)
	i := 0
	if err := ForEach[int](
		vec.Iterator(),
		func(item *int) {
			if *item != numbers[i] {
				t.Errorf(EXPECTED, numbers[i], item)
			}
			i++
		},
	); err != nil {
		t.Errorf(ERROR, err.Error())
	}
}

func TestEnumerate(t *testing.T) {
	vec, numbers := setup(t)
	i := 0
	if err := ForEach[EnumerateItem[int]](
		NewEnumerate[int](vec.Iterator()),
		func(item *EnumerateItem[int]) {
			if item.index != i {
				t.Errorf(EXPECTED, i, item.index)
			}
			if item.item != numbers[i] {
				t.Errorf(EXPECTED, numbers[i], item.item)
			}
			i++
		},
	); err != nil {
		t.Errorf(ERROR, err.Error())
	}
}
