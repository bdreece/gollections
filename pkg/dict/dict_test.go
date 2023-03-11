package dict

import (
	"testing"

	"github.com/bdreece/gollections/pkg/iterator"
	"github.com/bdreece/gollections/pkg/slice"
)

func setup() Dict[string, int] {
	return From[string, int](
		slice.From([]Pair[string, int]{
			{"apple", 1},
			{"banana", 2},
		}))
}

func TestDictNew(t *testing.T) {
	d := New[string, int]()
	if d.Count() != 0 {
		t.Errorf("Count == %d\n", d.Count())
	}
}

func TestDictFrom(t *testing.T) {
	d := setup()
	if d.Count() != 2 {
		t.Errorf("Count == %d\n", d.Count())
	}
}

func TestDictGet(t *testing.T) {
	d := setup()
	v := d.Get("banana")

	if v == nil {
		t.Error("v is nil")
		return
	}

	if *v != 2 {
		t.Errorf("v == %d\n", *v)
	}
}

func TestDictSet(t *testing.T) {
	d := setup()
	d.Set("cat", 3)

	v := d.Get("cat")

	if v == nil {
		t.Error("v is nil")
		return
	}

	if *v != 3 {
		t.Errorf("v == %d\n", *v)
	}
}

func TestDictRemove(t *testing.T) {
	d := setup()

	d.Remove("banana")

	if d.Count() != 1 {
		t.Errorf("Count == %d\n", d.Count())
	}
}

func TestDictIter(t *testing.T) {
	s := slice.From([]Pair[string, int]{
		{"cat", 3},
		{"dog", 4},
	})

	d := setup().Concat(s)

	filtered := iterator.Filter(d.Iter(), func(item Pair[string, int]) bool {
		return len(item.Key) == 3
	})

	str := iterator.Reduce(filtered, func(prev string, item Pair[string, int]) string {
		return prev + item.Key
	}, "")

	if str != "catdog" {
		t.Errorf("str == %s\n", str)
	}
}
