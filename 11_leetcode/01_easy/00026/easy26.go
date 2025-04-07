package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// このインデックスは次のユニークな数値の挿入場所を示しています
	index := 1

	// 2番目の文字から始める
	for i := 1; i < len(nums); i++ {
		// ソートされている前提なので、前の要素と比較して異なる要素の場合は異なる文字と判明する
		if nums[i] != nums[i-1] {
			// indexの位置に対象の文字列を差し込む
			nums[index] = nums[i]
			index++
		}
	}

	return index
}

func main() {
	// ユニークな文字数Kを探しつつ、K番目までの要素をソートする（それより後ろは任意の文字）
	nums := []int{0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums)
	fmt.Println("Number of unique elements:", k)
	fmt.Println("Modified nums:", nums[:k]) // nums = [0 1 2 3 4 2 2 3 3 4] ← K番目より後ろはどうでもいい
}
