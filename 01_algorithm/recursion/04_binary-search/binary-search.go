package main

import "fmt"

func binarySearch(list []int, target int) int {
	left, right := 0, len(list)-1

	for left <= right {
		// 中間点
		mid := left + (right-left)/2
		midVal := list[mid]

		if midVal == target {
			return mid
		} else if midVal < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// 基底ケース
	return -1
}

func main() {
	fmt.Println(binarySearch([]int{1}, 1))                             // 0
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)) // 2
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9)) // 8
	// テストケース
	fmt.Println(binarySearch([]int{1, 3, 5, 7, 9}, 5))       // 2
	fmt.Println(binarySearch([]int{2, 4, 6, 8, 10}, 2))      // 0
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 5))       // 4
	fmt.Println(binarySearch([]int{10, 20, 30, 40, 50}, 25)) // -1
	fmt.Println(binarySearch([]int{}, 1))                    // -1
	fmt.Println(binarySearch([]int{4}, 4))                   // 0
	fmt.Println(binarySearch([]int{4}, 3))                   // -1
}
