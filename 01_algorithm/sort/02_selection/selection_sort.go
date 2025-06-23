package main

import (
	util "coding/00_util/golang"
	"fmt"
)

func selectionSort(arr []int) []int {

	l := len(arr)
	for i := 0; i < l; i++ {
		tmpIndex := i
		for j := i; j < l; j++ {
			if arr[j] < arr[tmpIndex] {
				tmpIndex = j
			}
		}
		arr[i], arr[tmpIndex] = arr[tmpIndex], arr[i]
	}
	return arr
}

func main() {
	fmt.Println(selectionSort([]int{5, 4, 3, 2, 1}))
	fmt.Println(selectionSort(util.GenerateRandomArray(5)))
}
