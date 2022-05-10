package ordered

import (
	"github.com/bdreece/gollections/errors"
	"github.com/bdreece/gollections/iterator"
	"github.com/bdreece/gollections/maps"
	"github.com/bdreece/gollections/vector"
)

// Map provides an ordered map (type alias to
// built-in go map).
type Map[K comparable, V any] map[K]V

// New constructs a new ordered map
func New[K comparable, V any]() *Map[K, V] {
	m := make(Map[K, V])
	return &m
}

// Get retrieves a value from the ordered map
// associated with the specified key. Returns
// nil, errors.NotFound if the key does not
// exist. This method implements part of the
// gollections.Array interface.
func (m Map[K, V]) Get(key K) (*V, error) {
	if val, ok := map[K]V(m)[key]; ok {
		return &val, nil
	} else {
		return nil, errors.NotFound[K]{Key: key}
	}
}

// Set replaces the value associated with the
// specified key with the provided value. Returns
// errors.NotFound if the key does not exist.
// This method implements part of the gollections.Array
// interface.
func (m *Map[K, V]) Set(key K, val V) error {
	if _, ok := map[K]V(*m)[key]; !ok {
		return errors.NotFound[K]{Key: key}
	}
	map[K]V(*m)[key] = val
	return nil
}

// Ins inserts a value into the ordered map
// associated with the provided key, overwriting
// any value currently present. This method implements
// part of the gollections.Array interface. Note that
// this method does not throw errors, and only has the
// error return type to match the interface.
func (m *Map[K, V]) Ins(key K, val V) error {
	map[K]V(*m)[key] = val
	return nil
}

// Del removes and returns the value from the
// ordered map associated with the specified
// key. Returns nil, errors.NotFound if the
// key does not exist.
func (m *Map[K, V]) Del(key K) (*V, error) {
	if val, err := m.Get(key); err == nil {
		delete(map[K]V(*m), key)
		return val, nil
	} else {
		return nil, err
	}
}

// Keys returns an iterator over the keys present
// in the ordered map.
func (m *Map[K, V]) Keys() iterator.Iterator[K] {
	vec := vector.New[K]()
	for key := range map[K]V(*m) {
		vec.PushBack(key)
	}
	return &Iterator[K]{*vec, 0}
}

// Vals returns an iterator over the values present
// in the ordered map.
func (m *Map[K, V]) Vals() iterator.Iterator[V] {
	vec := vector.New[V]()
	for _, val := range map[K]V(*m) {
		vec.PushBack(val)
	}
	return &Iterator[V]{*vec, 0}
}

// IntoIterator returns an iterator over the key-value
// pairs present in the ordered map.
func (m *Map[K, V]) IntoIterator() iterator.Iterator[maps.Pair[K, V]] {
	return NewPairIterator(m)
}
