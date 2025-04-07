package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// ダミーヘッドを使って結果リストを初期化
	dummy := &ListNode{} // これはダミーノード！リストではない
	current := dummy

	fmt.Println("これからループを開始します")
	fmt.Println()

	// l1とl2が両方ともnullでない限りループ
	for l1 != nil && l2 != nil {

		fmt.Println("loop毎のログ------------------")
		fmt.Println(l1.Val, l2.Val)

		// l1とl2のどちらが小さいかを比較し、小さい方を結果リストに追加
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next // l1の次のノードをl1に代入
		} else {
			current.Next = l2
			l2 = l2.Next // l2の次のノードをl2に代入
		}
		current = current.Next
		fmt.Println(current)
	}

	// l1またはl2のいずれかがまだ要素を持っていれば、それを結果リストに追加
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, " -> ")
		node = node.Next
	}
	fmt.Println("nil")
}

func main() {
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	result := mergeTwoLists(list1, list2)
	printList(result)
}
