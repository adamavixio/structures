package atom

import (
	"testing"

	"github.com/adamavixio/structures/pkg/test"
)

func TestNewRing(t *testing.T) {
	ring := NewRing[int]()
	test.NotEqual(t, ring, nil)
	test.Equal(t, ring.size, 0)
	test.Equal(t, ring.head, 0)
	test.Equal(t, ring.tail, 0)
	test.Equal(t, len(ring.data), 0)
	test.Equal(t, cap(ring.data), 0)
}

func TestIsEmptyRing(t *testing.T) {
	ring := NewRing[int]()
	test.Equal(t, ring.IsEmpty(), true)
}

func TestSizeRing(t *testing.T) {
	ring := NewRing[int]()
	test.Equal(t, ring.Size(), 0)
}

func TestPeekFrontRing(t *testing.T) {
	ring := NewRing[int]()

	value, err := ring.PeekFront()
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingEmpty)
}

func TestPeekBackRing(t *testing.T) {
	ring := NewRing[int]()

	value, err := ring.PeekBack()
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingEmpty)
}

func TestPeekIndexRing(t *testing.T) {
	ring := NewRing[int]()

	value, err := ring.PeekIndex(-1)
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingEmpty)
	value, err = ring.PeekIndex(0)
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingEmpty)
	value, err = ring.PeekIndex(1)
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingEmpty)

	ring.PushFront(1)
	value, err = ring.PeekIndex(-1)
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingIndex)
	value, err = ring.PeekIndex(0)
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekIndex(1)
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingIndex)
}

func TestPushFrontRing(t *testing.T) {
	ring := NewRing[int]()

	ring.PushFront(4)
	value, err := ring.PeekFront()
	test.Equal(t, value, 4)
	test.Equal(t, err, nil)
	value, err = ring.PeekBack()
	test.Equal(t, value, 4)
	test.Equal(t, err, nil)
	test.Equal(t, ring.size, 1)
	test.Equal(t, ring.head, 0)
	test.Equal(t, ring.tail, 0)
	test.Equal(t, len(ring.data), 1)
	test.Equal(t, cap(ring.data), 1)

	ring.PushFront(3)
	value, err = ring.PeekFront()
	test.Equal(t, value, 3)
	test.Equal(t, err, nil)
	value, err = ring.PeekBack()
	test.Equal(t, value, 4)
	test.Equal(t, err, nil)
	test.Equal(t, ring.size, 2)
	test.Equal(t, ring.head, 1)
	test.Equal(t, ring.tail, 1)
	test.Equal(t, len(ring.data), 2)
	test.Equal(t, cap(ring.data), 2)

	ring.PushFront(2)
	value, err = ring.PeekFront()
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	value, err = ring.PeekIndex(1)
	test.Equal(t, value, 3)
	test.Equal(t, err, nil)
	value, err = ring.PeekBack()
	test.Equal(t, value, 4)
	test.Equal(t, err, nil)
	test.Equal(t, ring.size, 3)
	test.Equal(t, ring.head, 3)
	test.Equal(t, ring.tail, 2)
	test.Equal(t, len(ring.data), 4)
	test.Equal(t, cap(ring.data), 4)

	ring.PushFront(1)
	value, err = ring.PeekFront()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekIndex(1)
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	value, err = ring.PeekIndex(2)
	test.Equal(t, value, 3)
	test.Equal(t, err, nil)
	value, err = ring.PeekBack()
	test.Equal(t, value, 4)
	test.Equal(t, err, nil)
	test.Equal(t, ring.size, 4)
	test.Equal(t, ring.head, 2)
	test.Equal(t, ring.tail, 2)
	test.Equal(t, len(ring.data), 4)
	test.Equal(t, cap(ring.data), 4)
}

