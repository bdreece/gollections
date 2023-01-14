package dll

import "github.com/bdreece/gollections/pkg/iterator"

type dllIterator[TItem any] struct {
	current *node[TItem]
}

func (d *dll[TItem]) Iter() iterator.Iterator[TItem] {
	return &dllIterator[TItem]{d.first}
}

func (d *dllIterator[TItem]) Next() *TItem {
	if d.current == nil {
		return nil
	}
	val := d.current.item
	d.current = d.current.next
	return &val
}
