package iterator

type FilterFunc[TItem any] func(TItem) bool

type filterIterator[TItem any] struct {
	iter Iterator[TItem]
	pred FilterFunc[TItem]
}

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
