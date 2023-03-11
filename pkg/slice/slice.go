package slice

import (
	"errors"

	"github.com/bdreece/gollections/pkg/collection"
	"github.com/bdreece/gollections/pkg/iterator"
)

var (
	ErrIndexOutOfBounds = errors.New("index out of bounds")
)

// Slice provides a thin abstraction over the
// primitive Go slice
type Slice[TItem any] interface {
	collection.Collection[TItem]

	// First returns the first item in the Slice
	First() *TItem

	// Last returns the last item in the Slice
	Last() *TItem

	// Add appends the given item to the Slice
	Add(TItem)

	// Get returns the item at the given index
	// in the Slice, otherwise returns nil and
	// the respective error
	Get(int) (*TItem, error)

	// Set value at the given index to the given
	// item in the Slice, returning any resulting
	// error
	Set(int, TItem) error

	// Insert adds the given item at the given index
	// in the Slice, returning any resulting error
	Insert(int, TItem) error

	// Remove removes and returns the item at the given
	// index in the size, otherwise returns nil and the
	// respective error
	Remove(int) (*TItem, error)
}

type slice[TItem any] []TItem

// New creates a new Slice with the given size
// and capacity
func New[TItem any](size, capacity int) Slice[TItem] {
	var s slice[TItem] = make([]TItem, size, capacity)
	return &s
}

// From returns a Slice by casting a Go slice
func From[TItem any](s []TItem) Slice[TItem] {
	var _s slice[TItem] = s
	return &_s
}

func (s *slice[TItem]) First() *TItem {
	if len(*s) == 0 {
		return nil
	}
	val := (*s)[0]
	return &val
}

func (s *slice[TItem]) Last() *TItem {
	if len(*s) == 0 {
		return nil
	}
	index := len(*s) - 1
	val := (*s)[index]
	return &val
}

func (s *slice[TItem]) Count() int {
	return len(*s)
}

func (s *slice[TItem]) Append(item TItem) collection.Collection[TItem] {
	s.Add(item)
	return s
}

func (s *slice[TItem]) Concat(
	into iterator.IntoIterator[TItem],
) collection.Collection[TItem] {
	return s.Collect(into.Iter())
}

func (s *slice[TItem]) Collect(
	iter iterator.Iterator[TItem],
) collection.Collection[TItem] {
	iterator.ForEach(iter, func(item TItem) {
		s.Append(item)
	})
	return s
}

func (s *slice[TItem]) Add(item TItem) {
	*s = append(*s, item)
}

func (s *slice[TItem]) Get(index int) (*TItem, error) {
	if index > len(*s) || index < 0 {
		return nil, ErrIndexOutOfBounds
	}
	val := (*s)[index]
	return &val, nil
}

func (s *slice[TItem]) Set(index int, item TItem) error {
	if index > len(*s) || index < 0 {
		return ErrIndexOutOfBounds
	}
	(*s)[index] = item
	return nil
}

func (s *slice[TItem]) Insert(index int, item TItem) error {
	if index > len(*s) || index < 0 {
		return ErrIndexOutOfBounds
	}
	before := (*s)[:index]
	after := append([]TItem{item}, (*s)[index:]...)
	*s = append(before, after...)

	return nil
}

func (s *slice[TItem]) Remove(index int) (*TItem, error) {
	if index > len(*s) || index < 0 {
		return nil, ErrIndexOutOfBounds
	}

	before := (*s)[:index]
	after := (*s)[index+1:]
	item := (*s)[index]
	*s = append(before, after...)
	return &item, nil
}
