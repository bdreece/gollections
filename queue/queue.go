// MIT License
// Copyright (c) 2022 Brian Reece

package queue

type Queue[T any] interface {
	PushBack(T)
	PopFront() (*T, error)
}