func TestPushBackRing(t *testing.T) {
	ring := NewRing[int]()

	ring.PushBack(1)
	value, err := ring.PeekFront()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekBack()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	test.Equal(t, ring.size, 1)
	test.Equal(t, ring.head, 0)
	test.Equal(t, ring.tail, 0)
	test.Equal(t, len(ring.data), 1)
	test.Equal(t, cap(ring.data), 1)

	ring.PushBack(2)
	value, err = ring.PeekFront()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekBack()
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	test.Equal(t, ring.size, 2)
	test.Equal(t, ring.head, 0)
	test.Equal(t, ring.tail, 0)
	test.Equal(t, len(ring.data), 2)
	test.Equal(t, cap(ring.data), 2)

	ring.PushBack(3)
	value, err = ring.PeekFront()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekIndex(1)
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	value, err = ring.PeekBack()
	test.Equal(t, value, 3)
	test.Equal(t, err, nil)
	test.Equal(t, ring.size, 3)
	test.Equal(t, ring.head, 0)
	test.Equal(t, ring.tail, 3)
	test.Equal(t, len(ring.data), 4)
	test.Equal(t, cap(ring.data), 4)

	ring.PushBack(4)
	value, err = ring.PeekFront()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekIndex(1)
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	value, err = ring.PeekIndex(2)
	test.Equal(t, value, 3)
	test.Equal(t, err, nil)
	value, err = ring.PeekBack()
	test.Equal(t, value, 4)
	test.Equal(t, err, nil)
	test.Equal(t, ring.size, 4)
	test.Equal(t, ring.head, 0)
	test.Equal(t, ring.tail, 0)
	test.Equal(t, len(ring.data), 4)
	test.Equal(t, cap(ring.data), 4)
}

func TestPopFrontRing(t *testing.T) {
	ring := NewRing[int]()
	value, err := ring.PopFront()
	test.Equal(t, ring.size, 0)
	test.Equal(t, ring.head, 0)
	test.Equal(t, ring.tail, 0)
	test.Equal(t, len(ring.data), 0)
	test.Equal(t, cap(ring.data), 0)
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingEmpty)

	ring = NewRing[int]()
	ring.PushFront(1)
	value, err = ring.PopFront()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	test.Equal(t, ring.size, 0)
	test.Equal(t, ring.head, 0)
	test.Equal(t, ring.tail, 0)
	test.Equal(t, len(ring.data), 1)
	test.Equal(t, cap(ring.data), 1)

	ring = NewRing[int]()
	ring.PushFront(2)
	ring.PushFront(1)
	value, err = ring.PopFront()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	test.Equal(t, ring.size, 1)
	test.Equal(t, ring.head, 0)
	test.Equal(t, ring.tail, 1)
	test.Equal(t, len(ring.data), 2)
	test.Equal(t, cap(ring.data), 2)
	value, err = ring.PopFront()
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	test.Equal(t, ring.size, 0)
	test.Equal(t, ring.head, 1)
	test.Equal(t, ring.tail, 1)
	test.Equal(t, len(ring.data), 2)
	test.Equal(t, cap(ring.data), 2)

	ring = NewRing[int]()
	ring.PushBack(1)
	value, err = ring.PopFront()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	test.Equal(t, ring.size, 0)
	test.Equal(t, ring.head, 0)
	test.Equal(t, ring.tail, 0)
	test.Equal(t, len(ring.data), 1)
	test.Equal(t, cap(ring.data), 1)

	ring = NewRing[int]()
	ring.PushBack(1)
	ring.PushBack(2)
	value, err = ring.PopFront()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	test.Equal(t, ring.size, 1)
	test.Equal(t, ring.head, 1)
	test.Equal(t, ring.tail, 0)
	test.Equal(t, len(ring.data), 2)
	test.Equal(t, cap(ring.data), 2)
	value, err = ring.PopFront()
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	test.Equal(t, ring.size, 0)
	test.Equal(t, ring.head, 0)
	test.Equal(t, ring.tail, 0)
	test.Equal(t, len(ring.data), 2)
	test.Equal(t, cap(ring.data), 2)
}

