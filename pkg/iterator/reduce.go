package iterator

// ReduceFunc represents the predicate function passed
// to Reduce
type ReduceFunc[TItem, TAggregate any] func(TAggregate, TItem) TAggregate

// Reduce applies the given predicate function to
// all items in the given iterator, aggregating the
// result that is eventually returned.
func Reduce[TItem, TAggregate any](
	iter Iterator[TItem],
	reduce ReduceFunc[TItem, TAggregate],
	initial TAggregate,
) TAggregate {
	ForEach(iter, func(item TItem) {
		initial = reduce(initial, item)
	})
	return initial
}
