package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	fmt.Println()
	fmt.Println(nums)
	maxSoFar := nums[0]      // これまでの最大部分配列の和
	maxEndingHere := nums[0] // 現在検討中の部分配列の最大和

	for i := 1; i < len(nums); i++ {
		fmt.Println("maxSoFar:", maxSoFar)
		fmt.Println("maxEndingHere", maxEndingHere)
		fmt.Println("nums[i]", nums[i])

		current := nums[i]                                  // 現在考慮中の配列の要素
		maxEndingHere = max(current, maxEndingHere+current) // 現在の要素を、これまでに見つかった最大和を持つ部分配列の和に加えた値と、現在の要素の大きい方を選ぶ
		maxSoFar = max(maxSoFar, maxEndingHere)
	}

	return maxSoFar
}

// 比較して大きい方を返す
func max(a, b int) int {
	fmt.Println("a:", a, "b:", b)
	if a > b {
		fmt.Printf("%d > %d\n", a, b)
		return a
	}
	fmt.Printf("%d <= %d\n", a, b)
	return b
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
	fmt.Println(maxSubArray([]int{-1, -2, -3, -4}))
	fmt.Println(maxSubArray([]int{-2, 1}))
	fmt.Println(maxSubArray([]int{-2, -1}))
}
