package iter

type Iterator[T any] interface {
	Next() (*T, error)
}

func ForEach[T any](iter Iterator[T], fn func(*T)) error {
	for {
		item, err := iter.Next()
		if err != nil {
			return err
		}
		if item == nil {
			break
		}
		(fn)(item)
	}

	return nil
}

func Reduce[T any](iter Iterator[T], fn func(*T, T) *T) (*T, error) {
	total := new(T)
	for {
		item, err := iter.Next()
		if err != nil {
			return total, err
		}
		if item == nil {
			break
		}
		(fn)(total, *item)
	}
	return total, nil
}

func Fold[T any, U any](iter Iterator[T], init *U, fn func(*U, T) *U) (*U, error) {
	for {
		item, err := iter.Next()
		if err != nil {
			return init, err
		}
		if item == nil {
			break
		}
		(fn)(init, *item)
	}
	return init, nil
}
