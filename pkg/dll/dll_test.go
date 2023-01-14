package dll_test

import (
	"testing"

	"github.com/bdreece/gollections/pkg/dll"
	"github.com/bdreece/gollections/pkg/slice"
)

func TestDLLNew(t *testing.T) {
	d := dll.New[int]()
	if d.Count() != 0 {
		t.Error("Count != 0")
	}
}

func TestDLLFrom(t *testing.T) {
	d := dll.From[int](slice.From([]int{1, 2, 3}))
	if d.Count() != 3 {
		t.Error("Count != 0")
	}

	first := d.Iter().Next()
	if first == nil {
		t.Error("first is nil")
		return
	}
	if *first != 1 {
		t.Error("first != 1")
	}
}
