package iterator

// Iterator provides the required method Next
// for use in other iterator functions
type Iterator[TItem any] interface {
	// Next returns the next item in the
	// iterator, otherwise nil
	Next() *TItem
}

// IntoIterator provides the Iter method used
// to obtain an iterator over a collection
type IntoIterator[TItem any] interface {
	Iter() Iterator[TItem]
}

// From returns an iterator from an IntoIterator
func From[TItem any](into IntoIterator[TItem]) Iterator[TItem] {
	return into.Iter()
}

// ForEachFunc represents the predicate function
// passed to ForEach
type ForEachFunc[TItem any] func(TItem)

// ForEach iterates over all items in the
// given iterator, passing each to the given
// predicate function
func ForEach[TItem any](
	iter Iterator[TItem],
	pred ForEachFunc[TItem],
) {
	for val := iter.Next(); val != nil; val = iter.Next() {
		pred(*val)
	}
}
