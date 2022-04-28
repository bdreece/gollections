// MIT License
// Copyright (c) 2022 Brian Reece

package hashmap

import (
	"github.com/bdreece/gollections/vector"
	"hash/maphash"
)

type MapPair[K, V any] struct {
	key K
	val V
}

type HashMap[K comparable, V any] struct {
	*vector.Vector[map[K]V]
	maphash.Hash
}

func New[K comparable, V any]() HashMap[K, V] {
	return HashMap[K, V]{
		vector.New[map[K]V](),
		maphash.Hash{},
	}
}

func (m *HashMap[K, V]) hash(key K) uint64 {
	bytes := any(key).([]byte)
	m.Hash.Write(bytes)
	return m.Hash.Sum64()
}

func (m *HashMap[K, V]) Set(key K, val V) error {
	hash := m.hash(key)
	bucket, err := m.Vector.Get(int(hash))
	if err != nil {
		return err
	}
	(*bucket)[key] = val
	err = m.Vector.Set(int(hash), *bucket)
	if err != nil {
		return err
	}
	return nil
}

func (m *HashMap[K, V]) Collect(values ...MapPair[K, V]) {
	for _, value := range values {
		m.Set(value.key, value.val)
	}
}
