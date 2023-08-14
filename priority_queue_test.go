package prioritize

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestNewHeapedQueue(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		items := []Item[string]{
			NewItem[string]("hello world", 4),
		}

		got := NewPriorityQueue[string](items)

		want := &PriorityQueue[string]{
			mu:           sync.RWMutex{},
			internalHeap: &internalHeap[string]{items: items},
		}

		assert.Equal(t, want, got)
	})
}

func TestPriorityQueue_Push(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		items := []Item[string]{
			NewItem[string]("hello world", 4),
			NewItem[string]("hello world 2", 3),
			NewItem[string]("hello world 3", 2),
		}

		pq := NewPriorityQueue[string](items)
		pq.Push(NewItem[string]("hello world 4", 1))

		got := pq.internalHeap.items

		want := []Item[string]{
			{
				Value:    "hello world",
				Priority: 4,
				index:    0,
			},
			{
				Value:    "hello world 2",
				Priority: 3,
				index:    1,
			},
			{
				Value:    "hello world 3",
				Priority: 2,
				index:    2,
			},
			{
				Value:    "hello world 4",
				Priority: 1,
				index:    3,
			},
		}

		assert.Equal(t, want, got)
	})
}

func TestPriorityQueue_Pop(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		items := []Item[string]{
			NewItem[string]("hello world", 4),
			NewItem[string]("hello world 2", 3),
			NewItem[string]("hello world 3", 2),
			NewItem[string]("hello world 4", 1),
		}

		pq := NewPriorityQueue[string](items)
		got := pq.Pop()

		want := Item[string]{
			Value:    "hello world",
			Priority: 4,
			index:    0,
		}

		assert.Equal(t, want, got)
	})
}

func TestPriorityQueue_Peek(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		items := []Item[string]{
			NewItem[string]("hello world", 4),
			NewItem[string]("hello world 2", 3),
			NewItem[string]("hello world 3", 2),
			NewItem[string]("hello world 4", 1),
		}

		pq := NewPriorityQueue[string](items)
		got := pq.Peek()

		want := Item[string]{
			Value:    "hello world",
			Priority: 4,
			index:    0,
		}

		assert.Equal(t, want, got)
	})
}

func TestPriorityQueue_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		items := []Item[string]{
			NewItem[string]("hello world", 4),
			NewItem[string]("hello world 2", 3),
			NewItem[string]("hello world 3", 2),
			NewItem[string]("hello world 4", 1),
		}

		pq := NewPriorityQueue[string](items)
		pq.Update(pq.internalHeap.items[0], "hello world updated", 5)

		got := pq.internalHeap.items

		want := []Item[string]{
			{
				Value:    "hello world updated",
				Priority: 5,
				index:    0,
			},
			{
				Value:    "hello world 2",
				Priority: 3,
				index:    1,
			},
			{
				Value:    "hello world 3",
				Priority: 2,
				index:    2,
			},
			{
				Value:    "hello world 4",
				Priority: 1,
				index:    3,
			},
		}

		assert.Equal(t, want, got)
	})
}

func TestPriorityQueue_Range(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		items := []Item[string]{
			NewItem[string]("hello world", 4),
			NewItem[string]("hello world 2", 3),
			NewItem[string]("hello world 3", 2),
			NewItem[string]("hello world 4", 1),
		}

		pq := NewPriorityQueue[string](items)
		got := pq.Range()

		want := []Item[string]{
			{
				Value:    "hello world",
				Priority: 4,
				index:    0,
			},
			{
				Value:    "hello world 2",
				Priority: 3,
				index:    1,
			},
			{
				Value:    "hello world 3",
				Priority: 2,
				index:    2,
			},
			{
				Value:    "hello world 4",
				Priority: 1,
				index:    3,
			},
		}

		assert.Equal(t, want, got)
	})
}
