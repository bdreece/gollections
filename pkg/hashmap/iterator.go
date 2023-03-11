package hashmap

import (
	"github.com/bdreece/gollections/pkg/dict"
	"github.com/bdreece/gollections/pkg/iterator"
)

func (h *hashMap[TKey, TValue]) Iter() iterator.Iterator[dict.Pair[TKey, TValue]] {
	return iterator.FlatMap(h.buckets.Iter(), func(
		l dict.Dict[TKey, TValue],
	) iterator.IntoIterator[dict.Pair[TKey, TValue]] {
		return l
	})
}
