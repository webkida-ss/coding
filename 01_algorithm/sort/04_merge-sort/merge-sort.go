package main

import "fmt"

func mergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mid := len(arr) / 2
	left := make([]int, mid)
	right := make([]int, len(arr)-mid)

	copy(left, arr[:mid])
	copy(right, arr[mid:])

	// 左右の配列を再帰的にソート
	mergeSort(left)
	mergeSort(right)

	// マージ処理
	merge(arr, left, right)
}

// 二つの配列をマージする関数
func merge(arr, left, right []int) {
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	// 残りの要素をコピー：left,rightいずれかどっちがだけ残るので、どっちかのforは実行されない
	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("元の配列:", arr)
	mergeSort(arr)
	fmt.Println("ソートされた配列:", arr)
}
