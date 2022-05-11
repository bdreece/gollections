package hash

import "github.com/bdreece/gollections/maps"

type Iterator[K comparable, V any] struct {
	*Map[K, V]
}

func (i *Iterator[K, V]) Next() (*maps.Pair[K, V], error) {
	// TODO: figure out how to aggregate pairs from vector of maps
	return nil, nil
}
