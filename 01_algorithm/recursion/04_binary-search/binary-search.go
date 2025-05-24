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
	type test struct {
		input  []int
		target int
		expect int
	}
	tests := []test{
		{input: []int{1}, target: 1, expect: 0},
		{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 3, expect: 2},
		{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 9, expect: 8},
		{input: []int{1, 3, 5, 7, 9}, target: 5, expect: 2},
		{input: []int{2, 4, 6, 8, 10}, target: 2, expect: 0},
		{input: []int{1, 2, 3, 4, 5}, target: 5, expect: 4},
		{input: []int{10, 20, 30, 40, 50}, target: 25, expect: -1},
		{input: []int{}, target: 1, expect: -1},
		{input: []int{4}, target: 4, expect: 0},
		{input: []int{4}, target: 3, expect: -1},
	}

	fmt.Println("binarySearch results:")
	for i, tc := range tests {
		result := binarySearch(tc.input, tc.target)
		if result == tc.expect {
			fmt.Printf("[OK]   #%d input=%v, target=%d, got=%d\n", i+1, tc.input, tc.target, result)
		} else {
			fmt.Printf("[FAIL] #%d input=%v, target=%d, got=%d, expect=%d\n", i+1, tc.input, tc.target, result, tc.expect)
		}
	}

	fmt.Println("---------------------------------------------------")
	fmt.Println("binarySearchRecursion results:")
	for i, tc := range tests {
		result := binarySearchRecursion(tc.input, tc.target)
		if result == tc.expect {
			fmt.Printf("[OK]   #%d input=%v, target=%d, got=%d\n", i+1, tc.input, tc.target, result)
		} else {
			fmt.Printf("[FAIL] #%d input=%v, target=%d, got=%d, expect=%d\n", i+1, tc.input, tc.target, result, tc.expect)
		}
	}
}
