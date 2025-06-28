package main

import (
	util "coding/00_util/golang"
	"fmt"
)

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
	arr1 := []int{5, 4, 1, 8}
	mergeSort(arr1)
	fmt.Println(arr1)
	arr2 := []int{5, 4, 3}
	mergeSort(arr2)
	fmt.Println(arr2)
	arr3 := []int{5, 4, 1, 8, 7, 3, 2, 9}
	mergeSort(arr3)
	fmt.Println(arr3)
	arr4 := util.GenerateRandomArray(10)
	mergeSort(arr4)
	fmt.Println(arr4)

}
