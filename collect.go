// MIT License
// Copyright (c) 2022 Brian Reece

package gollections

type Collect[T any] interface {
	Collect(...T)
}
