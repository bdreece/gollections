package hashmap_test

import (
	"fmt"
	"testing"

	"github.com/bdreece/gollections/pkg/dict"
	"github.com/bdreece/gollections/pkg/hashmap"
	"github.com/bdreece/gollections/pkg/iterator"
	"github.com/bdreece/gollections/pkg/slice"
)

func TestHashmapNew(t *testing.T) {
	h := hashmap.New[string, int](10)
	if h.Count() != 0 {
		t.Error("Count != 0")
	}
}

func setup() hashmap.HashMap[string, int] {
	return hashmap.From[string, int](
		slice.Marshal([]dict.Pair[string, int]{
			{"apple", 1},
			{"banana", 2},
			{"cake", 3},
		}))
}

func TestHashmapFrom(t *testing.T) {
	h := setup()
	if h.Count() != 3 {
		t.Errorf("Count == %d\n", h.Count())
	}
}

func TestHashmapSet(t *testing.T) {
	h := setup()
	h.Set("dog", 4)

	if h.Count() != 4 {
		t.Errorf("Count == %d\n", h.Count())
	}
}

func TestHashmapRemove(t *testing.T) {
	h := setup()
	h.Remove("banana")

	if h.Count() != 2 {
		t.Errorf("Count == %d\n", h.Count())
	}
}

func TestHashmapCollect(t *testing.T) {
	h := setup()
	h.Collect(iterator.Map(
		iterator.Range(0, 100),
		func(i int) dict.Pair[string, int] {
			return dict.Pair[string, int]{fmt.Sprint(i), i}
		}))

	if h.Count() != 103 {
		t.Errorf("Count == %d\n", h.Count())
	}
}
