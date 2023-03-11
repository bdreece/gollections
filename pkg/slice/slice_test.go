package slice_test

import (
	"fmt"
	"testing"

	"github.com/bdreece/gollections/pkg/iterator"
	"github.com/bdreece/gollections/pkg/slice"
)

func TestNewSlice(t *testing.T) {
	s := slice.New[int](3, 3)

	if s.Count() != 3 {
		t.Error("Count != 3")
	}
}

func TestFromSlice(t *testing.T) {
	s := slice.Marshal([]int{1, 2, 3})

	if s.Count() != 3 {
		t.Error("Count != 3")
	}

	first := s.First()
	if first == nil {
		t.Error("first is nil")
		return
	}

	if *first != 1 {
		t.Error("First != 1")
	}
}

func TestSliceEmptyFirst(t *testing.T) {
	s := slice.New[int](0, 0)
	v := s.First()
	if v != nil {
		t.Error("v != nil")
	}
}

func TestSliceEmptyLast(t *testing.T) {
	s := slice.New[int](0, 0)
	v := s.Last()
	if v != nil {
		t.Error("v != nil")
	}
}

func TestSliceAppend(t *testing.T) {
	s := slice.New[int](0, 0)

	v := s.Append(1).
		Iter().
		Next()

	if s.Count() != 1 {
		t.Error("Count != 1")
	}

	if *v != 1 {
		t.Error("First != 1")
	}
}

func TestSliceConcat(t *testing.T) {
	s := slice.New[int](0, 0)
	v := s.Concat(slice.Marshal([]int{1, 2, 3})).
		Iter().
		Next()

	if s.Count() != 3 {
		t.Error("Count != 3")
	}

	if *v != 1 {
		t.Error("First != 1")
	}
}

func TestSliceCollect(t *testing.T) {
	s := slice.New[string](0, 0)
	v := s.Collect(
		iterator.Map(
			slice.Marshal([]int{1, 2, 3}).Iter(),
			func(item int) string {
				return fmt.Sprint(item)
			},
		)).
		Iter().
		Next()

	if s.Count() != 3 {
		t.Error("Count != 3")
	}

	if *v != "1" {
		t.Error("First != \"1\"")
	}
}

func TestSliceAdd(t *testing.T) {
	s := slice.Marshal([]int{1, 2, 3})
	s.Add(4)

	if s.Count() != 4 {
		t.Error("Count != 4")
	}

	last := s.Last()
	if *last != 4 {
		t.Error("Last != 4")
	}
}

func TestSliceGet(t *testing.T) {
	s := slice.Marshal([]int{1, 2, 3})

	v, err := s.Get(0)

	if err != nil {
		t.Error(err)
	}

	if v == nil {
		t.Error("v is nil")
		return
	}

	if *v != 1 {
		t.Error("First != 1")
	}
}

func TestSliceBadGet(t *testing.T) {
	s := slice.Marshal([]int{1, 2, 3})
	v, err := s.Get(4)
	if err == nil {
		t.Error("err is nil")
	}

	if v != nil {
		t.Error("v != nil")
	}
}

func TestSliceSet(t *testing.T) {
	s := slice.Marshal([]int{1, 2, 3})
	err := s.Set(1, 3)
	if err != nil {
		t.Error(err)
	}

	if s.Count() != 3 {
		t.Error("Count != 3")
	}

	v, err := s.Get(1)
	if err != nil {
		t.Error(err)
	}

	if v == nil {
		t.Error("v is nil")
		return
	}

	if *v != 3 {
		t.Error("v != 3")
	}
}

func TestSliceBadSet(t *testing.T) {
	s := slice.Marshal([]int{1, 2, 3})
	err := s.Set(4, 5)
	if err == nil {
		t.Error("err is nil")
	}
}

func TestSliceRemove(t *testing.T) {
	s := slice.Marshal([]int{1, 2, 3})

	v, err := s.Remove(1)
	if err != nil {
		t.Error(err)
	}

	if s.Count() != 2 {
		t.Error("Count != 2")
	}

	if v == nil {
		t.Error("v is nil")
		return
	}

	if *v != 2 {
		t.Error("v != 2")
	}
}

func TestSliceBadRemove(t *testing.T) {
	s := slice.Marshal([]int{1, 2, 3})
	v, err := s.Remove(4)
	if err == nil {
		t.Error("err is nil")
	}

	if v != nil {
		t.Error("v != nil")
	}
}
