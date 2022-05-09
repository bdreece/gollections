// MIT License
// Copyright (c) 2022 Brian Reece

package ringbuf

import (
	"testing"

	"github.com/bdreece/gollections/iterator"
)

func TestForEach(t *testing.T) {
	ringbuf, numbers := setup()
	i := 0
	if err := iterator.ForEach(ringbuf.IntoIterator(), func(item *int) {
		if *item != numbers[i] {
			t.Errorf(EXPECTED, "val", numbers[i], *item)
		}
		i++
	}); err != nil {
		t.Errorf(ERROR, err.Error())
	}
}
