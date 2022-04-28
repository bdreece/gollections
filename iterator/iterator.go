// MIT License
// Copyright (c) 2022 Brian Reece

package iterator

type Iterator[T any] interface {
	Next() (*T, error)
}

func Any[T any](iter Iterator[T], fn func(*T) bool) (bool, error) {
	for {
		item, err := iter.Next()
		if item == nil {
			break
		}
		if err != nil {
			return false, err
		}
		if (fn)(item) {
			return true, nil
		}
	}
	return false, nil
}

func All[T any](iter Iterator[T], fn func(*T) bool) (bool, error) {
	for {
		item, err := iter.Next()
		if item == nil {
			break
		}
		if err != nil {
			return true, err
		}
		if !(fn)(item) {
			return false, nil
		}
	}
	return true, nil
}

func None[T any](iter Iterator[T], fn func(*T) bool) (bool, error) {
	for {
		item, err := iter.Next()
		if item == nil {
			break
		}
		if err != nil {
			return true, err
		}
		if (fn)(item) {
			return false, err
		}
	}
	return true, nil
}

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

func Reduce[T any](iter Iterator[T], fn func(*T, T) *T) (*T, error) {
	total := new(T)
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

func Fold[T any, U any](iter Iterator[T], init *U, fn func(*U, T) *U) (*U, error) {
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
