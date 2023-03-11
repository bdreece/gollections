package collection

import "github.com/bdreece/gollections/pkg/iterator"

// Collection provides shared behavior between all collections,
// including chainable mutations and other utilities
type Collection[TItem any] interface {
	iterator.IntoIterator[TItem]

	// Concat appends all items of the IntoIterator
	// to the end of this Collection
	Concat(iterator.IntoIterator[TItem]) Collection[TItem]

	// Collect appends all items of the Iterator
	// to the end of this Collection
	Collect(iterator.Iterator[TItem]) Collection[TItem]

	// Append appends an item to the end of
	// this Collection
	Append(TItem) Collection[TItem]

	// Count returns the number of items in this
	// Collection
	Count() int
}
