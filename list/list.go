package list

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

type List[T any] struct {
	head   *node[T]
	length int
}

func New[T any]() *List[T] {
	return &List[T]{head: nil, length: 0}
}

func (l *List[T]) Push(val T) {
	l.length += 1
	if l.head == nil {
		l.head = new(node[T])
		l.head.value = val
		return
	}

	walk := l.head
	for walk.next != nil {
		walk = walk.next
	}
	walk.next = new(node[T])
	walk.next.value = val
}

func (l *List[T]) Pop() *T {
	if l.head == nil {
		return nil
	}

	value := new(T)
	*value = l.head.value
	l.head = l.head.next
	l.length -= 1
	return value
}
