// MIT License
// Copyright (c) 2022 Brian Reece

package gollections

// Queue provides an interface for
// queue-like (i.e. FIFO) collections.
type Queue[T any] interface {
	// Peek provides the Peek functionality
	// to the Queue.
	Peek[T]

	// Enqueue appends an element to the
	// end of the queue. That is, in such
	// a way to ensure it is dequeued after
	// all present elements in the queue;
	// reference concrete implementations
	// for details.
	Enqueue(T) error

	// Dequeue removes and returns the
	// element at the beginning of the
	// queue. That is, in such a way to
	// ensure that the oldest enqueued
	// element is dequeued; reference
	// concrete implementations for details.
	Dequeue() (*T, error)
}
