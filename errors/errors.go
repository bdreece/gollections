package errors

import "fmt"

type IndexOutOfBounds struct {
	index  int
	bounds int
}

func NewIndexOutOfBounds(index, bounds int) IndexOutOfBounds {
	return IndexOutOfBounds{index, bounds}
}

func (e IndexOutOfBounds) Error() string {
	return fmt.Sprintf("index %d > bounds %d", e.index, e.bounds)
}

type Empty struct{}

func (e Empty) Error() string { return "collection is empty" }
