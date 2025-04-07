package main

import "fmt"

func generateParenthesis(n int) []string {
	var res []string
	dfs(n, n, "", &res)
	return res
}

func dfs(left, right int, cur string, res *[]string) {

	fmt.Println("left:", left)
	fmt.Println("right:", right)
	fmt.Println("cur:", cur)
	fmt.Println("res:", res)
	fmt.Println()

	if left == 0 && right == 0 {
		*res = append(*res, cur)
		return
	}
	if left > 0 {
		dfs(left-1, right, cur+"(", res)
	}
	if right > left {
		dfs(left, right-1, cur+")", res)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
