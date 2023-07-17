package molecule

import (
	"testing"

	"github.com/adamavixio/structures/pkg/test"
)

func TestNewQueue(t *testing.T) {
	queue := NewQueue[int]()
	test.NotEqual(t, queue, nil)
	test.NotEqual(t, queue.data, nil)
}

func TestIsEmpty(t *testing.T) {
	queue := NewQueue[int]()
	test.Equal(t, queue.IsEmpty(), true)
}

func TestSize(t *testing.T) {
	queue := NewQueue[int]()
	test.Equal(t, queue.Size(), 0)
}

func TestPeek(t *testing.T) {
	queue := NewQueue[int]()

	value, err := queue.Peek()
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrQueueEmpty)
}

func TestEnqueue(t *testing.T) {
	queue := NewQueue[int]()

	queue.Enqueue(1)
	value, err := queue.Peek()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)

	queue.Enqueue(2)
	value, err = queue.Peek()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
}

func TestDequeue(t *testing.T) {
	queue := NewQueue[int]()

	value, err := queue.Dequeue()
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrQueueEmpty)

	queue.Enqueue(1)
	value, err = queue.Dequeue()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = queue.Dequeue()
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrQueueEmpty)

	queue.Enqueue(1)
	queue.Enqueue(2)
	value, err = queue.Dequeue()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = queue.Dequeue()
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	value, err = queue.Dequeue()
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrQueueEmpty)
}
