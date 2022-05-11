package unordered

import (
	"testing"

	"github.com/bdreece/gollections/iterator"
	"github.com/bdreece/gollections/vector"
)

func TestKeys(t *testing.T) {
	ours, theirs := setup()
	their_keys := vector.New[string]()
	for key := range theirs {
		their_keys.Enqueue(key)
	}
	if err := iterator.ForEach(
		ours.Keys(),
		func(key *string) {
			if matches, err := iterator.Any(
				their_keys.IntoIterator(),
				func(their_key *string) bool {
					return *their_key == *key
				},
			); err != nil {
				t.Errorf(ERROR, err.Error())
			} else if !matches {
				t.Errorf(EXPECTED, "key", *key, "none")
			}
		},
	); err != nil {
		t.Errorf(ERROR, err.Error())
	}
}
func TestVals(t *testing.T) {
	ours, theirs := setup()
	their_vals := vector.New[int]()
	for _, val := range theirs {
		their_vals.Enqueue(val)
	}
	if err := iterator.ForEach(
		ours.Vals(),
		func(val *int) {
			if matches, err := iterator.Any(
				their_vals.IntoIterator(),
				func(their_val *int) bool {
					return *their_val == *val
				},
			); err != nil {
				t.Errorf(ERROR, err.Error())
			} else if !matches {
				t.Errorf(EXPECTED, "val", *val, "none")
			}
		},
	); err != nil {
		t.Errorf(ERROR, err.Error())
	}
}
