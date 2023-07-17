package atom

// Set is a custom data structure that represents a collection of unique elements.
type Set[K comparable, V any] map[K]V

// HashMapIter[K comparable, V any] represents an iterator over a hash map.
type SetIter[K comparable, V any] struct {
	key   K
	value V
}

// NewSet creates and returns a new, empty Set.
func NewSet[K comparable, V any]() Set[K, V] {
	return Set[K, V]{}
}

// Size returns the number of elements in the set.
func (set Set[K, V]) Size() int {
	return len(set)
}

// Add adds a new element to the set.
// If the element already exists, it has no effect.
func (set Set[K, V]) Add(key K, value V) {
	set[key] = value
}

// Remove removes an element from the set.
// If the element doesn't exist, it has no effect.
func (set Set[K, V]) Remove(key K) {
	delete(set, key)
}

// Contains checks whether a given element exists in the set.
func (set Set[K, V]) Contains(key K) bool {
	_, ok := set[key]
	return ok
}

// Clear removes all elements from the set.
func (set Set[K, V]) Clear() {
	for key := range set {
		delete(set, key)
	}
}

// Enumerate returns two slices containing the keys and values of the set.
func (set Set[K, V]) Enumerate() ([]K, []V) {
	keys, values := []K{}, []V{}
	for key, value := range set {
		keys = append(keys, key)
		values = append(values, value)
	}

	return keys, values
}

// Iterate creates and returns a channel that can be used to iterate over all key-value pairs in the set.
func (set Set[K, V]) Iterate() <-chan SetIter[K, V] {
	channel := make(chan SetIter[K, V], set.Size())
	go func() {
		for key, value := range set {
			channel <- SetIter[K, V]{
				key:   key,
				value: value,
			}
		}
		close(channel)
	}()
	return channel
}

// Union returns a new set that contains all the elements
// that are in either the current set or another set.
func (set Set[K, V]) Union(other Set[K, V]) Set[K, V] {
	unionSet := NewSet[K, V]()

	for key, value := range set {
		unionSet.Add(key, value)
	}

	for key, value := range other {
		if !unionSet.Contains(key) {
			unionSet.Add(key, value)
		}
	}

	return unionSet
}

// Intersection returns a new set that contains all the elements
// that are in both the current set and another set.
func (set Set[K, V]) Intersection(other Set[K, V]) Set[K, V] {
	intersectionSet := NewSet[K, V]()

	for key, value := range set {
		if other.Contains(key) {
			intersectionSet.Add(key, value)
		}
	}

	return intersectionSet
}

// Difference returns a new set that contains all the elements
// that are in the current set but not in another set.
func (set Set[K, V]) Difference(other Set[K, V]) Set[K, V] {
	differenceSet := NewSet[K, V]()

	for key, value := range set {
		if !other.Contains(key) {
			differenceSet.Add(key, value)
		}
	}

	return differenceSet
}

// SymmetricDifference returns a new set that contains all the elements
// that are unique in each set (i.e., the elements that are in either
// the current set or another set but not in both).
func (set Set[K, V]) SymmetricDifference(other Set[K, V]) Set[K, V] {
	diffSet := NewSet[K, V]()

	for key, value := range set {
		if !other.Contains(key) {
			diffSet.Add(key, value)
		}
	}

	for key, value := range other {
		if !set.Contains(key) {
			diffSet.Add(key, value)
		}
	}

	return diffSet
}

// Subset checks if all elements in the current set are also in another set.
func (set Set[K, V]) Subset(other Set[K, V]) bool {
	for key := range set {
		if !other.Contains(key) {
			return false
		}
	}

	return true
}
