package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {

	// ソート： [-1, 2, 1, -4] -> [-4, -1, 1, 2]
	// ↓
	// [-4,           -1,         1,          2]
	//  ↑             ↑           ↑           ↑
	// index = 0, left = 1,                right = 3
	// ↓
	// index = 0, left = 1,    right = 2
	// ↓
	//              index = 1, left = 2,    right = 3

	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		// 左を1つ進める、右を一番右にする
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			// 絶対値を比較
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-closest)) {
				closest = sum
			}
			// targetとの差が小さくなるようにleftとrightを移動
			if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return closest
}

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target)) // Output: 2
}
