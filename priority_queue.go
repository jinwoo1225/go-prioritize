package prioritize

import (
	"container/heap"
	"sync"
)

type PriorityQueuer[T any] interface {
	Pop() Item[T]
	Push(x Item[T])
	Update(item Item[T], value T, priority int64)
	Peek() Item[T]
	Range() []Item[T]
}

type PriorityQueue[T any] struct {
	mu           sync.RWMutex
	internalHeap *internalHeap[T]
}

func NewPriorityQueue[T any](items []Item[T]) *PriorityQueue[T] {
	for i := range items {
		items[i].index = i
	}

	iHeap := &internalHeap[T]{
		items: items,
	}

	heap.Init(iHeap)

	return &PriorityQueue[T]{
		mu:           sync.RWMutex{},
		internalHeap: iHeap,
	}
}

func (h *PriorityQueue[T]) Push(x Item[T]) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.internalHeap.Push(x)
	h.internalHeap.Fix(h.internalHeap.Len() - 1)
}

func (h *PriorityQueue[T]) Pop() Item[T] {
	h.mu.Lock()
	defer h.mu.Unlock()

	return h.internalHeap.Pop().(Item[T])
}

func (h *PriorityQueue[T]) Peek() Item[T] {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if h.internalHeap.Len() == 0 {
		return Item[T]{}
	}

	return h.internalHeap.items[0]
}

func (h *PriorityQueue[T]) Update(item Item[T], value T, priority int64) {
	h.mu.Lock()
	defer h.mu.Unlock()

	item.Value = value
	item.Priority = priority

	h.internalHeap.items[item.index] = item
	h.internalHeap.Fix(item.index)
}

func (h *PriorityQueue[T]) Range() []Item[T] {
	h.mu.RLock()
	defer h.mu.RUnlock()

	var items []Item[T]
	for _, item := range h.internalHeap.items {
		items = append(items, item)
	}

	return items
}
