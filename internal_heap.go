package prioritize

import "container/heap"

type internalHeap[T any] struct {
	items []Item[T]
}

var _ heap.Interface = &internalHeap[any]{}

func (h *internalHeap[T]) Len() int {
	return len(h.items)
}

func (h *internalHeap[T]) Less(i, j int) bool {
	return h.items[i].Priority > h.items[j].Priority
}

func (h *internalHeap[T]) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
	h.items[i].index = i
	h.items[j].index = j
}

func (h *internalHeap[T]) Push(x interface{}) {
	item := x.(Item[T])
	item.index = len(h.items)
	h.items = append(h.items, item)
}

func (h *internalHeap[T]) Pop() interface{} {
	old := h.items
	n := len(old)
	item := old[0]
	h.items = old[1:n]
	return item
}

func (h *internalHeap[T]) Fix(i int) {
	heap.Fix(h, i)
}
