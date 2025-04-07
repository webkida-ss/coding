package main

import "fmt"

func gcd(a, b int) int {
	// 余り
	r := a % b

	// 基底ケース
	if r == 0 {
		return b
	}

	// 再帰呼び出し： a → b, b → r　とするのがユークリッドの互除法
	return gcd(b, r)
}

func main() {
	fmt.Println(gcd(28, 14)) // 14
	// テストケース
	fmt.Println(gcd(14, 28))   // 14
	fmt.Println(gcd(8, 12))    // 4
	fmt.Println(gcd(17, 31))   // 1
	fmt.Println(gcd(270, 192)) // 6
	fmt.Println(gcd(25, 25))   // 25
	fmt.Println(gcd(0, 15))    // 15
}
