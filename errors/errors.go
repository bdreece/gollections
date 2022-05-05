// MIT License
// Copyright (c) 2022 Brian Reece

package errors

import "fmt"

// IndexOutOfBounds represents an error when
// index < 0 or index > bounds.
type IndexOutOfBounds struct {
	Index  int
	Bounds int
}

// Error returns a message representing the
// IndexOutOfBounds error.
func (e IndexOutOfBounds) Error() string {
	return fmt.Sprintf("index (%d) out of bounds (%d)", e.Index, e.Bounds)
}

// Empty represents an error when
// the collection is empty.
type Empty struct{}

// Error returns a message representing the
// Empty error.
func (e Empty) Error() string { return "collection is empty" }

// NotFound represents an error when
// the specified key cannot be found
type NotFound[K any] struct { Key K }

// Error returns a message representing the
// key not found error.
func (n NotFound[K]) Error() string { 
    return fmt.Sprintf("key not found: %v\n", n.Key)
}

