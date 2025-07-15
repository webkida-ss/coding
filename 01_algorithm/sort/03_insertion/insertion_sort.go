package main

import "fmt"

func insertionSort(arr []int) []int {
	l := len(arr)
	for i := 0; i < l; i++ {
		tmp := arr[i] // 1個ずつずらすから値を保持しておく
		j := i - 1
		for j >= 0 && arr[j] > tmp {
			arr[j+1] = arr[j] // 右にずらす
			j--
		}
		arr[j+1] = tmp // j+1は帳尻合わせ
	}
	return arr
}

func main() {
	fmt.Println(insertionSort([]int{1, 3, 7, 2, 8, 5}))
	fmt.Println(insertionSort([]int{10, 7, 3, 2, 8, 5}))
}

// 3 9 6 1 2
// ↓
// 3 9 6 1 2
//   ↓
// 3 9 6 1 2
//     ↓
// 3 6 9 1 2
//       ↓
// 1 3 6 9 2
//         ↓
// 1 2 3 6 9
