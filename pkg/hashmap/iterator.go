package hashmap

import (
	"github.com/bdreece/gollections/pkg/iterator"
	"github.com/bdreece/gollections/pkg/list"
)

func (h *hashMap[TKey, TValue]) Iter() iterator.Iterator[Pair[TKey, TValue]] {
	return iterator.FlatMap(h.list.Iter(), func(
		l list.List[Pair[TKey, TValue]],
	) iterator.IntoIterator[Pair[TKey, TValue]] {
		return l
	})
}
