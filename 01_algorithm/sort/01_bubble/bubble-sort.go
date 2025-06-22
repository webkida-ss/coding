package main

import (
	util "coding/00_util/golang"
	"fmt"
)

func bubbleSort(arr []int) {
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := 0; j < l-1-i; j++ {
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

func main() {
	sort([]int{12, 11, 13, 5, 6, 7})
	sort(util.GenerateRandomArray(5))
	sort(util.GenerateRandomArray(5))
}
