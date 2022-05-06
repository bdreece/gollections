// MIT License
// Copyright (c) 2022 Brian Reece

package collection

// Indexed provides an interface for
// collections whose elements can be
// accessed at random.
type Indexed[T, U any] interface {
	Collection[T]
	Set(U, T) error
	Ins(U, T) error
	Get(U) (*T, error)
	Del(U) (*T, error)
}
