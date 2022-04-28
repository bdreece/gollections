package errors

import "fmt"

type IndexOutOfBoundsError struct {
	index  int
	bounds int
}

func NewIndexOutOfBoundsError(index, bounds int) IndexOutOfBoundsError {
	return IndexOutOfBoundsError{index, bounds}
}

func (e *IndexOutOfBoundsError) Error() string {
	return fmt.Sprintf("index %d > bounds %d", e.index, e.bounds)
}
