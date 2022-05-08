// MIT License
// Copyright (c) 2022 Brian Reece

package hashmap

import (
	"hash/maphash"

	"github.com/bdreece/gollections/errors"
	"github.com/bdreece/gollections/iterator"
	"github.com/bdreece/gollections/vector"
)

// MapPair is a key-value pair for the HashMap
// structure
type MapPair[K, V any] struct {
	Key K
	Val V
}

// HashMap is a hash map data structure,
// using hash/maphash for key hashing.
type HashMap[K comparable, V any] struct {
	*vector.Vector[map[K]V]
	maphash.Hash

	loadFactor float32
}

// New constructs a new HashMap
func New[K comparable, V any](loadFactor float32, capacity int) *HashMap[K, V] {
	vec := vector.New[map[K]V]()
	vec.Reserve(5)

	iterator.ForEach(vec.Iterator(), func(bucket *map[K]V) {
		*bucket = make(map[K]V)
	})

	return &HashMap[K, V]{
		vec,
		maphash.Hash{},
		loadFactor,
	}
}

func (m *HashMap[K, V]) hash(key K) int {
	bytes := any(key).([]byte)
	m.Hash.Write(bytes)
	return int(m.Hash.Sum64()) % len(*m.Vector)
}

// Get retrieves the value associated with a specified
// key. Returns nil, errors.NotFound if key does not
// exist.
func (m HashMap[K, V]) Get(key K) (*V, error) {
	hash := m.hash(key)
	bucket, err := m.Vector.Get(hash)
	if err != nil {
		return nil, err
	}
	val, ok := (*bucket)[key]
	if !ok {
		return nil, errors.NotFound{}
	}
	return &val, nil
}

// Set inserts a key value pair into the hash map,
// overwriting the value if the key already exists.
// Returns error on vector error.
func (m *HashMap[K, V]) Set(key K, val V) error {
	hash := m.hash(key)
	bucket, err := m.Vector.Get(hash)
	if err != nil {
		return err
	}
	(*bucket)[key] = val
	return nil
}

// Collect inserts a variable number of key-value pairs
// into the hash map. This method implements part of the
// iterator.Collection interface.
func (m *HashMap[K, V]) Collect(pairs ...MapPair[K, V]) {
	for _, pair := range pairs {
		m.Set(pair.Key, pair.Val)
	}
}

// Iterator returns an iterator over the key-value pairs
// in the hash map. This method implements part of the
// iterator.Collection interface.
func (m *HashMap[K, V]) Iterator() *Iterator[K, V] {
	return &Iterator[K, V]{m}
}
