package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	var result []int
	if len(matrix) == 0 {
		return result
	}
	rowBegin, rowEnd := 0, len(matrix)-1
	colBegin, colEnd := 0, len(matrix[0])-1
	for rowBegin <= rowEnd && colBegin <= colEnd {
		// 右へ移動
		for i := colBegin; i <= colEnd; i++ {
			result = append(result, matrix[rowBegin][i])
		}
		rowBegin++

		// 下へ移動
		for i := rowBegin; i <= rowEnd; i++ {
			result = append(result, matrix[i][colEnd])
		}
		colEnd--

		// 左へ移動
		if rowBegin <= rowEnd {
			for i := colEnd; i >= colBegin; i-- {
				result = append(result, matrix[rowEnd][i])
			}
		}
		rowEnd--

		// 上へ移動
		if colBegin <= colEnd {
			for i := rowEnd; i >= rowBegin; i-- {
				result = append(result, matrix[i][colBegin])
			}
		}
		colBegin++
	}

	return result
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(matrix))
}
