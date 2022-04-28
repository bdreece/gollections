package iter

type Collection[T any] interface {
	Iterator() Iterator[T]
	Collect(...T)
}

func Collect[T any, U any](iter Iterator[T], coll Collection[T]) (*Collection[T], error) {
	var (
		err error
		buf []T
	)
	for {
		item, err := iter.Next()
		if err != nil || item == nil {
			break
		}
		buf = append(buf, *item)
	}
	coll.Collect(buf...)
	return &coll, err
}
