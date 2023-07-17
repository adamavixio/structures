package atom

import "sync"

//
// NOT DONE!!!
//
// Error Handling: Your implementation doesn't handle the case where someone tries to get a value from the map using a key that does not exist. In this case, you would be returning a zero-value for the specific type which might be valid in certain contexts and cause confusion. To prevent this, you could return an error.
// Documentation: Adding more comprehensive comments to the methods will help other developers understand what each function does. Describe what each method does, its arguments, return values, and any errors it can return.
// Testing: While not directly related to the code implementation, you should write tests to make sure your hashmap behaves as expected. This includes tests for concurrent reads and writes to make sure your locks are functioning properly.
// Performance: In Enumerate(), Keys(), Values() methods, you are using append which might cause unnecessary allocations and reallocations as the slice grows. Instead, you could initialize the slices with the length of the map using make([]K, 0, len(hashMap)) and make([]V, 0, len(hashMap)) respectively.
// Methods: Depending on your needs, you could add methods to handle multiple keys. For instance, SetMultiple, GetMultiple, or DeleteMultiple could handle bulk operations. Please note that for thread safety these should be implemented within a single lock/unlock scope.

// HashMap[K comparable, V any] represents a key-value map.
type HashMap[K comparable, V any] struct {
	data map[K]V

	mu *sync.RWMutex
}

// HashMapIter[K comparable, V any] represents an iterator over a hash map.
type HashMapIter[K comparable, V any] struct {
	Key   K
	Value V
}

// NewHashMap creates a new instance of a HashMap.
func NewHashMap[K comparable, V any]() HashMap[K, V] {
	return HashMap[K, V]{
		data: map[K]V{},
		mu:   &sync.RWMutex{},
	}
}

// Len returns the number of key-value pairs in the hash map.
func (hashMap *HashMap[K, V]) Len() int {
	hashMap.mu.RLock()
	defer hashMap.mu.RUnlock()

	return len(hashMap.data)
}

// Contains checks whether the hash map contains the specified key.
func (hashMap *HashMap[K, V]) Contains(key K) bool {
	hashMap.mu.RLock()
	defer hashMap.mu.RUnlock()

	_, ok := hashMap.data[key]
	return ok
}

// Get retrieves the value associated with the provided key in the hash map.
// The second return value indicates whether the key was found in the hash map.
func (hashMap *HashMap[K, V]) Get(key K) (V, bool) {
	hashMap.mu.RLock()
	defer hashMap.mu.RUnlock()

	value, ok := hashMap.data[key]
	return value, ok
}

// Set inserts or updates the value for the key in the hash map.
func (hashMap *HashMap[K, V]) Set(key K, value V) {
	hashMap.mu.Lock()
	defer hashMap.mu.Unlock()

	hashMap.data[key] = value
}

// Delete removes the key and its associated value from the hash map.
func (hashMap *HashMap[K, V]) Delete(key K) {
	hashMap.mu.Lock()
	defer hashMap.mu.Unlock()

	delete(hashMap.data, key)
}

// Clear removes all key-value pairs from the hash map.
func (hashMap *HashMap[K, V]) Clear() {
	hashMap.mu.Lock()
	defer hashMap.mu.Unlock()

	hashMap.data = map[K]V{}
}

// Enumerate returns two slices containing the keys and values of the hash map.
func (hashMap *HashMap[K, V]) Enumerate() ([]K, []V) {
	hashMap.mu.RLock()
	defer hashMap.mu.RUnlock()

	keys, values := []K{}, []V{}
	for key, value := range hashMap.data {
		keys = append(keys, key)
		values = append(values, value)
	}

	return keys, values
}

// Iterate creates and returns a channel that can be used to iterate
// over all key-value pairs in the hash map.
func (hashMap *HashMap[K, V]) Iterate() <-chan HashMapIter[K, V] {
	channel := make(chan HashMapIter[K, V], hashMap.Len())
	go func() {
		hashMap.mu.RLock()
		defer hashMap.mu.RUnlock()

		for key, value := range hashMap.data {
			channel <- HashMapIter[K, V]{
				Key:   key,
				Value: value,
			}
		}
		close(channel)
	}()
	return channel
}

// Keys returns a slice of all keys present in the hash map.
func (hashMap *HashMap[K, V]) Keys() []K {
	hashMap.mu.RLock()
	defer hashMap.mu.RUnlock()

	keys := make([]K, 0, hashMap.Len())
	for key := range hashMap.data {
		keys = append(keys, key)
	}
	return keys
}

// Values returns a slice of all values present in the hash map.
func (hashMap *HashMap[K, V]) Values() []V {
	hashMap.mu.RLock()
	defer hashMap.mu.RUnlock()

	values := make([]V, 0, hashMap.Len())
	for _, value := range hashMap.data {
		values = append(values, value)
	}
	return values
}

// Copy creates a new hash map and copies all key-value pairs from the current map to the new one.
func (hashMap *HashMap[K, V]) Copy() HashMap[K, V] {
	hashMap.mu.RLock()
	defer hashMap.mu.RUnlock()

	copied := NewHashMap[K, V]()
	copied.mu.Lock()
	defer copied.mu.Unlock()
	for key, value := range hashMap.data {
		copied.data[key] = value
	}

	return copied
}
