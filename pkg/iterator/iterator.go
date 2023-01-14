package iterator

type Iterator[TItem any] interface {
	Next() *TItem
}

type IntoIterator[TItem any] interface {
	Iter() Iterator[TItem]
}

func From[TItem any](into IntoIterator[TItem]) Iterator[TItem] {
	return into.Iter()
}

type ForEachFunc[TItem any] func(TItem)

func ForEach[TItem any](
	iter Iterator[TItem],
	pred ForEachFunc[TItem],
) {
	for val := iter.Next(); val != nil; val = iter.Next() {
		pred(*val)
	}
}
