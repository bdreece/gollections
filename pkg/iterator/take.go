package iterator

type takeIterator[TItem any] struct {
	iter  Iterator[TItem]
	count int
	max   int
}

func Take[TItem any](iter Iterator[TItem], max int) Iterator[TItem] {
	return &takeIterator[TItem]{iter, 0, max}
}

func (t *takeIterator[TItem]) Next() *TItem {
	val := t.iter.Next()
	t.count++

	if val == nil || t.count > t.max {
		return nil
	}

	return val
}
