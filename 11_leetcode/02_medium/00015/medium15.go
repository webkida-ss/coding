package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {

	sort.Ints(nums)

	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {

		left, right := i+1, len(nums)-1

		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 重複削除
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right-1] == nums[right] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-4, -1, -1, 0, 1, 2}))
	fmt.Println(threeSum([]int{-4, -1, -1, -1, -1, 2}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
}
