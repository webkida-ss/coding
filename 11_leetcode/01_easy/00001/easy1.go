package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	// mapの初期化（キー:numsの値、値:numsのindex）
	m := make(map[int]int)

	// O(n)でnumsをループ
	for i, v := range nums {
		// 補数を計算
		complement := target - v
		// 補数がmapにあるか確認
		if _, ok := m[complement]; ok {
			// あれば、mapにあるindexと現在のindexを返す
			return []int{m[complement], i}
		}
		// なければ、mapに現在の値とindexを追加
		m[v] = i
	}
	return nil
}
