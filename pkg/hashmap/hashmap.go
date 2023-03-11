package hashmap

import (
	"fmt"
	"hash/maphash"

	"github.com/bdreece/gollections/pkg/collection"
	"github.com/bdreece/gollections/pkg/dict"
	"github.com/bdreece/gollections/pkg/iterator"
	"github.com/bdreece/gollections/pkg/slice"
)

type HashMap[TKey comparable, TValue any] interface {
	collection.Collection[dict.Pair[TKey, TValue]]
	Get(TKey) (*TValue, error)
	Set(TKey, TValue) error
	Remove(TKey) (*TValue, error)
}

type hashMap[TKey comparable, TValue any] struct {
	buckets  slice.Slice[dict.Dict[TKey, TValue]]
	seed     maphash.Seed
	capacity int
}

func New[TKey comparable, TValue any](capacity int) HashMap[TKey, TValue] {
	s := slice.New[dict.Dict[TKey, TValue]](0, capacity)
	s.Collect(iterator.Map(
		iterator.Range(0, capacity),
		func(_ int) dict.Dict[TKey, TValue] {
			return dict.New[TKey, TValue]()
		}))

	return &hashMap[TKey, TValue]{
		buckets:  s,
		seed:     maphash.MakeSeed(),
		capacity: capacity,
	}
}

func From[TKey comparable, TValue any](
	c collection.Collection[dict.Pair[TKey, TValue]],
) HashMap[TKey, TValue] {
	h := New[TKey, TValue](c.Count())
	h.Concat(c)
	return h
}

func (h *hashMap[TKey, TValue]) Count() int {
	return iterator.Reduce(h.buckets.Iter(),
		func(count int, d dict.Dict[TKey, TValue]) int {
			return count + d.Count()
		}, 0)
}

func (h *hashMap[TKey, TValue]) Concat(
	into iterator.IntoIterator[dict.Pair[TKey, TValue]],
) collection.Collection[dict.Pair[TKey, TValue]] {
	return h.Collect(into.Iter())
}

func (h *hashMap[TKey, TValue]) Collect(
	iter iterator.Iterator[dict.Pair[TKey, TValue]],
) collection.Collection[dict.Pair[TKey, TValue]] {
	iterator.ForEach(iter, func(item dict.Pair[TKey, TValue]) {
		h.Append(item)
	})
	return h
}

func (h *hashMap[TKey, TValue]) Append(
	item dict.Pair[TKey, TValue],
) collection.Collection[dict.Pair[TKey, TValue]] {
	h.Set(item.Key, item.Value)
	return h
}

func (h *hashMap[TKey, TValue]) Get(key TKey) (*TValue, error) {
	hash := h.hash(key)
	bucket, err := h.buckets.Get(hash)
	if err != nil {
		return nil, err
	}

	return (*bucket).Get(key), nil
}

func (h *hashMap[TKey, TValue]) Set(key TKey, value TValue) error {
	hash := h.hash(key)
	bucket, err := h.buckets.Get(hash)
	if err != nil {
		return err
	}

	(*bucket).Set(key, value)
	h.buckets.Set(hash, *bucket)

	numBuckets := h.buckets.Count()
	if h.Count() > int(0.75*float32(numBuckets*numBuckets)) {
		newHashmap := rehash(h, h.capacity*2)
		*h = *newHashmap
	}

	return nil
}

func (h *hashMap[TKey, TValue]) Remove(key TKey) (*TValue, error) {
	hash := h.hash(key)
	bucket, err := h.buckets.Get(hash)
	if err != nil {
		return nil, err
	}

	item := (*bucket).Remove(key)
	if item == nil {
		return nil, iterator.ErrNotFound
	}

	h.buckets.Set(hash, *bucket)
	return item, nil
}

func rehash[TKey comparable, TValue any](h *hashMap[TKey, TValue], capacity int) *hashMap[TKey, TValue] {
	iter := h.Iter()
	s := slice.New[dict.Dict[TKey, TValue]](0, capacity)
	s.Collect(iterator.Map(
		iterator.Range(0, capacity),
		func(_ int) dict.Dict[TKey, TValue] {
			return dict.New[TKey, TValue]()
		}))

	rehashed := &hashMap[TKey, TValue]{
		buckets:  s,
		seed:     h.seed,
		capacity: capacity,
	}

	rehashed.Collect(iter)
	return rehashed
}

func (h *hashMap[TKey, TValue]) hash(key TKey) int {
	var hash maphash.Hash
	hash.WriteString(fmt.Sprint(key))
	return int(hash.Sum64() % uint64(h.capacity))
}
