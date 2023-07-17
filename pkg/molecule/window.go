package molecule

import "github.com/adamavixio/structures/pkg/atom"

type Window[T any] struct {
	filter func(T) bool
	data   *atom.Ring[T]
}

func NewWindow[T any](filter func(T) bool) *Window[T] {
	return &Window[T]{
		filter: filter,
		data:   atom.NewRing[T](),
	}
}

func (window *Window[T]) Insert(value T) ([]T, []T, error) {
	front := []T{}
	for value, err := window.data.PeekFront(); window.filter(value); {
		if err != nil {
			return nil, nil, err
		}
		value, err := window.data.PopFront()
		if err != nil {
			return nil, nil, err
		}
		front = append(front, value)
	}

	back := []T{}
	for value, err := window.data.PeekBack(); window.filter(value); {
		if err != nil {
			return nil, nil, err
		}
		value, err := window.data.PopBack()
		if err != nil {
			return nil, nil, err
		}
		front = append(back, value)
	}

	return front, back, nil
}
