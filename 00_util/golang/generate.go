package util

import (
	"math/rand"
	"time"
)

func GenerateRandomArray(size int) []int {
	// シード値を現在の時刻に基づいて設定
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, size)
	for i := range arr {
		arr[i] = r.Intn(100) // 0から99までのランダムな整数
	}
	return arr
}
