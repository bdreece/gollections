// MIT License
// Copyright (c) 2022 Brian Reece

package gollections

// Array provides an interface for
// array-like data structures (i.e.
// random element access).
type Array[I, T any] interface {
    // Get retrieves the element in
    // the collection associated with
    // the provided index, or nil and
    // the relevant error. Note that
    // the value returned is not
    // necessarily a pointer into the
    // collection; reference the concrete
    // implementations for details.
    Get(I) (*T, error)

    // Del removes the element in
    // the collection associated with
    // the provided index, returning
    // either the removed element, or
    // nil and the relevant error. Note
    // that the value returned is not
    // a pointer into the collection, as
    // the element will have been removed.
    Del(I) (*T, error)
    
    // Set overwrites the element in
    // the collection associated with
    // the provided index, propagating
    // internal errors.
    Set(I, T) error

    // Ins inserts an element into
    // the collection at the provided
    // index. Note that this method may
    // invalidate indices of other elements
    // in the collection; reference the
    // concrete implementations for details.
    Ins(I, T) error
}
