package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var head *ListNode

	head = appendNode(head, 1)
	printList(head)

	head = appendNode(head, 2)
	printList(head)

	head = appendNode(head, 3)
	printList(head)

	head = deleteNode(head, 2)
	printList(head)

	node := searchNode(head, 2)
	if node != nil {
		fmt.Println("Found:", node.Val) // Found: 2
	} else {
		fmt.Println("Not Found")
	}

}

func appendNode(head *ListNode, val int) *ListNode {
	newNode := &ListNode{Val: val}
	if head == nil {
		return newNode
	}

	// head：HeadNodeのアドレス
	// ① HeadNode：{Val：0, Next：Node1のアドレス}
	// ② Node1   ：{Val：1, Next：nil}

	current := head
	for current.Next != nil {
		current = current.Next
	}
	// この時点でcurrent = Node1のアドレス

	current.Next = newNode

	// HeadNode：{Val：0, Next：Node1のアドレス}
	// Node1   ：{Val：1, Node2のアドレス}
	// Node2   ：{Val：2, nil} // newNode

	return head // リンクなので返すのはヘッド
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil // リストが空の場合
	}
	if head.Val == val {
		return head.Next // 先頭のノードが削除対象の場合
	}

	// HeadNode：{Val：0, Next：Node1のアドレス}
	// Node1   ：{Val：1, Node2のアドレス}
	// Node2   ：{Val：2, Node3のアドレス} // 削除対象
	// Node3   ：{Val：3, nil}
	current := head
	for current.Next != nil && current.Next.Val != val { // currentの次のノードを比較
		current = current.Next // 削除対象のノードを探す
	}
	// current = Node1であり、削除対象がNode2でNextで構えている
	if current.Next != nil {
		// currentのNextに次の次のノードを設定する
		// Node1.Next = Node3のアドレス
		current.Next = current.Next.Next // 削除対象のノードをスキップ
	}
	return head // リストの先頭を返す
}

func searchNode(head *ListNode, val int) *ListNode {
	current := head
	for current != nil {
		if current.Val == val {
			return current // ノードが見つかった場合
		}
		current = current.Next // 次のノードに移動
	}
	return nil // ノードが見つからなかった場合
}

func printList(head *ListNode) {
	current := head
	for current != nil { // アドレスがあるということは値がある
		fmt.Print(current.Val, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}
