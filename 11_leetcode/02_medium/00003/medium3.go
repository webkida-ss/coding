package main

import (
	"fmt"
)

// 文字列s内の重複しない最長の部分文字列の長さを返す関数
func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int) // 既に見た文字の最後の位置を保存
	maxLen := 0             // 最長部分文字列の長さ
	start := 0              // 現在の部分文字列の開始位置

	// sの各文字に対してループを行う。endは現在の文字の位置、charは現在の文字です。
	for end, char := range s {
		fmt.Println("========== 文字数分ループ ==========")
		fmt.Println(end, string(char))

		// 最後に見た位置と、その文字を前に見たかどうかを取得。
		lastPos, exists := m[char]
		fmt.Println("lastPos:", lastPos)
		fmt.Println("exists:", exists)

		// もし現在の文字が以前に見られそれが現在の部分文字列内にある場合、部分文字列の開始位置を更新
		if exists && lastPos >= start {
			start = lastPos + 1
		}
		// 現在の文字とその位置をマップに保存
		m[char] = end

		// 最長の部分文字列の長さを更新する必要があるかどうかを確認
		if end-start+1 > maxLen {
			fmt.Println("maxLen:", maxLen)
			maxLen = end - start + 1
		}
		fmt.Println()
	}

	return maxLen
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s)) // 3, "abc"
	// s = "bbbbbbb"
	// fmt.Println(lengthOfLongestSubstring(s)) // 1, "b"
	// s = "pwwkew"
	// fmt.Println(lengthOfLongestSubstring(s)) // 3, "wke"
}
