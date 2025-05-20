package main

import "fmt"

func jump(nums []int) int {
	// ジャンプ数
	jumps := 0
	// 現在のジャンプで到達できる最も遠い位置。つまり、現在のジャンプが終わる位置です。この位置に達したら、新たなジャンプを開始します。
	currentEnd := 0
	//　在の位置から次にジャンプするときに、到達可能な最も遠い位置。つまり、次のジャンプの可能な範囲の最遠端です。

	currentFarthest := 0
	for i := 0; i < len(nums)-1; i++ {
		fmt.Println("------------------")
		fmt.Println("jumps:", jumps)
		fmt.Println("currentEnd:", currentEnd)
		fmt.Println("currentFarthest:", currentFarthest)
		fmt.Println("i:", i)
		fmt.Println("nums[i]:", nums[i])
		fmt.Println()

		if i+nums[i] > currentFarthest {
			currentFarthest = i + nums[i]
		}
		if i == currentEnd {
			jumps++
			currentEnd = currentFarthest
		}
	}
	return jumps
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	result := jump(nums)
	println(result)
}
