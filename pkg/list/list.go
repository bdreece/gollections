package list

import (
	"github.com/bdreece/gollections/pkg/collection"
)

type List[TItem any] interface {
	collection.Collection[TItem]
	First() *TItem
	Last() *TItem
	Add(TItem)
	Get(int) (*TItem, error)
	Set(int, TItem) error
	Remove(int) (*TItem, error)
}
