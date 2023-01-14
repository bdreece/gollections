package iterator

type emptyIterator[TItem any] struct{}

func Empty[TItem any]() Iterator[TItem] {
	return emptyIterator[TItem]{}
}

func (e emptyIterator[TItem]) Next() *TItem {
	return nil
}
