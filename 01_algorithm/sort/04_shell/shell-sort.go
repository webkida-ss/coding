package main

import "fmt"

func shellSort(arr []int) []int {
	l := len(arr)
	gap := l / 2
	for gap > 0 {
		for i := 0; i < l; i++ {
			tmp := arr[i]
			j := i
			for j >= gap && arr[j-gap] > tmp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = tmp
		}
		gap = gap / 2
	}
	return arr
}

func main() {
	fmt.Println(shellSort([]int{5, 6, 9, 2, 3}))
}
