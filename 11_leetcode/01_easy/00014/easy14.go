package main

import "fmt"

// こっちのほうが早い：めっちゃシンプル
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 最初の文字列を基準にする
	prefix := strs[0]

	// 起点となる文字列の文字数ループ
	for i := 0; i < len(prefix); i++ {
		char := prefix[i]

		// 残りの文字列を比較
		for _, str := range strs[1:] {
			// 文字列の長さが足りないか、文字が一致しない場合
			// 終了条件を調べている、直前の文字までは一致していることを保証済み
			if i >= len(str) || str[i] != char {
				// 明示的に空文字は返さない（一致しない場合はここで勝手に空文字になる）
				return prefix[:i]
			}
		}
	}
	return prefix
}

func main() {
	tests := [][]string{
		{"flower", "flow", "flight"},
		{"flower", "flow", "flo"},
		{"flo", "flow", "flower"},
		{"flow", "flower", "flight"},
		{"dog", "racecar", "car"},
		{"a"},
		{"acb", "b", "", ""},
		{},
	}
	for _, test := range tests {
		fmt.Println(longestCommonPrefix(test))
	}
}

// 削っていく方式
// func answer1_longestCommonPrefix(strs []string) string {
// 	if len(strs) == 0 {
// 		return ""
// 	}

// 	// 最初の文字列をprefixとして設定
// 	prefix := strs[0]
// 	// 最初の文字列を起点にしているので、2番目の文字列からループを回す
// 	fmt.Println("=====================================")
// 	fmt.Println("prefix:", prefix)
// 	for _, str := range strs[1:] {

// 		fmt.Println()
// 		fmt.Println("str:", str)

// 		// 以下の場合にループを続行
// 		// ① 対象文字列が最初の文字列よりも大きい：IndexOutOfRangeを避けるため
// 		//    対象文字列のほうが最初の文字より
// 		// or
// 		// ② 対象文字列の先頭が最初の文字列と一致しない場合：こっちが大事
// 		for len(str) < len(prefix) || str[:len(prefix)] != prefix {
// 			// 最初の文字列の末尾を削除していく
// 			prefix = prefix[:len(prefix)-1]
// 			fmt.Println("prefix:", prefix)
// 			// 最終的に最初の文字列が空文字になったら、共通のprefixは存在しない
// 			if len(prefix) == 0 {
// 				return ""
// 			}
// 		}

// 	}
// 	return prefix
// }
