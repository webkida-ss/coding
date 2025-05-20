package main

import (
	"fmt"
)

func maxArea(height []int) int {
	// 左ポインタ
	left := 0
	// 右ポインタ
	right := len(height) - 1
	maxArea := 0

	// 左ポインタが右ポインタより右に来るまでループ
	for left < right {
		// 高さの小さい方を選択
		h := min(height[left], height[right])
		// 面積を計算
		area := h * (right - left)
		// 今までの最大面積と比較
		maxArea = max(maxArea, area)

		// 高さの小さい方のポインタを進める
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}
