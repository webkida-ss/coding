package main

import (
	"fmt"
)

type entry[K comparable, V any] struct {
	key   K
	value V
}

type HashTable[K comparable, V any] struct {
	buckets [][]entry[K, V]
	size    int
}

func NewHashTable[K comparable, V any](size int) *HashTable[K, V] {
	return &HashTable[K, V]{
		buckets: make([][]entry[K, V], size),
		size:    size,
	}
}

func hash[K comparable](key K, size int) int {
	return int(fmt.Sprintf("%v", key)[0]) % size
}

func (ht *HashTable[K, V]) Set(key K, value V) {
	index := hash(key, ht.size)
	bucket := ht.buckets[index]

	for i, e := range bucket {
		if e.key == key {
			bucket[i].value = value
			return
		}
	}
	ht.buckets[index] = append(bucket, entry[K, V]{key, value})
}

func (ht *HashTable[K, V]) Get(key K) (V, bool) {
	index := hash(key, ht.size)
	bucket := ht.buckets[index]

	for _, e := range bucket {
		if e.key == key {
			return e.value, true
		}
	}
	var zero V
	return zero, false
}

func main() {
	ht := NewHashTable[int, string](10)
	ht.Set(123, "hello")
	ht.Set(456, "world")

	if value, ok := ht.Get(123); ok {
		fmt.Println("Found:", value)
	}
}
