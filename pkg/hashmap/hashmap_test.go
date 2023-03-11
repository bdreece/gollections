package hashmap_test

import (
	"testing"

	"github.com/bdreece/gollections/pkg/hashmap"
	"github.com/bdreece/gollections/pkg/slice"
)

func TestHashmapNew(t *testing.T) {
	h := hashmap.New[string, int](10)
	if h.Count() != 0 {
		t.Error("Count != 0")
	}
}

func TestHashmapFrom(t *testing.T) {
	h := hashmap.From[string, int](
		slice.From([]hashmap.Pair[string, int]{
			{"apple", 1},
			{"banana", 2},
			{"cake", 3},
		}))

	if h.Count() != 3 {
		t.Errorf("Count == %d\n", h.Count())
	}
}
