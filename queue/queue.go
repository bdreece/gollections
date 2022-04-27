package queue

type Queue[T any] []T

func New[T any]() *Queue[T] {
	return new(Queue[T])
}

func (q *Queue[T]) Enqueue(value T) {
	*q = append(*q, value)
}

func (q *Queue[T]) Dequeue() *T {
	n := len(*q)
	if n == 0 {
		return nil
	}

	value := new(T)
	*value = []T(*q)[0]
	*q = Queue[T]([]T(*q)[1:])
	return value
}

func (q Queue[T]) Peek() *T {
	n := len(q)
	if n == 0 {
		return nil
	}

	value := new(T)
	*value = []T(q)[0]
	return value
}
