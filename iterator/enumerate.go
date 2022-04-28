package iter

type EnumerateItem[T any] struct {
	item  T
	index int
}

type Enumerate[T any] struct {
	iter  Iterator[T]
	index int
}

func NewEnumerate[T any](iter Iterator[T]) Enumerate[T] {
	return Enumerate[T]{
		iter:  iter,
		index: 0,
	}
}

func (e *Enumerate[T]) Next() (*EnumerateItem[T], error) {
	item, err := e.iter.Next()
	enum_item := EnumerateItem[T]{
		item:  *item,
		index: e.index,
	}
	e.index += 1
	if err != nil {
		return &enum_item, err
	}
	return &enum_item, nil
}
