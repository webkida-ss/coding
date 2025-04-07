package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	dfs(candidates, target, 0, []int{}, &res)
	return res
}

func dfs(nums []int, target, idx int, path []int, res *[][]int) {
	// 合計数値が一致した場合にのみappendする（その他は捨てる）
	if target == 0 {
		*res = append(*res, append([]int(nil), path...))
		return
	}
	for i := idx; i < len(nums); i++ {
		if target-nums[i] < 0 {
			break
		}
		// 重複組み合わせを排除
		if i > idx && nums[i] == nums[i-1] {
			continue
		}
		dfs(nums, target-nums[i], i+1, append(path, nums[i]), res)
	}
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}
