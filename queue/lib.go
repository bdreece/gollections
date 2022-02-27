package queue

type Queue []interface{}

func (q *Queue) Enqueue(value []interface{}) {
	*q = append(*q, value)
}

func (q *Queue) Dequeue() interface{} {
  value := []interface{}(*q)[0]
  *q = Queue([]interface{}(*q)[1:])
  return value
}

func (q Queue) Peek() interface{} {
	return []interface{}(q)[0]
}

