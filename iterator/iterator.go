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

type Fill[T any] interface {
	Fill(int, T)
}

type Generate[T any] interface {
	Generate(int, func() T)
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
