package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	backtrack(candidates, target, &result, []int{}, 0)
	return result
}

func backtrack(candidates []int, target int, result *[][]int, current []int, index int) {
	fmt.Println("--------- backtrack ---------")
	fmt.Println(candidates)
	fmt.Println(target)
	fmt.Println(result)
	fmt.Println(current)
	fmt.Println(index)
	fmt.Println()

	if target < 0 {
		return
	}
	if target == 0 {
		*result = append(*result, append([]int{}, current...))
	}
	for i := index; i < len(candidates); i++ {
		current = append(current, candidates[i])
		backtrack(candidates, target-candidates[i], result, current, i)
		// 取り消し（最後の要素を取り除く）
		current = current[:len(current)-1]
	}
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}
