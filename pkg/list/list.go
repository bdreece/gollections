package list

import (
	"github.com/bdreece/gollections/pkg/collection"
	"github.com/bdreece/gollections/pkg/slice"
)

// List provides an abstraction over the Slice
// interface, allowing construction from other
// collections
type List[TItem any] interface {
	slice.Slice[TItem]
}

// From creates a new List by concatenating the items
// from the given collection
func From[TItem any](c collection.Collection[TItem]) List[TItem] {
	s := slice.New[TItem](0, c.Count())
	s.Concat(c)
	return s
}
