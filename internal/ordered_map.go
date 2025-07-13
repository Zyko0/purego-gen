package internal

import (
	"iter"
	"slices"
)

type OrderedMap[K comparable, V any] struct {
	keys  []K
	inner map[K]V
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		keys:  make([]K, 0),
		inner: make(map[K]V),
	}
}

func (om *OrderedMap[K, V]) Set(key K, value V) {
	if _, exists := om.inner[key]; !exists {
		om.keys = append(om.keys, key)
	}
	om.inner[key] = value
}

func (om *OrderedMap[K, V]) Get(key K) (V, bool) {
	value, exists := om.inner[key]
	return value, exists
}

func (om *OrderedMap[K, V]) Delete(key K) {
	if _, exists := om.inner[key]; exists {
		delete(om.inner, key)
		idx := slices.Index(om.keys, key)
		om.keys = slices.Delete(om.keys, idx, idx+1)
	}
}

func (om *OrderedMap[K, V]) Len() int {
	return len(om.keys)
}

func (om *OrderedMap[K, V]) All() iter.Seq2[K, V] {
	return iter.Seq2[K, V](func(yield func(K, V) bool) {
		for _, key := range om.keys {
			if !yield(key, om.inner[key]) {
				return
			}
		}
	})
}
