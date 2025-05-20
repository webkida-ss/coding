package main

import "fmt"

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	leftIdx := binarySearch(nums, target, true)

	// leftIdxが配列の長さと等しいか、targetが見つからない場合
	if leftIdx == len(nums) || nums[leftIdx] != target {
		return result
	}

	result[0] = leftIdx
	result[1] = binarySearch(nums, target, false) - 1

	return result
}

func binarySearch(nums []int, target int, left bool) int {
	lo, hi := 0, len(nums)

	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] > target || (left && target == nums[mid]) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target)) // [3 4]
	fmt.Println()

	nums = []int{5, 7, 7, 8, 8, 10}
	target = 6
	fmt.Println(searchRange(nums, target)) // [-1 -1]

}
