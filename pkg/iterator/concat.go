package iterator

type concatIterator[TItem any] struct {
	first    Iterator[TItem]
	second   Iterator[TItem]
	hasFirst bool
}

// Concat concatenates the second iterator onto the end
// of the first iterator
func Concat[TItem any](first, second Iterator[TItem]) Iterator[TItem] {
	return &concatIterator[TItem]{first, second, true}
}

func (c *concatIterator[TItem]) Next() *TItem {
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
