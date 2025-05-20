package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	prev := countAndSay(n - 1)
	result := ""
	count := 1

	fmt.Println("--------------")
	fmt.Println(n)
	fmt.Println(prev)
	fmt.Println(result)
	fmt.Println(count)
	fmt.Println()

	// n-1番目の項目（prev）をスキャンし、n番目の項目（result）を生成するためのループ
	for i := 1; i < len(prev); i++ {
		fmt.Println("prev[i]:", prev[i])
		fmt.Println("prev[i-1]:", prev[i-1])
		if prev[i] == prev[i-1] {
			count++
		} else {
			result += strconv.Itoa(count) + string(prev[i-1])
			count = 1
		}
	}

	// カウント+文字列
	fmt.Println("count", count)
	fmt.Println("string(prev[len(prev)-1]:", string(prev[len(prev)-1]))

	// prev[len(prev)-1]:最後のパターンをここで拾う
	result += strconv.Itoa(count) + string(prev[len(prev)-1])
	return result
}

func main() {
	fmt.Println(countAndSay(6)) // Output: "111221"
}
