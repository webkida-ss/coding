package main

import "fmt"

func factorial(n int) int {

	// 基底ケース
	if n <= 1 {
		return 1
	}

	// 再帰
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(0))
	fmt.Println(factorial(1))
	fmt.Println(factorial(2))
	fmt.Println(factorial(3))
	fmt.Println(factorial(4))
	fmt.Println(factorial(5))
	fmt.Println(factorial(10))
}
