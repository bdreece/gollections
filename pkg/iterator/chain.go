package iterator

type chainIterator[TItem any] struct {
	first    Iterator[TItem]
	second   Iterator[TItem]
	hasFirst bool
}

func Chain[TItem any](first, second Iterator[TItem]) Iterator[TItem] {
	return &chainIterator[TItem]{first, second, true}
}

func (c *chainIterator[TItem]) Next() *TItem {
	var val *TItem
	if c.hasFirst {
		val = c.first.Next()
		if val == nil {
			c.hasFirst = false
			val = c.second.Next()
		}
	} else {
		val = c.second.Next()
	}
	return val
}
