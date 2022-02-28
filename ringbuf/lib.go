package ringbuf

type RingBuf struct {
  data   []interface{}
  size   int
  length int
  head   int
  tail   int
}

func New(size int) RingBuf {
  return RingBuf{
    data: make([]interface{}, size),
    size: size,
    length: 0,
    head: 0,
    tail: 0,
  }
}

func (b *RingBuf) Read() interface{} {
  if b.length <= 0 {
    return nil
  }

  val := b.data[b.head]
  
  b.head += 1
  if b.head >= b.size {
    b.head = 0
  }

  b.length -= 1
  return val
}

func (b RingBuf) Peek() interface{} {
  if b.length <= 0 {
    return nil
  }
  return b.data[b.head]
}

func (b *RingBuf) Write(val interface{}) {
  b.data[b.tail] = val
  
  b.tail += 1
  if b.tail >= b.size {
    b.tail = 0
  }
  
  b.length += 1
}

func (b *RingBuf) Clear() {
  b.data = make([]interface{}, b.size)
}
