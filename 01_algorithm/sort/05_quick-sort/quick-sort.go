package main

import "fmt"

func quickSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}
	left, right := 0, len(arr)-1

	// ピボットはスライスの中央を選択
	pivotIndex := len(arr) / 2

	// ピボットの要素と配列の最後の要素を交換
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// 配列全体の要素でループ
	// 配列: [3, 7, 8, 2, 1, 5, 4]
	// ピボット: 5 (最後の要素)
	// ↓
	// ↓ 5をピボットとして左は5よりも小さい、右は5よりも大きい
	// 配列: [3, 2, 1, 4,   5,    7, 8]
	for i := range arr {
		// ☆ 現在の要素が配列の最後の要素（ピボット）より小さいかをチェック
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
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

	fmt.Println("----------------------------------------")
	arr0 := []int{1}
	fmt.Println("Original array:", arr0)
	sortedArr0 := quickSort(arr0)
	fmt.Println("Sorted array:", sortedArr0)
	fmt.Println()

	fmt.Println("----------------------------------------")
	arr1 := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	fmt.Println("Original array:", arr1)
	sortedArr1 := quickSort(arr1)
	fmt.Println("Sorted array:", sortedArr1)
}
