package ordered

import (
	"github.com/bdreece/gollections/errors"
	super "github.com/bdreece/gollections/maps"
	"github.com/bdreece/gollections/vector"
)

type Iterator[T any] struct {
	vector.Vector[T]
	int
}

func NewPairIterator[K comparable, V any](m *Map[K, V]) *Iterator[super.Pair[K, V]] {
	vec := vector.New[super.Pair[K, V]]()
	for key, val := range map[K]V(*m) {
		vec.PushBack(super.Pair[K, V]{Key: key, Val: val})
	}
	return &Iterator[super.Pair[K, V]]{vec, 0}
}

func (iter *Iterator[T]) Next() (*T, error) {
	if iter.int >= len([]T(iter.Vector)) {
		return nil, errors.Empty{}
	}
	iter.int++
	return iter.Vector.Get(iter.int - 1)
}
