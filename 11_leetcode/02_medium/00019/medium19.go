package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := dummy, dummy

	// firstを先にn分進める
	for i := 1; i <= n+1; i++ {
		first = first.Next
	}

	// firstが終端に来たら終了：n分先に進んでいるからsecond後ろから2になる
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// リンクドリストからn番目のノードを削除しています。
	second.Next = second.Next.Next
	return dummy.Next
}

func main() {
	// リンクドリストの作成
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}

	head = removeNthFromEnd(head, 2)
	for node := head; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}
