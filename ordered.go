package orderedmap

import (
	"maps"
	"slices"
)

// Ordered represents generic map, it will iterate keys always in same order - as elements were added (FIFO).
type Ordered[K comparable, V any] struct {
	m map[K]V
	o []K
}

// NewOrdered creates empty ordered map.
func NewOrdered[K comparable, V any]() (rv *Ordered[K, V]) {
	return &Ordered[K, V]{}
}

// Set adds new element to map.
func (o *Ordered[K, V]) Set(key K, value V) {
	if o.m == nil {
		o.m, o.o = make(map[K]V), make([]K, 0, 1)
	}

	_, exists := o.m[key]
	o.m[key] = value

	if !exists {
		o.o = append(o.o, key)
	}
}

// Del removes element from map.
func (o *Ordered[K, V]) Del(key K) {
	if o.m == nil {
		return
	}

	_, exists := o.m[key]
	if !exists {
		return
	}

	delete(o.m, key)

	if n := slices.Index(o.o, key); n >= 0 {
		o.o = slices.Delete(o.o, n, n+1)
	}
}

// Get returns element from map (if any).
func (o *Ordered[K, V]) Get(key K) (value V, ok bool) {
	if o.m == nil {
		return
	}

	value, ok = o.m[key]

	return
}

// Len returns number of elements in map.
func (o *Ordered[K, V]) Len() int {
	return len(o.o)
}

// Iter iterates elements in order of addition.
func (o *Ordered[K, V]) Iter(next func(key K, value V) bool) {
	for _, key := range o.o {
		value := o.m[key]

		if !next(key, value) {
			break
		}
	}
}

// Has checks key existence.
func (o *Ordered[K, V]) Has(key K) (yes bool) {
	_, yes = o.Get(key)

	return
}

// Clone creates shallow copy of map.
func (o *Ordered[K, V]) Clone() (rv *Ordered[K, V]) {
	return &Ordered[K, V]{
		m: maps.Clone(o.m),
		o: slices.Clone(o.o),
	}
}
