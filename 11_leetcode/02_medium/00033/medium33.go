package main

import "fmt"

func search(nums []int, target int) int {
	// 左右ポインタ
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		fmt.Println("------------------")
		fmt.Println("left:", left)
		fmt.Println("right:", right)
		fmt.Println("mid:", mid)
		fmt.Println()

		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {
			fmt.Println("左側がソートされている場合")
			// nums[left] <= target < nums[mid] の場合
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}

		} else {
			fmt.Println("右側がソートされている場合")
			// nums[mid] < target <= nums[right] の場合
			if target <= nums[right] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {

	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target)) //4
	fmt.Println()

	// nums = []int{4, 5, 6, 7, 0, 1, 2}
	// target = 3
	// fmt.Println(search(nums, target)) //1
}
