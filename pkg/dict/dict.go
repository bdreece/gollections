package dict

import (
	"github.com/bdreece/gollections/pkg/collection"
	"github.com/bdreece/gollections/pkg/iterator"
	"github.com/bdreece/gollections/pkg/slice"
)

type Pair[TKey comparable, TValue any] struct {
	Key   TKey
	Value TValue
}

type Dict[TKey comparable, TValue any] interface {
	collection.Collection[Pair[TKey, TValue]]
	Get(TKey) *TValue
	Set(TKey, TValue)
	Remove(TKey) *TValue
}

type dict[TKey comparable, TValue any] map[TKey]TValue

func New[TKey comparable, TValue any]() Dict[TKey, TValue] {
	var d dict[TKey, TValue] = make(map[TKey]TValue)
	return &d
}

func From[TKey comparable, TValue any](
	c collection.Collection[Pair[TKey, TValue]],
) Dict[TKey, TValue] {
	d := New[TKey, TValue]()
	d.Concat(c)
	return d
}

func (d *dict[TKey, TValue]) Count() int {
	return len(*d)
}

func (d *dict[TKey, TValue]) Iter() iterator.Iterator[Pair[TKey, TValue]] {
	s := slice.New[Pair[TKey, TValue]](0, len(*d))
	for k, v := range *d {
		s.Add(Pair[TKey, TValue]{k, v})
	}

	return s.Iter()
}

func (d *dict[TKey, TValue]) Concat(
	into iterator.IntoIterator[Pair[TKey, TValue]],
) collection.Collection[Pair[TKey, TValue]] {
	return d.Collect(into.Iter())
}

func (d *dict[TKey, TValue]) Collect(
	iter iterator.Iterator[Pair[TKey, TValue]],
) collection.Collection[Pair[TKey, TValue]] {
	iterator.ForEach(iter, func(pair Pair[TKey, TValue]) {
		d.Append(pair)
	})

	return d
}

func (d *dict[TKey, TValue]) Append(
	pair Pair[TKey, TValue],
) collection.Collection[Pair[TKey, TValue]] {
	d.Set(pair.Key, pair.Value)
	return d
}

func (d *dict[TKey, TValue]) Get(key TKey) *TValue {
	if v, ok := (*d)[key]; ok {
		return &v
	}

	return nil
}

func (d *dict[TKey, TValue]) Set(key TKey, value TValue) {
	(*d)[key] = value
}

func (d *dict[TKey, TValue]) Remove(key TKey) *TValue {
	v, ok := (*d)[key]
	delete(*d, key)
	if ok {
		return &v
	}

	return nil
}
