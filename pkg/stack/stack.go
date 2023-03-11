package stack

import (
	"github.com/bdreece/gollections/pkg/collection"
)

// Stack provides the required methods of the stack
// data structure
type Stack[TItem any] interface {
	collection.Collection[TItem]

	// Push prepends the Stack with the given item
	Push(TItem)

	// Pop removes and returns the first item in the Stack
	Pop() *TItem

	// Peek returns the first item in the Stack without
	// removing it
	Peek() *TItem
}