func TestPopBackRing(t *testing.T) {
	ring := NewRing[int]()

	value, err := ring.PopBack()
	test.Equal(t, ring.size, 0)
	test.Equal(t, ring.head, 0)
	test.Equal(t, ring.tail, 0)
	test.Equal(t, len(ring.data), 0)
	test.Equal(t, cap(ring.data), 0)
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingEmpty)

	ring = NewRing[int]()
	ring.PushFront(1)
	value, err = ring.PopBack()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	test.Equal(t, ring.size, 0)
	test.Equal(t, ring.head, 0)
	test.Equal(t, ring.tail, 0)
	test.Equal(t, len(ring.data), 1)
	test.Equal(t, cap(ring.data), 1)

	ring = NewRing[int]()
	ring.PushFront(2)
	ring.PushFront(1)
	value, err = ring.PopBack()
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	test.Equal(t, ring.size, 1)
	test.Equal(t, ring.head, 1)
	test.Equal(t, ring.tail, 0)
	test.Equal(t, len(ring.data), 2)
	test.Equal(t, cap(ring.data), 2)
	value, err = ring.PopBack()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	test.Equal(t, ring.tail, 1)
	test.Equal(t, ring.size, 0)
	test.Equal(t, ring.head, 1)
	test.Equal(t, len(ring.data), 2)
	test.Equal(t, cap(ring.data), 2)

	ring = NewRing[int]()
	ring.PushBack(1)
	value, err = ring.PopBack()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	test.Equal(t, ring.size, 0)
	test.Equal(t, ring.head, 0)
	test.Equal(t, ring.tail, 0)
	test.Equal(t, len(ring.data), 1)
	test.Equal(t, cap(ring.data), 1)

	ring = NewRing[int]()
	ring.PushBack(1)
	ring.PushBack(2)
	value, err = ring.PopBack()
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	test.Equal(t, ring.size, 1)
	test.Equal(t, ring.head, 0)
	test.Equal(t, ring.tail, 1)
	test.Equal(t, len(ring.data), 2)
	test.Equal(t, cap(ring.data), 2)
	value, err = ring.PopBack()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	test.Equal(t, ring.size, 0)
	test.Equal(t, ring.head, 0)
	test.Equal(t, ring.tail, 0)
	test.Equal(t, len(ring.data), 2)
	test.Equal(t, cap(ring.data), 2)
}

func TestRing(t *testing.T) {
	ring := NewRing[int]()

	testRing(t, ring, func(r *Ring[int]) {
		ring.PushFront(2)
		ring.PushFront(1)
	})

	testRing(t, ring, func(r *Ring[int]) {
		ring.PushFront(1)
		ring.PushBack(2)
	})

	testRing(t, ring, func(r *Ring[int]) {
		ring.PushBack(2)
		ring.PushFront(1)
	})

	testRing(t, ring, func(r *Ring[int]) {
		ring.PushBack(1)
		ring.PushBack(2)
	})
}

func testRing(t *testing.T, ring *Ring[int], before func(*Ring[int])) {
	before(ring)

	value, err := ring.PeekFront()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekIndex(0)
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekBack()
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	value, err = ring.PeekIndex(1)
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)

	value, err = ring.PopFront()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekFront()
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	value, err = ring.PeekIndex(0)
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	value, err = ring.PeekBack()
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)

	value, err = ring.PopFront()
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	value, err = ring.PeekFront()
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingEmpty)
	value, err = ring.PeekBack()
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingEmpty)
	value, err = ring.PeekIndex(0)
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingEmpty)

	before(ring)

	value, err = ring.PeekFront()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekIndex(0)
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekBack()
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	value, err = ring.PeekIndex(1)
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)

	value, err = ring.PopFront()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekFront()
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	value, err = ring.PeekIndex(0)
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	value, err = ring.PeekBack()
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)

	value, err = ring.PopBack()
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	value, err = ring.PeekFront()
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingEmpty)
	value, err = ring.PeekBack()
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingEmpty)
	value, err = ring.PeekIndex(0)
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingEmpty)

	before(ring)

	value, err = ring.PeekFront()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekIndex(0)
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekBack()
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	value, err = ring.PeekIndex(1)
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)

	value, err = ring.PopBack()
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	value, err = ring.PeekFront()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekIndex(0)
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekBack()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)

	value, err = ring.PopFront()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekFront()
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingEmpty)
	value, err = ring.PeekBack()
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingEmpty)
	value, err = ring.PeekIndex(0)
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingEmpty)

	before(ring)

	value, err = ring.PeekFront()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekIndex(0)
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekBack()
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	value, err = ring.PeekIndex(1)
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)

	value, err = ring.PopBack()
	test.Equal(t, value, 2)
	test.Equal(t, err, nil)
	value, err = ring.PeekFront()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekIndex(0)
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekBack()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)

	value, err = ring.PopBack()
	test.Equal(t, value, 1)
	test.Equal(t, err, nil)
	value, err = ring.PeekFront()
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingEmpty)
	value, err = ring.PeekBack()
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingEmpty)
	value, err = ring.PeekIndex(0)
	test.Equal(t, value, 0)
	test.Equal(t, err, ErrRingEmpty)
}
