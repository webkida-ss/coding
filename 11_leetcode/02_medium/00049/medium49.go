package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {

	// まずはキー（文字列の昇順）でマップ作成
	m := make(map[string][]string) // キーに対する文字列リスト
	for _, str := range strs {
		// 1文字ずつに分割 → ソート → 結合
		s := strings.Split(str, "")
		sort.Strings(s)
		sortedStr := strings.Join(s, "")
		m[sortedStr] = append(m[sortedStr], str)
	}

	// キーに対するバリューを取り出して2次元配列に変換する
	var result [][]string
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
