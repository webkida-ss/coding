package main

import "fmt"

func strStr(haystack string, needle string) int {
	// needleが空文字の場合
	if needle == "" {
		return 0
	}

	n, m := len(haystack), len(needle)

	// haystackの各インデックスを確認する
	for i := 0; i <= n-m; i++ {
		if haystack[i:i+m] == needle {
			return i
		}
	}

	// needleがhaystackに含まれていない場合
	return -1
}

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStr(haystack, needle))
}
