package vector

import (
	"testing"

	"github.com/bdreece/gollections/iterator"
)

func TestForEach(t *testing.T) {
	vec, numbers := setup()
	i := 0
	if err := iterator.ForEach(
		vec.IntoIterator(),
		func(elem *int) {
			if *elem != numbers[i] {
				t.Errorf(EXPECTED, "val", numbers[i], *elem)
			}
			i++
		},
	); err != nil {
		t.Errorf(ERROR, err.Error())
	}
}

func TestEnumerate(t *testing.T) {
	vec, numbers := setup()
	if err := iterator.ForEach[iterator.EnumerateItem[int]](
		iterator.NewEnumerate(vec.IntoIterator()),
		func(elem *iterator.EnumerateItem[int]) {
			expected := numbers[elem.Index]
			got := elem.Item
			if got != expected {
				t.Errorf(EXPECTED, "val", expected, got)
			}
		},
	); err != nil {
		t.Errorf(ERROR, err.Error())
	}
}
