package main

import (
	"fmt"
	"strings"
)

// 各行を連結しつつ文字列を作成し、最後に連結した文字列を返す
func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	// 各行の文字を格納するスライスを初期化
	rows := make([]string, numRows)
	curRow := 0        // 現在の行
	goingDown := false // ジグザグの方向

	// 文字列を1文字ずつ処理
	for _, char := range s {
		rows[curRow] += string(char) // 現在の行に文字を追加（連結）

		// ジグザグパターンの方向を切り替える（行が1番上か1番下に到達したら）
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}

		// ダウンなら行を1つ下へ、アップなら行を1つ上へ
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	// 全行を連結して返す
	return strings.Join(rows, "")
}

func main() {
	result1 := convert("PAYPALISHIRING", 3)
	fmt.Println(result1) // 出力: PAHNAPLSIIGYIR
	result2 := convert("PAYPALISHIRING", 4)
	fmt.Println(result2) // 出力: PINALSIGYAHRPI
	result3 := convert("A", 1)
	fmt.Println(result3) // 出力: A
}
