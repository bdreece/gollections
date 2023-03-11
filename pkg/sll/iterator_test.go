package sll_test

import (
	"testing"

	"github.com/bdreece/gollections/pkg/slice"
	"github.com/bdreece/gollections/pkg/sll"
)

func TestSllIterator(t *testing.T) {
	s1 := sll.From[int](slice.Marshal([]int{1, 2, 3}))
	s2 := sll.New[int]()
	v := s2.Concat(s1).
		Iter().
		Next()

	if s2.Count() != 3 {
		t.Error("Count != 3")
	}
	if v == nil {
		t.Error("v is nil")
	}
	if *v != 1 {
		t.Error("v != 1")
	}
}
