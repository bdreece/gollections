package iterator

type ReduceFunc[TItem, TAggregate any] func(TAggregate, TItem) TAggregate

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
