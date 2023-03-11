package iterator

// FilterFunc represents the predicate function passed
// to Filter. Items that would result in this function
// returning true will be included in the resulting
// iterator
type FilterFunc[TItem any] func(TItem) bool

type filterIterator[TItem any] struct {
	iter Iterator[TItem]
	pred FilterFunc[TItem]
}

// Filter returns an iterator over the items in the
// given iterator where the given predicate function
// returns true
func Filter[TItem any](
	iter Iterator[TItem],
	pred FilterFunc[TItem],
) Iterator[TItem] {
	return &filterIterator[TItem]{
		iter,
		pred,
	}
}

func (f *filterIterator[TItem]) Next() *TItem {
	for {
		val := f.iter.Next()
		if val == nil {
			return nil
		}

		if f.pred(*val) {
			return val
		}
	}
}
