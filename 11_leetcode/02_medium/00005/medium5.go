package main

import (
	"fmt"
)

// 文字列sから最も長い回文部分文字列を返す関数
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	// 文字列の開始位置と終了位置
	start, end := 0, 0

	// 文字列の各文字を中心として、その周りで回文を探します。
	for i := 0; i < len(s); i++ {
		fmt.Println("========== 文字数分ループ ==========")

		// 奇数長の回文を探す
		len1 := expandAroundCenter(s, i, i)
		// 偶数長の回文を探す：rightを1つずらす
		len2 := expandAroundCenter(s, i, i+1)
		// 大きい方
		maxLen := max(len1, len2)
		fmt.Println("len1:", len1)
		fmt.Println("len2:", len2)
		fmt.Println("maxLen:", maxLen)

		// maxLen（現在の最大の回文）と end-start（前回までの最大の回文）を比較
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2

			fmt.Println("start:", start)
			fmt.Println("end:", end)
		}

		fmt.Println()
	}

	// 最も長い回文部分文字列を返す
	return s[start : end+1]
}

// 中心から左右に拡張して最大の回文の長さを返す関数
func expandAroundCenter(s string, left int, right int) int {
	if len(s) == 0 || left > right {
		return 0
	}
	// 文字列の境界を超えていない限り、左右に拡張
	for left >= 0 && right < len(s) && s[left] == s[right] {
		// left >= 0： 左が0以上
		// right < len(s)： 右が文字列の最大長
		// s[left] == s[right]：左右に拡張したときに文字が一致するか
		left--
		right++
	}

	// 回文の長さを返す
	return right - left - 1
}

// 2つの整数のうち大きい方を返す
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s)) // "bab" または "aba"

	// s = "cbbd"
	// fmt.Println(longestPalindrome(s)) // "bb"
}
