package dll

import (
	"github.com/bdreece/gollections/pkg/collection"
	"github.com/bdreece/gollections/pkg/iterator"
	"github.com/bdreece/gollections/pkg/queue"
	"github.com/bdreece/gollections/pkg/stack"
)

// DLL provides the implementation of a doubly-linked list
// of nodes
type DLL[TItem any] interface {
	stack.Stack[TItem]
	queue.Queue[TItem]
}

type node[TItem any] struct {
	item TItem
	prev *node[TItem]
	next *node[TItem]
}

type dll[TItem any] struct {
	first *node[TItem]
	last  *node[TItem]
}

// New creates an empty DLL
func New[TItem any]() DLL[TItem] {
	return &dll[TItem]{nil, nil}
}

// From creates a new DLL by concatenating the items
// of the given collection
func From[TItem any](c collection.Collection[TItem]) DLL[TItem] {
	d := New[TItem]()
	d.Concat(c)
	return d
}

func (d *dll[TItem]) Count() int {
	i := 0
	current := d.first
	for current != nil {
		current = current.next
		i++
	}

	return i
}

func (d *dll[TItem]) Concat(into iterator.IntoIterator[TItem]) collection.Collection[TItem] {
	return d.Collect(into.Iter())
}

func (d *dll[TItem]) Collect(iter iterator.Iterator[TItem]) collection.Collection[TItem] {
	iterator.ForEach(iter, func(item TItem) {
		d.Append(item)
	})
	return d
}

func (d *dll[TItem]) Append(item TItem) collection.Collection[TItem] {
	if d.first == nil {
		d.first = &node[TItem]{item, nil, nil}
		d.last = d.first
		return d
	}

	d.last.next = &node[TItem]{item, d.last, nil}
	d.last = d.last.next
	return d
}

func (d *dll[TItem]) Push(item TItem) {
	if d.first == nil {
		d.first = &node[TItem]{item, nil, nil}
		d.last = d.first
		return
	}

	d.first.prev = &node[TItem]{item, nil, d.first}
	d.first = d.first.prev
}

func (d *dll[TItem]) Pop() *TItem {
	if d.first == nil {
		return nil
	}

	val := d.first.item
	d.first = d.first.next
	return &val
}

func (d *dll[TItem]) Peek() *TItem {
	if d.first == nil {
		return nil
	}

	return &d.first.item
}

func (d *dll[TItem]) Enqueue(item TItem) {
	d.Append(item)
}

func (d *dll[TItem]) Dequeue() *TItem {
	return d.Pop()
}
