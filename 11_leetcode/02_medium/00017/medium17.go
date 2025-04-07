package main

import "fmt"

var phone = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	result := []string{""}
	// 1数字ずつ見ていく
	for _, digit := range digits {
		temp := []string{}
		// 今までの結果に今回の数字に対応する文字を足していく
		for _, res := range result {
			// 今回の数字に対応する文字を足していく
			for _, char := range phone[string(digit)] {
				temp = append(temp, res+string(char))
			}
		}
		result = temp
	}
	return result
}

func main() {
	fmt.Println(letterCombinations("23"))
}
