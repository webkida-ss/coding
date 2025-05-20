package main

import "fmt"

func insertionSort(arr []int) []int {
	l := len(arr)
	for i := 0; i < l; i++ {
		tmp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > tmp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = tmp
	}
	return arr
}

func main() {
	fmt.Println(insertionSort([]int{1, 7, 3, 2, 8, 5}))
}
