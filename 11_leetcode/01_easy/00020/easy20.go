package main

import (
	"fmt"
)

func isValid(s string) bool {

	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []rune{} // make([rune],0)

	// 文字列を1文字ずつ処理
	for _, char := range s {

		switch char {

		// 開始括弧の場合はスタックに追加
		case '(', '{', '[':
			stack = append(stack, char)

		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			// 閉じ括弧が正しい場合：スタックの最後の要素を削除
			stack = stack[:len(stack)-1]

		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([)]"))   // false
	fmt.Println(isValid("{[]}"))   // true
	fmt.Println(isValid("]"))      // false
}
