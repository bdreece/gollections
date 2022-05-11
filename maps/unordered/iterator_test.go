package unordered

import (
	"testing"

	"github.com/bdreece/gollections/iterator"
	"github.com/bdreece/gollections/maps"
	"github.com/bdreece/gollections/vector"
)

func TestForEach(t *testing.T) {
	ours, theirs := setup()

	if err := iterator.ForEach(
		ours.IntoIterator(),
		func(item *maps.Pair[string, int]) {
			if item.Val != theirs[item.Key] {
				t.Errorf(EXPECTED, "val", theirs[item.Key], item.Val)
			}
		},
	); err != nil {
		t.Errorf(ERROR, err.Error())
	}
}

func TestFromIterator(t *testing.T) {
	ours_expected, theirs := setup()
	ours_got := New[string, int]()
	their_pairs := vector.New[maps.Pair[string, int]]()
	for key, val := range theirs {
		their_pairs.Enqueue(maps.Pair[string, int]{Key: key, Val: val})
	}
	ours_got.FromIterator(their_pairs.IntoIterator())
	if err := iterator.ForEach(
		ours_expected.IntoIterator(),
		func(pair *maps.Pair[string, int]) {
			val, err := ours_got.Get(pair.Key)
			if err != nil {
				t.Errorf(ERROR, err.Error())
			}
			if *val != pair.Val {
				t.Errorf(EXPECTED, "val", pair.Val, *val)
			}
		},
	); err != nil {
		t.Errorf(ERROR, err.Error())
	}
}
