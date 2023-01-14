package iterator

type EnumerateFunc[TItem any] func(TItem, int)

func Enumerate[TItem any](iter Iterator[TItem], pred EnumerateFunc[TItem]) {
	i := 0
	for val := iter.Next(); val != nil; val = iter.Next() {
		pred(*val, i)
		i++
	}
}
