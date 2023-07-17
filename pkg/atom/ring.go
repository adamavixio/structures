package atom

import (
	"errors"
	"sync"
)

var (
	ErrRingEmpty = errors.New("ring is empty")
	ErrRingIndex = errors.New("index out of range")
)

type Ring[T any] struct {
	size int
	head int
	tail int
	data []T

	mu *sync.RWMutex
}

func NewRing[T any]() *Ring[T] {
	return &Ring[T]{
		data: []T{},
		mu:   &sync.RWMutex{},
	}
}

func (ring *Ring[T]) IsEmpty() bool {
	ring.mu.RLock()
	defer ring.mu.RUnlock()

	return ring.size == 0
}

func (ring *Ring[T]) Size() int {
	ring.mu.RLock()
	defer ring.mu.RUnlock()

	return ring.size
}

func (ring *Ring[T]) PeekFront() (T, error) {
	ring.mu.RLock()
	defer ring.mu.RUnlock()

	if ring.size == 0 {
		var empty T
		return empty, ErrRingEmpty
	}
	return ring.data[ring.head], nil
}

func (ring *Ring[T]) PeekBack() (T, error) {
	ring.mu.RLock()
	defer ring.mu.RUnlock()

	if ring.size == 0 {
		var empty T
		return empty, ErrRingEmpty
	}
	tailIndex := (ring.tail - 1 + len(ring.data)) % len(ring.data)
	return ring.data[tailIndex], nil
}

func (ring *Ring[T]) PeekIndex(index int) (T, error) {
	ring.mu.RLock()
	defer ring.mu.RUnlock()

	if ring.size == 0 {
		var empty T
		return empty, ErrRingEmpty
	}
	if index < 0 || index >= ring.size {
		var empty T
		return empty, ErrRingIndex
	}
	return ring.data[(ring.head+index)%len(ring.data)], nil
}

func (ring *Ring[T]) PushFront(value T) {
	ring.mu.Lock()
	defer ring.mu.Unlock()

	if ring.size == len(ring.data) {
		ring.resize()
	}
	ring.head = (ring.head - 1 + len(ring.data)) % len(ring.data)
	ring.data[ring.head] = value
	ring.size++
}

func (ring *Ring[T]) PushBack(value T) {
	ring.mu.Lock()
	defer ring.mu.Unlock()

	if ring.size == len(ring.data) {
		ring.resize()
	}
	ring.data[ring.tail] = value
	ring.tail = (ring.tail + 1) % len(ring.data)
	ring.size++
}

func (ring *Ring[T]) PopFront() (T, error) {
	ring.mu.Lock()
	defer ring.mu.Unlock()

	if ring.size == 0 {
		var empty T
		return empty, ErrRingEmpty
	}
	value := ring.data[ring.head]
	ring.head = (ring.head + 1) % len(ring.data)
	ring.size--
	return value, nil
}

func (ring *Ring[T]) PopBack() (T, error) {
	ring.mu.Lock()
	defer ring.mu.Unlock()

	if ring.size == 0 {
		var empty T
		return empty, ErrRingEmpty
	}
	ring.tail = (ring.tail - 1 + len(ring.data)) % len(ring.data)
	value := ring.data[ring.tail]
	ring.size--
	return value, nil
}

func (ring *Ring[T]) resize() {
	size := len(ring.data)
	if size == 0 {
		size = 1
	} else {
		size *= 2
	}

	newData := make([]T, size)
	n := copy(newData, ring.data[ring.head:])
	copy(newData[n:], ring.data[:ring.tail])

	ring.head = 0
	ring.tail = ring.size
	ring.data = newData
}
