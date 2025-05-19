package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(arr []int) {
	arrLength := len(arr)
	for i := 0; i < arrLength; i++ {
		for j := 0; j < arrLength-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func sort(arr []int) {
	fmt.Println("----------------------------------------")
	fmt.Println("元の配列:", arr)
	bubbleSort(arr)
	fmt.Println("ソートされた配列:", arr)
	fmt.Println("----------------------------------------")
}

func generateRandomArray(size int) []int {
	// シード値を現在の時刻に基づいて設定
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, size)
	for i := range arr {
		arr[i] = r.Intn(100) // 0から99までのランダムな整数
	}
	return arr
}

func main() {
	sort([]int{12, 11, 13, 5, 6, 7})
	sort(generateRandomArray(5))
}
