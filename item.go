package prioritize

// Item for generic support
type Item[T any] struct {
	// The Value of the item; arbitrary.
	Value T
	// The Priority of the item in the queue.
	Priority int64
	// The index of the item in the heap.
	// This is needed by update and is maintained by the heap.Interface methods.
	index int
}

func NewItem[T any](value T, priority int64) Item[T] {
	return Item[T]{
		Value:    value,
		Priority: priority,
		index:    -1,
	}
}
