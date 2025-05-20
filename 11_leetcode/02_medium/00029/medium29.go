package main

import "fmt"

// << : ビットシフト演算子（2のn乗）
func divide(dividend int, divisor int) int {

	// オーバーフローのケースを処理
	if dividend == -1<<31 && divisor == -1 {
		return 1<<31 - 1
	}

	// 結果の符号を決定
	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}

	// dividendとdivisorの絶対値を取得
	dividend = abs(dividend)
	divisor = abs(divisor)

	// 商を初期化
	quotient := 0
	// dividendがdivisor以上である限り
	for dividend >= divisor {
		temp := divisor
		m := 1
		// divisorがdividend以下になるまでdivisorを2倍にする
		for (temp << 1) <= dividend {
			temp <<= 1 // temp <<= 1は、temp = temp << 1
			m <<= 1
		}
		// 最終的なtempの値をdividendから引く
		dividend -= temp
		// 商にmを加算する
		quotient += m
	}
	// 最初に決定した符号を商に適用する
	return sign * quotient
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(divide(10, 3)) // Output: 3
	fmt.Println(divide(7, -3)) // Output: -2
}
