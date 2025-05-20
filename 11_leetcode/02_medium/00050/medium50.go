package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1 // 基底ケース
	} else if n < 0 {
		// 負のべき乗
		return 1 / myPow(x, -n)
	} else if n%2 == 0 {
		// 指数が偶数の場合（分割統治法によりこの場合を高速化）
		return myPow(x*x, n/2)
	} else {
		// 指数が奇数の場合
		return x * myPow(x, n-1)
	}
}

func main() {
	fmt.Println(myPow(2.00000, 10)) // Output: 1024.00000
	fmt.Println(myPow(2.10000, 3))  // Output: 9.26100
	fmt.Println(myPow(2.00000, -2)) // Output: 0.25000
}
