// MIT License
// Copyright (c) 2022 Brian Reece

package iterator

import (
	"fmt"
	"testing"

	"github.com/bdreece/gollections/errors"
)

const (
	EXPECTED string = "expected %s: (%d), got (%d)\n"
	ERROR    string = "experienced error: \"%s\"\n"
)

// Slice provides a test collection for iteration
// (need to do this to prevent import loops).
type Slice[T any] []T

// IntoIterator implements the IntoIterator interface.
func (s *Slice[T]) IntoIterator() Iterator[T] {
	return NewSliceIter(*s...)
}

// FromIterator implements the FromIterator interface.
func (s *Slice[T]) FromIterator(iter Iterator[T]) error {
	for {
		elem, err := iter.Next()
		if elem == nil {
			break
		}
		if err != nil {
			return err
		}
		*s = append([]T(*s), *elem)
	}
	return nil
}

func (s *Slice[T]) Append(elem T) error {
	*s = append([]T(*s), elem)
	return nil
}

// SliceIter provides a rudimentary iterator
// for testing
type SliceIter[T any] struct {
	data  []T
	index int
}

// NewSliceIter constructs a new SliceIter
func NewSliceIter[T any](data ...T) *SliceIter[T] {
	return &SliceIter[T]{data, 0}
}

// Next returns the next item in the SliceIter.
// This method implements the Iterator interface.
func (i *SliceIter[T]) Next() (*T, error) {
	if i.index >= len(i.data) {
		return nil, errors.Empty{}
	}
	i.index++
	return &i.data[i.index-1], nil
}

// TestAny asserts that the Any function returns
// true if for any item in the iterator, the predicate
// function returns true.
func TestAny(t *testing.T) {
	iter := NewSliceIter(1, 2, 3, 4, 5)
	even, err := Any[int](iter, func(item *int) bool {
		return *item%2 == 0
	})
	if err != nil {
		t.Errorf(ERROR, err.Error())
	}
	if !even {
		t.Errorf(EXPECTED, "val", 0, 1)
	}
}

// TestAll asserts that the All function returns
// true if for all items in the iterator, the predicate
// function returns true.
func TestAll(t *testing.T) {
	iter := NewSliceIter(1, 2, 3, 4, 5)
	positive, err := All[int](iter, func(item *int) bool {
		return *item > 0
	})
	if err != nil {
		t.Errorf(ERROR, err.Error())
	}
	if !positive {
		t.Errorf(EXPECTED, "val", 0, 1)
	}
}

// TestNone asserts that the None function returns
// true if for all items in the iterator, the predicate
// function returns false
func TestNone(t *testing.T) {
	iter := NewSliceIter(1, 2, 3, 4, 5)
	odd, err := None[int](iter, func(item *int) bool {
		return *item%2 == 0
	})
	if err != nil {
		t.Errorf(ERROR, err.Error())
	}
	if odd {
		t.Errorf(EXPECTED, "val", 2, 1)
	}

}

// TestForEach asserts that the ForEach function
// properly applies a function to each item in the
// iterator.
func TestForEach(t *testing.T) {
	iter := NewSliceIter(1, 2, 3, 4, 5)
	i := 1
	if err := ForEach[int](iter, func(item *int) {
		if *item != i {
			t.Errorf(EXPECTED, "val", i, *item)
		}
		i++
	}); err != nil {
		t.Errorf(ERROR, err.Error())
	}
}

// TestReduce asserts that the Reduce function properly
// reduces the items in the iterator to a single value.
func TestReduce(t *testing.T) {
	iter := NewSliceIter(1, 2, 3, 4, 5)
	total, err := Reduce[int](iter, func(total *int, item int) { *total += item })
	if err != nil {
		t.Errorf(ERROR, err.Error())
	}
	if *total != 15 {
		t.Errorf(EXPECTED, "val", 15, *total)
	}
}

// TestFold asserts that the Fold function properly
// folds the items in the iterator to a single value,
// initialized with a specified value.
func TestFold(t *testing.T) {
	iter := NewSliceIter(1, 2, 3, 4, 5)
	init := 10
	total, err := Fold[int](iter, &init, func(total *int, item int) { *total += item })
	if err != nil {
		t.Errorf(ERROR, err.Error())
	}
	if *total != 25 {
		t.Errorf(EXPECTED, "val", 25, *total)
	}
}

func TestMap(t *testing.T) {
	num_iter := NewSliceIter(1, 2, 3, 4, 5)
	str_iter, err := Map[int, string](num_iter, new(Slice[string]), func(item *int) string {
		return fmt.Sprint(*item)
	})
	if err != nil {
		t.Errorf(ERROR, err.Error())
	}
	if err := ForEach[ZipItem[int, string]](
		NewZip[int](num_iter, str_iter),
		func(item *ZipItem[int, string]) {
			if fmt.Sprint(item.A) != *item.B {
				t.Errorf(EXPECTED, "val", item.A, item.A)
			}
		},
	); err != nil {
		t.Errorf(ERROR, err.Error())
	}
}
