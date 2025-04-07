package main

import (
	"fmt"
)

func romanToInt(s string) int {
	// ローマ字マップ
	romanValues := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// 合計値
	result := 0

	// 解法①：for iでループして1文字ずつ取得
	// iで取れるのはバイトごとなので、マルチバイト文字の場合はできない

	// 文字列を1文字ずつ処理
	for i := 0; i < len(s); i++ { // len(s)-1としてはイケない：最後の文字を処理できなくなる
		// 最後ではない && 次の文字の方が大きい場合
		currentValue := romanValues[rune(s[i])] // mapへのアクセスは複数回やると結構遅いので、回数を留める
		if i < len(s)-1 && currentValue < romanValues[rune(s[i+1])] {
			// I,X,Cなど引く場合
			result -= currentValue
		} else {
			// その他足す場合
			result += currentValue
		}
	}

	// 解法②：for range
	// rune型でマルチバイトも考慮してやってくれる、ただし、preValueのように別変数を用意しておく必要がある
	// total := 0
	// prevValue := 0

	// for _, char := range s {
	// 	value := romanMap[char]
	// 	if value > prevValue {
	// 		// 前の値を引く（すでに加算されているため）
	// 		total += value - 2*prevValue
	// 	} else {
	// 		total += value
	// 	}
	// 	prevValue = value
	// }

	return result
}

func main() {
	tests := []string{"III", "IV", "IX", "LVIII", "MCMXCIV"}
	for _, test := range tests {
		fmt.Printf("%s = %d\n", test, romanToInt(test))
	}
}
