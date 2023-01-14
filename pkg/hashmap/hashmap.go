package hashmap

import (
	"fmt"
	"hash/maphash"

	"github.com/bdreece/gollections/pkg/collection"
	"github.com/bdreece/gollections/pkg/iterator"
	"github.com/bdreece/gollections/pkg/list"
	"github.com/bdreece/gollections/pkg/slice"
)

type Pair[TKey comparable, TValue any] struct {
	Key   TKey
	Value TValue
}

type HashMap[TKey comparable, TValue any] interface {
	collection.Collection[Pair[TKey, TValue]]
	Get(TKey) (*TValue, error)
	Set(TKey, TValue) error
	Remove(TKey) (*TValue, error)
}

type hashMap[TKey comparable, TValue any] struct {
	list     list.List[list.List[Pair[TKey, TValue]]]
	seed     maphash.Seed
	capacity int
}

func New[TKey comparable, TValue any](capacity int) HashMap[TKey, TValue] {
	return &hashMap[TKey, TValue]{
		list:     slice.New[list.List[Pair[TKey, TValue]]](0, capacity),
		seed:     maphash.MakeSeed(),
		capacity: capacity,
	}
}

func (h *hashMap[TKey, TValue]) Count() int {
	return h.list.Count()
}

func (h *hashMap[TKey, TValue]) Concat(
	into iterator.IntoIterator[Pair[TKey, TValue]],
) collection.Collection[Pair[TKey, TValue]] {
	return h.Collect(into.Iter())
}

func (h *hashMap[TKey, TValue]) Collect(
	iter iterator.Iterator[Pair[TKey, TValue]],
) collection.Collection[Pair[TKey, TValue]] {
	iterator.ForEach(iter, func(item Pair[TKey, TValue]) {
		h.Append(item)
	})
	return h
}

func (h *hashMap[TKey, TValue]) Append(
	item Pair[TKey, TValue],
) collection.Collection[Pair[TKey, TValue]] {
	h.Set(item.Key, item.Value)
	return h
}

func (h *hashMap[TKey, TValue]) Get(key TKey) (*TValue, error) {
	hash := h.hash(key)
	bucket, err := h.list.Get(hash)
	if err != nil {
		return nil, err
	}

	pair, err := iterator.Find((*bucket).Iter(),
		func(p Pair[TKey, TValue]) bool {
			return p.Key == key
		})

	if err != nil {
		return nil, err
	}

	return &pair.Value, nil
}

func (h *hashMap[TKey, TValue]) Set(key TKey, value TValue) error {
	hash := h.hash(key)
	bucket, err := h.list.Get(hash)
	if err != nil {
		return err
	}

	iterator.Enumerate((*bucket).Iter(),
		func(pair Pair[TKey, TValue], index int) {
			if pair.Key == key {
				err = (*bucket).Set(index, Pair[TKey, TValue]{key, value})
			}
		})

	if err != nil {
		return err
	}

	h.list.Set(hash, *bucket)
	return nil
}

func (h *hashMap[TKey, TValue]) Remove(key TKey) (*TValue, error) {
	hash := h.hash(key)
	bucket, err := h.list.Get(hash)
	if err != nil {
		return nil, err
	}

	var item *TValue
	iterator.Enumerate((*bucket).Iter(),
		func(pair Pair[TKey, TValue], index int) {
			if pair.Key == key {
				item = &pair.Value
				(*bucket).Remove(index)
			}
		})

	if item == nil {
		return nil, iterator.ErrNotFound
	}

	h.list.Set(hash, *bucket)
	return item, nil
}

func (h *hashMap[TKey, TValue]) hash(key TKey) int {
	var hash maphash.Hash
	hash.WriteString(fmt.Sprint(key))
	return h.capacity % int(hash.Sum64())
}
