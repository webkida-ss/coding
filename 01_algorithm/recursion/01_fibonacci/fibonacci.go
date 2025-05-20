package main

import "fmt"

func fibonacci(n int) int {
	// 終了条件（基底ケース）
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(1))  // 1
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(3))  // 2
	fmt.Println(fibonacci(4))  // 3
	fmt.Println(fibonacci(5))  // 5
	fmt.Println(fibonacci(10)) // 55
}
