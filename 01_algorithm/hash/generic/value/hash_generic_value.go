package main

import (
	"fmt"
)

// エントリを格納する構造体
type entry[V any] struct {
	key   string
	value V
}

// ハッシュテーブルの定義
type HashTable[V any] struct {
	buckets [][]entry[V]
	size    int
}

// ハッシュ関数（シンプルな例）
func hash(key string, size int) int {
	hashValue := 0
	for _, c := range key {
		hashValue += int(c)
	}
	return hashValue % size
}

// ハッシュテーブルを初期化
func NewHashTable[V any](size int) *HashTable[V] {
	return &HashTable[V]{
		buckets: make([][]entry[V], size),
		size:    size,
	}
}

// 値を追加
func (ht *HashTable[V]) Set(key string, value V) {
	index := hash(key, ht.size)
	bucket := ht.buckets[index]

	// すでに同じキーがある場合は更新
	for i, e := range bucket {
		if e.key == key {
			ht.buckets[index][i].value = value
			return
		}
	}

	// 新しいエントリを追加
	ht.buckets[index] = append(bucket, entry[V]{key, value})
}

// 値を取得
func (ht *HashTable[V]) Get(key string) (V, bool) {
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

// 使用例
func main() {
	ht := NewHashTable[string](10)
	ht.Set("name", "Taro")
	ht.Set("age", "30")

	if val, ok := ht.Get("name"); ok {
		fmt.Println("name:", val)
	}
}
