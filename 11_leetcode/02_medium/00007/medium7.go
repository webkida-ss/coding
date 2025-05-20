package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var rev int
	for x != 0 {
		pop := x % 10 // 一番下の桁を取得（10で割ったときの余り=一番下の桁）
		x /= 10       // 一番下の桁を削除

		// int32の最大値と最小値をチェック
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && pop < -8) {
			return 0
		}

		rev = rev*10 + pop // 取得した桁を反転した数の末尾に追加
	}
	return rev
}

func main() {
	fmt.Println(reverse(123))  // 321
	fmt.Println(reverse(-123)) // -321
	fmt.Println(reverse(120))  // 21
	fmt.Println(reverse(0))    // 0
	// 次の例はオーバーフローするので0を返すべき
	fmt.Println(reverse(1534236469)) // 0
}
