package queue

import (
	"sync"
)

// Queue is a thread-safe generic queue implementation
type Queue[T any] struct {
	items []T
	mutex sync.RWMutex
}

// NewQueue creates a new queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0),
	}
}

// Enqueue adds an item to the queue
func (q *Queue[T]) Enqueue(item T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.items = append(q.items, item)
}

// Dequeue removes and returns an item from the queue
func (q *Queue[T]) Dequeue() (T, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	var zero T
	if len(q.items) == 0 {
		return zero, false
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item, true
}

// Peek returns the next item without removing it
func (q *Queue[T]) Peek() (T, bool) {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	var zero T
	if len(q.items) == 0 {
		return zero, false
	}

	return q.items[0], true
}

// IsEmpty checks if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	return len(q.items) == 0
}

// Size returns the number of items in the queue
func (q *Queue[T]) Size() int {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	return len(q.items)
}

// Clear removes all items from the queue
func (q *Queue[T]) Clear() {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.items = make([]T, 0)
}

// Items returns a copy of all items in the queue
func (q *Queue[T]) Items() []T {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// Create a copy to avoid race conditions
	items := make([]T, len(q.items))
	copy(items, q.items)

	return items
}
