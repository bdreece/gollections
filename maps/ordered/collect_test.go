package ordered

import (
	"testing"

	"github.com/bdreece/gollections/iterator"
	"github.com/bdreece/gollections/maps"
	"github.com/bdreece/gollections/vector"
)

func TestCollect(t *testing.T) {
	ours, _ := setup()
	keys := vector.New[string]()
	vals := vector.New[int]()
	keys.Collect("filet", "grahams", "halibut")
	vals.Collect(5, 7, 8)
	if err := iterator.ForEach[iterator.ZipItem[string, int]](
		iterator.NewZip(keys.IntoIterator(), vals.IntoIterator()),
		func(item *iterator.ZipItem[string, int]) {
			ours.Collect(maps.Pair[string, int]{Key: *item.A, Val: *item.B})
		},
	); err != nil {
		t.Errorf(ERROR, err.Error())
	}

	for _, key := range *keys {
		if _, err := ours.Get(key); err != nil {
			t.Errorf(ERROR, err.Error())
		}
	}
}
