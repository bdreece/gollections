package stack

import (
	"github.com/bdreece/gollections/pkg/collection"
)

type Stack[TItem any] interface {
	collection.Collection[TItem]
	Push(TItem)
	Pop() *TItem
}
