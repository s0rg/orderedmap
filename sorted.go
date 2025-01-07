package orderedmap

import (
	"cmp"
	"maps"
	"slices"
)

// Sorted represents generic map, it will iterate keys in sorted order.
type Sorted[K cmp.Ordered, V any] struct {
	m map[K]V
	o []K
}

// NewSorted creates empty sorted map.
func NewSorted[K cmp.Ordered, V any]() (rv *Sorted[K, V]) {
	return &Sorted[K, V]{}
}

// Set adds new element to map.
func (s *Sorted[K, V]) Set(key K, value V) {
	if s.m == nil {
		s.m, s.o = make(map[K]V), make([]K, 0, 1)
	}

	_, exists := s.m[key]
	s.m[key] = value

	if !exists {
		s.o = append(s.o, key)
		slices.Sort(s.o)
	}
}

// Del removes element from map.
func (s *Sorted[K, V]) Del(key K) {
	if s.m == nil {
		return
	}

	_, exists := s.m[key]
	if !exists {
		return
	}

	delete(s.m, key)

	if n, found := slices.BinarySearch(s.o, key); found {
		s.o = slices.Delete(s.o, n, n+1)
	}
}

// Get returns element from map (if any).
func (s *Sorted[K, V]) Get(key K) (value V, ok bool) {
	if s.m == nil {
		return
	}

	value, ok = s.m[key]

	return
}

// Len returns number of elements in map.
func (s *Sorted[K, V]) Len() int {
	return len(s.o)
}

// Iter iterates elements in sorted order.
func (s *Sorted[K, V]) Iter(next func(key K, value V) bool) {
	for _, key := range s.o {
		value := s.m[key]

		if !next(key, value) {
			break
		}
	}
}

// Has checks key existence.
func (s *Sorted[K, V]) Has(key K) (yes bool) {
	_, yes = s.Get(key)

	return
}

// Clone creates shallow copy of map.
func (s *Sorted[K, V]) Clone() (rv *Sorted[K, V]) {
	return &Sorted[K, V]{
		m: maps.Clone(s.m),
		o: slices.Clone(s.o),
	}
}
