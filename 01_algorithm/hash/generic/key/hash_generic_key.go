package main

import (
	"fmt"
)

type entry[K comparable] struct {
	key   K
	value string
}

type HashTable[K comparable] struct {
	buckets [][]entry[K]
	size    int
}

func NewHashTable[K comparable](size int) *HashTable[K] {
	return &HashTable[K]{
		buckets: make([][]entry[K], size),
		size:    size,
	}
}

func hash[K comparable](key K, size int) int {
	return int(fmt.Sprintf("%v", key)[0]) % size
}

func (ht *HashTable[K]) Set(key K, value string) {
	index := hash(key, ht.size)
	bucket := ht.buckets[index]

	for i, e := range bucket {
		if e.key == key {
			bucket[i].value = value
			return
		}
	}
	ht.buckets[index] = append(bucket, entry[K]{key, value})
}

func (ht *HashTable[K]) Get(key K) (string, bool) {
	index := hash(key, ht.size)
	bucket := ht.buckets[index]

	for _, e := range bucket {
		if e.key == key {
			return e.value, true
		}
	}
	return "", false
}

func main() {
	ht := NewHashTable[int](10)
	ht.Set(123, "hello")
	ht.Set(456, "world")

	if value, ok := ht.Get(123); ok {
		fmt.Println("Found:", value)
	}
}
