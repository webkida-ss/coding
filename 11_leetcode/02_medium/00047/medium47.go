package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	results := [][]int{}
	backtrack(&results, nums, []int{}, map[int]bool{})
	return results
}

func backtrack(results *[][]int, nums []int, current []int, used map[int]bool) {
	fmt.Println("------------ backtrack ------------")
	fmt.Println("results:", results)
	fmt.Println("nums:", nums)
	fmt.Println("current:", current)
	fmt.Println("used:", used)
	fmt.Println()

	if len(current) == len(nums) {
		fmt.Println("add result")
		fmt.Println()
		// 現在の順列をコピーし、そのコピーを結果のリストに追加
		temp := make([]int, len(current))
		copy(temp, current)
		*results = append(*results, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		fmt.Println("i:", i)
		fmt.Println()

		// 現在の要素が既に使用されているか、または現在の要素が前の要素と同じで、前の要素がまだ使用されていない場合に、ループをスキップ
		if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
			fmt.Println("skip")
			fmt.Println()
			continue
		}
		fmt.Println("not skip")
		used[i] = true                          // 現在の要素を「使用済み」とマーク
		current = append(current, nums[i])      // 現在の要素を現在の順列に追加
		backtrack(results, nums, current, used) // 再帰的にバックトラッキング関数を呼び出し

		fmt.Println("backtrack")
		// バックトラッキング
		current = current[:len(current)-1] // 現在の順列から最後に追加した要素を削除
		used[i] = false                    // 現在の要素を「未使用」にマーク
	}
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}
