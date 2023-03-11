package dict

import (
	"github.com/bdreece/gollections/pkg/collection"
	"github.com/bdreece/gollections/pkg/iterator"
	"github.com/bdreece/gollections/pkg/slice"
)

// Pair represents a key-value pair in a Dict
type Pair[TKey comparable, TValue any] struct {
	Key   TKey
	Value TValue
}

// Dict provides an associative array of key-value
// pairs, based on the primitive Go map
type Dict[TKey comparable, TValue any] interface {
	collection.Collection[Pair[TKey, TValue]]

	// Get returns the value associated with
	// the given key if it exists, otherwise nil
	Get(TKey) *TValue

	// Set assigns the given value with the given
	// key, overwriting an existing value if it
	// exists
	Set(TKey, TValue)

	// Remove removes the key-value pair associated
	// with the given key, returning the previous value
	// if it exists, otherwise nil
	Remove(TKey) *TValue
}

type dict[TKey comparable, TValue any] map[TKey]TValue

// New creates a new Dict with an arbitrary initial size
func New[TKey comparable, TValue any]() Dict[TKey, TValue] {
	var d dict[TKey, TValue] = make(map[TKey]TValue)
	return &d
}

func Marshal[TKey comparable, TValue any](
	m map[TKey]TValue,
) Dict[TKey, TValue] {
	var d dict[TKey, TValue] = m
	return &d
}

// From creates a new Dict by concatenating the items
// of the given collection
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
