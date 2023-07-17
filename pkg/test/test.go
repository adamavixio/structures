package test

import "testing"

func NotEqual[T comparable](t *testing.T, first, second T) {
	if first == second {
		t.Errorf("got %v, expected not equal %v", first, second)
	}
}

func Equal[T comparable](t *testing.T, first, second T) {
	if first != second {
		t.Errorf("got %v, expected %v", first, second)
	}
}
