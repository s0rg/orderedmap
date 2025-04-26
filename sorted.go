package orderedmap

import (
	"cmp"

	"github.com/s0rg/avl"
)

// Sorted represents generic map, it will iterate keys in sorted order.
type Sorted[K cmp.Ordered, V any] struct {
	tree *avl.Tree[K, V]
}

// NewSorted creates empty sorted map.
func NewSorted[K cmp.Ordered, V any]() (rv *Sorted[K, V]) {
	return &Sorted[K, V]{}
}

// Set adds new element to map.
func (s *Sorted[K, V]) Set(key K, value V) {
	if s.tree == nil {
		s.tree = avl.New[K, V]()
	}

	s.tree.Add(key, value)
}

// Del removes element from map.
func (s *Sorted[K, V]) Del(key K) {
	if s.tree == nil {
		return
	}

	s.tree.Del(key)
}

// Get returns element from map (if any).
func (s *Sorted[K, V]) Get(key K) (value V, ok bool) {
	if s.tree == nil {
		return
	}

	value, ok = s.tree.Get(key)

	return
}

// Len returns number of elements in map.
func (s *Sorted[K, V]) Len() (rv int) {
	if s.tree == nil {
		return
	}

	return s.tree.Len()
}

// Iter iterates elements in sorted order.
func (s *Sorted[K, V]) Iter(next func(key K, value V) bool) {
	if s.tree == nil {
		return
	}

	s.tree.Iter(next)
}

// Has checks key existence.
func (s *Sorted[K, V]) Has(key K) (yes bool) {
	_, yes = s.Get(key)

	return
}

// Clone creates shallow copy of map.
func (s *Sorted[K, V]) Clone() (rv *Sorted[K, V]) {
	rv = NewSorted[K, V]()

	if s.tree == nil {
		return
	}

	s.tree.Iter(func(key K, value V) bool {
		rv.Set(key, value)

		return true
	})

	return
}
