package vector

import "fmt"

type IndexOutOfBoundsError struct {
	index  int
	bounds int
}

func (e *IndexOutOfBoundsError) Error() string {
	return fmt.Sprintf("index %d > bounds %d", e.index, e.bounds)
}

type Vector[T any] []T

func New[T any]() *Vector[T] {
	return new(Vector[T])
}

func (v *Vector[T]) Push(value T) {
	*v = append([]T(*v), value)
}

func (v *Vector[T]) Pop() *T {
	n := len(*v)
	if n == 0 {
		return nil
	}
	val := new(T)
	if n > 1 {
		*val = []T(*v)[n-1]
		*v = []T(*v)[:n-1]
	} else {
		*val = []T(*v)[0]
		v = New[T]()
	}

	return val
}

func (v Vector[T]) Get(i int) (*T, error) {
	n := len(v)
	if i > n {
		return nil, &IndexOutOfBoundsError{
			index:  i,
			bounds: n,
		}
	}
	val := new(T)
	*val = []T(v)[i]
	return val, nil
}

func (v *Vector[T]) Set(value T, i int) error {
	n := len(*v)
	if i >= n {
		return &IndexOutOfBoundsError{
			index:  i,
			bounds: n,
		}
	}
	[]T(*v)[i] = value
	return nil
}

func (v *Vector[T]) Insert(value T, i int) error {
	n := len(*v)
	if i >= n {
		return &IndexOutOfBoundsError{
			index:  i,
			bounds: n,
		}
	}
	before := []T(*v)[:i]
	after := []T(*v)[i+1:]
	*v = append(before, value)
	*v = append(*v, after...)
	return nil
}
