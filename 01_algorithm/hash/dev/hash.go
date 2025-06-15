package main

import (
	"fmt"
)

// エントリを格納する構造体
type entry struct {
	key   string
	value string
}

// ハッシュテーブルの定義
type HashTable struct {
	buckets [][]entry
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
func NewHashTable(size int) *HashTable {
	return &HashTable{
		buckets: make([][]entry, size),
		size:    size,
	}
}

// 値を追加
func (ht *HashTable) Set(key, value string) {
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
	ht.buckets[index] = append(bucket, entry{key, value})
}

// 値を取得
func (ht *HashTable) Get(key string) (string, bool) {
	index := hash(key, ht.size)
	bucket := ht.buckets[index]

	for _, e := range bucket {
		if e.key == key {
			return e.value, true
		}
	}
	return "", false
}

// 使用例
func main() {
	ht := NewHashTable(10)
	ht.Set("name", "Taro")
	ht.Set("age", "30")

	if val, ok := ht.Get("name"); ok {
		fmt.Println("name:", val)
	}
}
