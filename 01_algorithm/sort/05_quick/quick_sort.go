package main

import (
	util "coding/00_util/golang"
	"fmt"
)

func quickSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}
	left, right := 0, len(arr)-1

	// ----------------------------------------------------------------------------------------
	// 中央をピボットに選ぶ理由は、平均的に良い分割を得やすいからです（特にソート済みや偏ったデータに対して、パフォーマンスが安定する傾向がある）。
	// ただし、実装がシンプルで良いのは末尾をピボットにする方法です。

	// 配列: [5, 4, 3, 1, 2]
	pivotIndex := len(arr) / 2                                // ピボットはスライスの中央を選択=3
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex] // ピボットの要素と配列の最後の要素を交換
	// 配列: [5, 4, 2, 1, 3]
	// ----------------------------------------------------------------------------------------

	// ピボット: 3 (最後の要素)：左は3よりも小さい、右は3よりも大きい に分ける
	for i := range arr {
		currentValue := arr[i]
		pivotValue := arr[right]
		if currentValue < pivotValue { // ☆ 現在の要素が配列の最後の要素（ピボット）より小さいかをチェック
			arr[i], arr[left] = arr[left], arr[i] // 現在の要素と左の要素を交換する
			left++                                // 左を1つ右にずらす
		}

	}
	// ピボットの位置を戻す
	arr[left], arr[right] = arr[right], arr[left]

	// 左側を再帰的にソートし直す
	quickSort(arr[:left])
	// 右側を再帰的にソートし直す
	quickSort(arr[left+1:])

	return arr
}

func main() {
	fmt.Println(quickSort([]int{5, 4, 3, 1, 2}))
	fmt.Println(quickSort([]int{4, 3, 2, 1}))
	fmt.Println(quickSort([]int{1}))
	fmt.Println(quickSort([]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}))
	fmt.Println(quickSort(util.GenerateRandomArray(10)))
}

// [5, 4, 3, 1, 2]
// pivot = 3
// pivotと最後を入れ替え
// [5, 4, 2, 1, 3]
// pivot = 3と 各要素を比較し、3よりも小さいものを左へ、大きいものを右へ
// ループ
//  i = 0 → val = 5 do nothing
//  i = 1 → val = 4 do nothing
//  i = 2 → val = 2 左の要素と入れ替え、左indexをインクリメント
//     [2, 4, 5, 1, 3]
//  i = 3 → val = 1 左の要素と入れ替え、左indexをインクリメント
//     [2, 1, 5, 4, 3]
//  i = 4 → val = 3 do nothing
// ループ終了
// [2, 1, 5, 4, 3]
// pivotを左（index=2,val=5）と入れ替え
// [2, 1, 3, 4, 5]
// 左index = 2,右index=5
// ☆☆☆ ここでわかるのは、左indexを堺にそのvalueよりも小さいものが左へ、大きいものが右になることがわかる ☆☆☆
//
