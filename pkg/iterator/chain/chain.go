package chain

import "github.com/bdreece/gollections/pkg/iterator"

// Chain provides chainable versions of the iterator functions
// avaiable in the parent package
type Chain[TInput, TOutput any] interface {
	Concat(iterator.Iterator[TInput]) Chain[TInput, TOutput]
	Enumerate(iterator.EnumerateFunc[TInput])
	Filter(iterator.FilterFunc[TInput]) Chain[TInput, TOutput]
	Find(iterator.FindFunc[TInput]) (*TInput, error)
	FlatMap(iterator.FlatMapFunc[TInput, TOutput]) Chain[TOutput, TOutput]
	ForEach(iterator.ForEachFunc[TInput])
	Map(iterator.MapFunc[TInput, TOutput]) Chain[TOutput, TOutput]
	Reduce(iterator.ReduceFunc[TInput, TOutput], TOutput) TOutput
	Take(int) Chain[TInput, TOutput]
	Value() iterator.Iterator[TInput]
}

type chain[TInput, TOutput any] struct {
	iterator.Iterator[TInput]
}

func From[TInput, TOutput any](iter iterator.Iterator[TInput]) Chain[TInput, TOutput] {
	return &chain[TInput, TOutput]{iter}
}

func (c *chain[TInput, TOutput]) Concat(
	second iterator.Iterator[TInput],
) Chain[TInput, TOutput] {
	return &chain[TInput, TOutput]{
		iterator.Concat(c.Iterator, second),
	}
}

func (c *chain[TInput, TOutput]) Enumerate(pred iterator.EnumerateFunc[TInput]) {
	iterator.Enumerate(c.Iterator, pred)
}

func (c *chain[TInput, TOutput]) Filter(
	pred iterator.FilterFunc[TInput],
) Chain[TInput, TOutput] {
	return &chain[TInput, TOutput]{
		iterator.Filter(c.Iterator, pred),
	}
}

func (c *chain[TInput, TOutput]) ForEach(pred iterator.ForEachFunc[TInput]) {
	iterator.ForEach(c.Iterator, pred)
}

func (c *chain[TInput, TOutput]) Find(pred iterator.FindFunc[TInput]) (*TInput, error) {
	return iterator.Find(c.Iterator, pred)
}

func (c *chain[TInput, TOutput]) FlatMap(
	pred iterator.FlatMapFunc[TInput, TOutput],
) Chain[TOutput, TOutput] {
	return &chain[TOutput, TOutput]{
		iterator.FlatMap(c.Iterator, pred),
	}
}

func (c *chain[TInput, TOutput]) Map(pred iterator.MapFunc[TInput, TOutput]) Chain[TOutput, TOutput] {
	return &chain[TOutput, TOutput]{
		iterator.Map(c.Iterator, pred),
	}
}

func (c *chain[TInput, TOutput]) Reduce(pred iterator.ReduceFunc[TInput, TOutput], initial TOutput) TOutput {
	return iterator.Reduce(c.Iterator, pred, initial)
}

func (c *chain[TInput, TOutput]) Take(n int) Chain[TInput, TOutput] {
	return &chain[TInput, TOutput]{
		iterator.Take(c.Iterator, n),
	}
}

func (c chain[TInput, TOutput]) Value() iterator.Iterator[TInput] {
	return c.Iterator
}
