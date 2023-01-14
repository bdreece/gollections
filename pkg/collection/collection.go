package collection

import "github.com/bdreece/gollections/pkg/iterator"

type Collection[TItem any] interface {
	iterator.IntoIterator[TItem]
	Concat(iterator.IntoIterator[TItem]) Collection[TItem]
	Collect(iterator.Iterator[TItem]) Collection[TItem]
	Append(TItem) Collection[TItem]
	Count() int
}
