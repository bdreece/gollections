package queue

import (
	"github.com/bdreece/gollections/pkg/collection"
)

type Queue[TItem any] interface {
	collection.Collection[TItem]
	Enqueue(TItem)
	Dequeue() *TItem
}
