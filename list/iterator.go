package list

import "github.com/bdreece/gollections/errors"

// Iterator provides an iterator over the List
// type.
type Iterator[T any] struct {
	*node[T]
}

// Next returns the next item in the iterator
// over a list. Returns nil, errors.Empty
// at the end of the iterator.
func (i *Iterator[T]) Next() (*T, error) {
	if i.node == nil {
		return nil, errors.Empty{}
	}
	ret := i.node
	i.node = i.node.next
	return &ret.value, nil
}
