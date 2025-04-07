package main

import "fmt"

func moveDisk(from, to string, disk int) {
	fmt.Printf("ディスク%dを %s から %s へ移動\n", disk, from, to)
}

// n:　上に重ねるディスクの数
func solveHanoi(n int, from, to, auxiliary string) {

	// 基底ケース
	if n == 1 {
		moveDisk(from, to, n)
		return
	}

	// ① n-1個のディスクを補助の棒に移動
	solveHanoi(n-1, from, auxiliary, to)

	// ② 残りの1つのディスクを目的の棒に移動
	moveDisk(from, to, n)

	// ③ n-1個のディスクを補助の棒から目的の棒に移動
	solveHanoi(n-1, auxiliary, to, from)
}

func main() {
	fmt.Println("ケース1")
	solveHanoi(1, "from", "to", "auxiliary")
	fmt.Println()
	fmt.Println("ケース2")
	solveHanoi(2, "FROM", "TO", "AUXILIARY")
	fmt.Println()
	fmt.Println("ケース3")
	solveHanoi(3, "A", "C", "B")
	fmt.Println()
	fmt.Println("ケース4")
	solveHanoi(4, "移動元", "移動先", "補助")	
}