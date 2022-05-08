package hashmap

type Iterator[K comparable, V any] struct {
    *HashMap[K, V]
}

func (i *Iterator[K, V]) Next() (*MapPair[K, V], error) {
    // TODO: figure out how to aggregate pairs from vector of maps
    return nil, nil
}
