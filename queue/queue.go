// MIT License
// Copyright (c) 2022 Brian Reece

package queue

// Queue is an interface for collections
// that can represent a FIFO queue.
type Queue[T any] interface {
	PushBack(T)
	PopFront() (*T, error)
}
