package iterator

// FlatMapFunc represents the predicate function passed to
// FlatMap
type FlatMapFunc[TInput, TOutput any] MapFunc[TInput, IntoIterator[TOutput]]

// FlatMap returns an iterator over the transformed and
// flattened items from the given iterator, according
// to the predicate function
func FlatMap[TInput, TOutput any](
	iter Iterator[TInput],
	pred FlatMapFunc[TInput, TOutput],
) Iterator[TOutput] {
	out := Empty[TOutput]()
	ForEach(iter, func(item TInput) {
		out = Concat(out, pred(item).Iter())
	})
	return out
}
