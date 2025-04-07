package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(1221))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(1))
	fmt.Println()
	fmt.Println(isPalindrome2(121))
	fmt.Println(isPalindrome2(1221))
	fmt.Println(isPalindrome2(-121))
	fmt.Println(isPalindrome2(10))
	fmt.Println(isPalindrome2(1))
}

// 実行速度こっちのほうが実行速度が早い
func isPalindrome(x int) bool {
	strX := strconv.Itoa(x)
	n := len(strX)
	for i := 0; i < n/2; i++ {
		// 逆から見ていく
		if strX[i] != strX[n-1-i] {
			// 違う時点で回分ではない
			return false
		}
	}
	return true
}

func isPalindrome2(x int) bool {
	strX := strconv.Itoa(x)

	left, right := 0, len(strX)-1
	for left < right {
		if strX[left] != strX[right] {
			return false
		}
		left++
		right--
	}

	return true
}
