package main

import "fmt"

func intToRoman(num int) string {
	// ローマ数字の値：900などは事前に定義しておく、大きい順に並べる
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	// 事前に定義した値を順番に見ていく
	for i := 0; i < len(values) && num >= 0; i++ {
		for values[i] <= num {
			num -= values[i]
			roman += symbols[i]
		}
	}
	return roman
}

func main() {
	fmt.Println(intToRoman(3))    // Output: III
	fmt.Println(intToRoman(58))   // Output: LVIII
	fmt.Println(intToRoman(1994)) // Output: MCMXCIV
}
