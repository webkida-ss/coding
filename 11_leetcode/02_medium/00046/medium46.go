package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	var result [][]int
	backtrack(nums, nil, &result)
	return result
}

func backtrack(nums []int, current []int, result *[][]int) {
	fmt.Println("--------- backtrack ---------")
	fmt.Println("nums:", nums)
	fmt.Println("current:", current)
	fmt.Println("result:", result)
	fmt.Println()

	if len(nums) == 0 {
		*result = append(*result, append([]int(nil), current...))
		return
	}
	for i := 0; i < len(nums); i++ {
		// numsスライスからi番目の要素を除いた新しいスライスnextを作成
		next := append([]int(nil), nums[:i]...)
		next = append(next, nums[i+1:]...)
		
		backtrack(next, append(current, nums[i]), result)
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
