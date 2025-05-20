package main

import "fmt"

func removeElement(nums []int, val int) int {
	pos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[pos] = nums[i]
			pos++
		}
	}
	return pos
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	result := removeElement(nums, val)

	fmt.Println("New length:", result)            // "New length: 2" を出力
	fmt.Println("Modified array:", nums[:result]) // "Modified array: [2 2]" を出力
}
