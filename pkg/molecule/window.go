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

func (window *Window[T]) Insert(value T) error {
	for value, err := window.data.PeekFront(); window.filter(value); {
		if err != nil {
			return err
		}
		window.data.PopFront()
	}
	for value, err := window.data.PeekBack(); window.filter(value); {
		if err != nil {
			return err
		}
		window.data.PopBack()
	}
	return nil
}
