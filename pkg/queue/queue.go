package queue

import (
	"github.com/bdreece/gollections/pkg/collection"
)

// Queue provides the required methods of the
// queue data structure
type Queue[TItem any] interface {
	collection.Collection[TItem]

	// Enqueue appends an item to the end of the queue
	Enqueue(TItem)

	// Dequeue removes and returns the first item in
	// the queue
	Dequeue() *TItem
}
