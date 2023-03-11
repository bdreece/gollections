package iterator

type takeIterator[TItem any] struct {
	iter  Iterator[TItem]
	count int
	max   int
}

// Take returns an iterator over a subset of the given iterator,
// taking the first n items
func Take[TItem any](iter Iterator[TItem], n int) Iterator[TItem] {
	return &takeIterator[TItem]{iter, 0, n}
}

func (t *takeIterator[TItem]) Next() *TItem {
	val := t.iter.Next()
	t.count++

	if val == nil || t.count > t.max {
		return nil
	}

	return val
}
