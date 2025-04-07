package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {

	// ソート
	sort.Ints(nums)

	// 結果リスト
	result := [][]int{}

	// 3つの数を選ぶので、最後から2つ目までループ
	for i := 0; i < len(nums)-2; i++ {
		// 重複をスキップ (i > 0 は、i-1が存在することを確認するため)
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 2つの数を選ぶので、最初と最後から1つずつループ
		l, r := i+1, len(nums)-1
		// rがlより大きい間ループ
		for l < r {
			// 合計が0なら結果に追加
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				// 重複をスキップ (l < r は、l+1が存在することを確認するため)
				for l < r && nums[l] == nums[l+1] {
					// lをインクリメント（左ポインタを右に移動）
					l++
				}
				// 重複をスキップ (l < r は、r-1が存在することを確認するため)
				for l < r && nums[r] == nums[r-1] {
					// rをデクリメント（右ポインタを左に移動）
					r--
				}
				// lをインクリメント（左ポインタを右に移動）
				l++
				// rをデクリメント（右ポインタを左に移動）
				r--
			} else if sum < 0 {
				// lをインクリメント（左ポインタを右に移動）
				l++
			} else {
				// rをデクリメント（右ポインタを左に移動）
				r--
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
