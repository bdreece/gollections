package iterator

type FlatMapFunc[TInput, TOutput any] MapFunc[TInput, IntoIterator[TOutput]]

func FlatMap[TInput, TOutput any](
	iter Iterator[TInput],
	pred FlatMapFunc[TInput, TOutput],
) Iterator[TOutput] {
	out := Empty[TOutput]()
	ForEach(iter, func(item TInput) {
		out = Chain(out, pred(item).Iter())
	})
	return out
}
