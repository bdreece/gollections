package gollections

// Peek provides an interface for
// peekable queues and stacks.
type Peek[T any] interface {
	// Peek retrieves the next element from
	// the collection without removing it.
	// Propagates any errors from the
	// underlying data structure.
	Peek() (*T, error)
}
