package iterator

// EnumerateFunc represents the predicate function passed to Enumerate
type EnumerateFunc[TItem any] func(TItem, int)

// Enumerate provides ForEach functionality, but also passes item index
// to the predicate function
func Enumerate[TItem any](iter Iterator[TItem], pred EnumerateFunc[TItem]) {
	i := 0
	for val := iter.Next(); val != nil; val = iter.Next() {
		pred(*val, i)
		i++
	}
}
