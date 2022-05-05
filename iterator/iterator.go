// MIT License
// Copyright (c) 2022 Brian Reece

package iterator

// Iterator is an iterface for iterating
// over collections
type Iterator[T any] interface {
	Next() (*T, error)
}

type Reverse[T any] interface {
	Prev() (*T, error)
}

// Any returns true if for any item in the iterator,
// pred returns true. Returns false, error upon collection
// error.
func Any[T any](iter Iterator[T], pred func(*T) bool) (bool, error) {
	for {
		item, err := iter.Next()
		if item == nil {
			break
		}
		if err != nil {
			return false, err
		}
		if (pred)(item) {
			return true, nil
		}
	}
	return false, nil
}

// All returns true if for all items in the iterator,
// pred returns true. Returns false, error upon collection
// error.
func All[T any](iter Iterator[T], pred func(*T) bool) (bool, error) {
	for {
		item, err := iter.Next()
		if item == nil {
			break
		}
		if err != nil {
			return true, err
		}
		if !(pred)(item) {
			return false, nil
		}
	}
	return true, nil
}

// None returns true if for all items in the iterator,
// pred returns false. Returns false, error upon collection
// error.
func None[T any](iter Iterator[T], pred func(*T) bool) (bool, error) {
	for {
		item, err := iter.Next()
		if item == nil {
			break
		}
		if err != nil {
			return true, err
		}
		if (pred)(item) {
			return false, err
		}
	}
	return true, nil
}

// ForEach calls fn with each item in the iterator.
// Returns error on collection error
func ForEach[T any](iter Iterator[T], fn func(*T)) error {
	for {
		item, err := iter.Next()
		if item == nil {
			break
		}
		if err != nil {
			return err
		}
		(fn)(item)
	}

	return nil
}

// Reduce computes a value from all the items in the iterator.
// fn takes two arguments: an accumulator, and an item from
// the iterator. After all items have been processed, the
// accumulator is returned. Returns nil, error on collection
// error.
func Reduce[T, U any](iter Iterator[T], fn func(*U, T)) (*U, error) {
	total := new(U)
	for {
		item, err := iter.Next()
		if item == nil {
			break
		}
		if err != nil {
			return total, err
		}
		(fn)(total, *item)
	}
	return total, nil
}

// Fold functions similarly to Reduce, except that an
// initial value is provided.
func Fold[T any, U any](iter Iterator[T], init *U, fn func(*U, T)) (*U, error) {
	for {
		item, err := iter.Next()
		if item == nil {
			break
		}
		if err != nil {
			return init, err
		}
		(fn)(init, *item)
	}
	return init, nil
}

// Map transforms an iterator over one type into an iterator
// over another type, converting each item via pred. Returns
// Iterator[U], error on collection error.
func Map[T, U any](iter Iterator[T], coll Collection[U], pred func(*T) U) (Iterator[U], error) {
	for {
		item, err := iter.Next()
		if item == nil {
			break
		}
		if err != nil {
			return coll.Iterator(), err
		}
		coll.Append((pred)(item))
	}
	return coll.Iterator(), nil
}
