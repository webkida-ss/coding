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

func binarySearchRecursion(list []int, target int) int {

	var _binarySearch func(_list []int, _target int, _leftIndex, _rightIndex int) int
	_binarySearch = func(_list []int, _target int, _leftIndex, _rightIndex int) int {

		if _leftIndex > _rightIndex {
			return -1
		}
		mid := (_leftIndex + _rightIndex) / 2
		if _list[mid] == target {
			return mid
		} else if _list[mid] < target {
			return _binarySearch(_list, target, mid+1, _rightIndex)
		} else {
			return _binarySearch(_list, target, _leftIndex, mid-1)
		}
	}

	return _binarySearch(list, target, 0, len(list)-1)

}

func main() {
	//type test struct {
	//	input  []int
	//	target int
	//	expect int
	//}
	//tests := []test{
	//	{
	//		input:  []int{1},
	//		target: 1,
	//		expect: 0,
	//	},
	//	{
	//		input:  []int{1},
	//		target: 1,
	//		expect: 0
	//	},
	//}
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
	fmt.Println("---------------------------------------------------")
	fmt.Println(binarySearchRecursion([]int{1}, 1))                             // 0
	fmt.Println(binarySearchRecursion([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)) // 2
	fmt.Println(binarySearchRecursion([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9)) // 8
	// テストケース
	fmt.Println(binarySearchRecursion([]int{1, 3, 5, 7, 9}, 5))       // 2
	fmt.Println(binarySearchRecursion([]int{2, 4, 6, 8, 10}, 2))      // 0
	fmt.Println(binarySearchRecursion([]int{1, 2, 3, 4, 5}, 5))       // 4
	fmt.Println(binarySearchRecursion([]int{10, 20, 30, 40, 50}, 25)) // -1
	fmt.Println(binarySearchRecursion([]int{}, 1))                    // -1
	fmt.Println(binarySearchRecursion([]int{4}, 4))                   // 0
	fmt.Println(binarySearchRecursion([]int{4}, 3))                   // -1
}
