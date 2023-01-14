package sll_test

import (
	"testing"

	"github.com/bdreece/gollections/pkg/slice"
	"github.com/bdreece/gollections/pkg/sll"
)

func TestSllNew(t *testing.T) {
	s := sll.New[int]()
	if s.Count() != 0 {
		t.Error("Count != 0")
	}
}

func TestSLLFrom(t *testing.T) {
	s := sll.From[int](slice.From([]int{1, 2, 3}))
	if s.Count() != 3 {
		t.Error("Count != 0")
	}
}

func TestSllAppend(t *testing.T) {
	s := sll.New[int]()
	v := s.Append(1).Iter().Next()

	if s.Count() != 1 {
		t.Error("Count != 1")
	}
	if v == nil {
		t.Error("v is nil")
		return
	}
	if *v != 1 {
		t.Error("v != 1")
	}
}

func TestSLLPush(t *testing.T) {
	s := sll.New[int]()
	s.Push(1)

	if s.Count() != 1 {
		t.Error("Count != 1")
	}

	first := s.Iter().Next()
	if first == nil {
		t.Error("First is nil")
		return
	}
	if *first != 1 {
		t.Error("First != 1")
	}
}

func TestSLLPop(t *testing.T) {
	s := sll.From[int](slice.From([]int{1, 2, 3}))
	v := s.Pop()
	if s.Count() != 2 {
		t.Error("Count != 2")
	}
	if v == nil {
		t.Error("v is nil")
		return
	}
	if *v != 1 {
		t.Error("v != 1")
	}
}

func TestSLLEmptyPop(t *testing.T) {
	s := sll.New[int]()
	v := s.Pop()
	if v != nil {
		t.Error("v != nil")
	}
}
