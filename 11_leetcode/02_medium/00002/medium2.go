package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{} // ダミーヘッド
	current := dummyHead
	carry := 0 // 上がり

	// どちらかのリストがnilになるまで繰り返す
	for l1 != nil || l2 != nil {
		fmt.Println("=================================")
		fmt.Println("l1,l2がnilになるまで繰り返す")
		fmt.Println("=================================")

		// 前回の上がりを足す
		sum := carry

		if l1 != nil {
			// 上がりにl1の値を足す
			sum += l1.Val
			// l1を次のノードに進める
			l1 = l1.Next
		}
		if l2 != nil {
			// 上がりにl2の値を足す
			sum += l2.Val
			// l2を次のノードに進める
			l2 = l2.Next
		}
		// 上がりを計算（小数点以下切り捨て）
		carry = sum / 10
		fmt.Println("sum:", sum)
		fmt.Println("carry:", carry)

		// 結果連結リストの現在ノードに次ノードのリンクを張る
		current.Next = &ListNode{Val: sum % 10}
		fmt.Println("current.Next:", current.Next)
		// 次のノードを現在ノードにする
		current = current.Next
		fmt.Println("current:", current)
		fmt.Println()
	}

	// ループが終わっても上がりがあれば、最後の上がりノードをを連結リストに追加する
	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	fmt.Println("☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆")
	fmt.Println(dummyHead.Next)
	fmt.Println("☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆")
	return dummyHead.Next
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
