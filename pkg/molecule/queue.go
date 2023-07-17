package molecule

import (
	"errors"

	"github.com/adamavixio/structures/pkg/atom"
)

var (
	ErrQueueEmpty = errors.New("queue is empty")
)

type Queue[T any] struct {
	data *atom.Ring[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		data: atom.NewRing[T](),
	}
}

func (queue *Queue[T]) IsEmpty() bool {
	return queue.data.Size() == 0
}

func (queue *Queue[T]) Size() int {
	return queue.data.Size()
}

func (queue *Queue[T]) Peek() (T, error) {
	value, err := queue.data.PeekFront()
	if err == atom.ErrRingEmpty {
		return value, ErrQueueEmpty
	}

	return value, nil
}

func (queue *Queue[T]) Enqueue(value T) {
	queue.data.PushBack(value)
}

func (queue *Queue[T]) Dequeue() (T, error) {
	value, err := queue.data.PopFront()
	if err == atom.ErrRingEmpty {
		var empty T
		return empty, ErrQueueEmpty
	}

	return value, nil
}
