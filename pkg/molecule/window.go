package molecule

import (
	"errors"

	"github.com/adamavixio/structures/pkg/atom"
)

var (
	ErrWindowEmpty = errors.New("queue is empty")
)

type Window[T any] struct {
	data *atom.Ring[T]
}

func NewWindow[T any]() *Window[T] {
	return &Window[T]{
		data: atom.NewRing[T](),
	}
}

func (window *Window[T]) PeekStart() (T, error) {
	value, err := window.data.PeekFront()
	if err == atom.ErrRingEmpty {
		return value, ErrWindowEmpty
	}

	return value, nil
}

func (window *Window[T]) InsertStart(value T) {
	window.data.PushFront(value)
}

func (window *Window[T]) FilterStart(filter func(T) bool) ([]T, error) {
	values := []T{}
	for {
		value, err := window.data.PeekFront()
		if err != nil {
			return nil, err
		}
		if !filter(value) {
			break
		}
		_, err = window.data.PopFront()
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, nil
}

func (window *Window[T]) PeekEnd() (T, error) {
	value, err := window.data.PeekBack()
	if err == atom.ErrRingEmpty {
		return value, ErrWindowEmpty
	}

	return value, nil
}

func (window *Window[T]) InsertEnd(value T) {
	window.data.PushBack(value)
}

func (window *Window[T]) FilterEnd(filter func(T) bool) ([]T, error) {
	values := []T{}
	for {
		value, err := window.data.PeekBack()
		if err != nil {
			return nil, err
		}
		if !filter(value) {
			break
		}
		_, err = window.data.PopBack()
		if err != nil {
			return nil, err
		}
		values = append([]T{value}, values...)
	}
	return values, nil
}
