package list

import (
	"github.com/bdreece/gollections/pkg/collection"
	"github.com/bdreece/gollections/pkg/slice"
)

type List[TItem any] interface {
	slice.Slice[TItem]
}

func From[TItem any](c collection.Collection[TItem]) List[TItem] {
	s := slice.New[TItem](0, c.Count())
	s.Concat(c)
	return s
}
