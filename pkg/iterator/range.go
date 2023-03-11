package iterator

type rangeIterator struct {
	current int
	end     int
}

func (iter *rangeIterator) Next() *int {
	if iter.current >= iter.end {
		return nil
	}

	c := iter.current
	iter.current += 1
	return &c
}

func Range(start, end int) Iterator[int] {
	return &rangeIterator{start, end}
}
