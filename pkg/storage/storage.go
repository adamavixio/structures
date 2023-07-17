package storage

import (
	"fmt"

	"github.com/adamavixio/structures/pkg/atom"
	"github.com/adamavixio/structures/pkg/molecule"
)

// Make this iterable? For example, work with specific arrays or value?
// Type cast for linked list, array, etc... with switch for add?
// Add iterable interface to type cast to?

// Memory is a generic key-value storage structure where each key maps to a slice of items of type T.
type Memory[K comparable, T any] struct {
	Storage atom.HashMap[K, *molecule.Queue[T]]
}

// NewMemory initializes and returns a new Memory instance.
func NewMemory[K comparable, T any]() Memory[K, T] {
	return Memory[K, T]{
		Storage: atom.NewHashMap[K, *molecule.Queue[T]](),
	}
}

// Count returns the number of items stored for a given key.
// If the key does not exist, it returns 0.
func (memory Memory[K, T]) Count(key K) int {
	queue, ok := memory.Storage.Get(key)
	if !ok {
		return 0
	}
	return queue.Size()
}

// Contains checks whether a given key exists in the storage and returns a boolean result.
func (memory Memory[K, T]) Contains(key K) bool {
	return memory.Storage.Contains(key)
}

// Get retrieves the slice of items associated with a given key.
// It returns an error if the key does not exist in the storage.
func (memory Memory[K, T]) Get(key K) (*molecule.Queue[T], error) {
	queue, ok := memory.Storage.Get(key)
	if !ok {
		return nil, fmt.Errorf("cannot get items: storage does not have key %v", key)
	}
	return queue, nil
}

// Add appends a slice of items to the slice associated with a given key.
// If the key does not exist, it creates a new slice for that key and then appends the items.
func (memory *Memory[K, T]) Add(key K, items ...T) {
	existingItems, ok := memory.Storage.Get(key)
	if !ok {
		existingItems = molecule.NewQueue[T]()
	}
	for _, item := range items {
		existingItems.Enqueue(item)
	}
	memory.Storage.Set(key, existingItems)
}

// Remove removes a specified number of items from the slice associated with a given key.
// It returns the removed items and an error if the key does not exist or there are not enough items to remove.
func (memory *Memory[K, T]) Remove(key K, amount int) ([]T, error) {
	queue, ok := memory.Storage.Get(key)
	if !ok {
		return nil, fmt.Errorf("cannot remove items: storage does not have key %v", key)
	}

	if amount > queue.Size() {
		return nil, fmt.Errorf("cannot remove items: storage has %v items, but tried to remove %v", queue.Size(), amount)
	}

	items := []T{}
	for i := 0; i < amount; i++ {
		item, err := queue.Dequeue()
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
