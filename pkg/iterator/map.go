package iterator

type MapFunc[TInput any, TOutput any] func(TInput) TOutput

type mapIterator[TInput any, TOutput any] struct {
	iter Iterator[TInput]
	pred MapFunc[TInput, TOutput]
}

func Map[TInput any, TOutput any](
	iter Iterator[TInput],
	pred MapFunc[TInput, TOutput],
) Iterator[TOutput] {
	return &mapIterator[TInput, TOutput]{
		iter,
		pred,
	}
}

func (m *mapIterator[TInput, TOutput]) Next() *TOutput {
	val := m.iter.Next()
	if val == nil {
		return nil
	}
	next := m.pred(*val)
	return &next
}
