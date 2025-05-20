package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	// ダミーノードを作成し、リンクリストの先頭を指すようにします。
	dummy := &ListNode{Next: head}
	// prevは、交換するノードペアの前のノードを追跡します。
	prev := dummy
	// リンクリストの終わりに達するか、次のノードがないノードに達するまでループします。
	for head != nil && head.Next != nil {
		// 交換するノードペアを取得します。
		first := head
		second := head.Next

		// prevの次のノードをsecondにリンクさせます。
		prev.Next = second
		// firstの次のノードをsecondの次のノードにリンクさせます。
		first.Next = second.Next
		// secondの次のノードをfirstにリンクさせます。
		second.Next = first

		// prevとheadを更新します。
		prev = first
		head = first.Next
	}
	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Create a linked list: 1 -> 2 -> 3 -> 4
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}

	// Print the original list
	fmt.Println("Original list:")
	printList(head)

	head = swapPairs(head)
	fmt.Println("After swapping pairs:")
	printList(head)

}
