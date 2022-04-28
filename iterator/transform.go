package iterator

import "github.com/bdreece/gollections/list"

type MapPredicate[T, U any] func(T) U

func Map[T, U any](iter Iterator[T], fn MapPredicate[T, U]) (Iterator[U], error) {
	l := list.New[U]()
	for {
		item, err := iter.Next()
		if item == nil {
			break
		}
		if err != nil {
			return l.Iterator(), err
		}
		l.PushBack((fn)(*item))
	}
	return l.Iterator(), nil
}
