package main

import "fmt"

type Node[V comparable] struct {
	Val  V
	Next *Node[V]
}

type LinkedList[V comparable] struct {
	Head *Node[V]
}

func (l *LinkedList[V]) Append(val V) {
	newNode := &Node[V]{Val: val}

	// If the head is nil, set the new node as the head.
	if l.Head == nil {
		l.Head = newNode
		return
	}

	// Traverse to the end of the list and append the new node
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (l *LinkedList[V]) Insert(val V) {
	newNode := &Node[V]{Val: val}
	newNode.Next = l.Head
	l.Head = newNode
}

func (l *LinkedList[V]) Remove(val V) {
	if l.Head == nil {
		return
	}

	// If the head needs to be removed
	if l.Head.Val == val {
		l.Head = l.Head.Next
		return
	}

	current := l.Head
	for current.Next != nil && current.Next.Val != val {
		current = current.Next
	}

	// If the node was found, remove it
	if current.Next != nil {
		current.Next = current.Next.Next
	}
}

func (l *LinkedList[V]) Reverse() {
	// To set nil at the end, prev needs to be initialized as nil.
	var prev *Node[V]
	current := l.Head

	// Change each one to point backward.
	// In each iteration, only one element is reversed to point backward.
	for current != nil {
		// Store next node
		nextTemp := current.Next
		// Reverse the link
		current.Next = prev
		// Move prev and current one step forward
		prev = current
		current = nextTemp
	}

	// Update head to point to the last node (which is now first)
	l.Head = prev

	// 元: A → B → C → D
	//     ↑
	//   current
	// nextTemp = B  (Bを一時保存)
	// A.Next = nil  (Aの矢印をnilに)
	// prev = A      (Aをprevに)
	// current = B   (Bに移動)
	// ------------------------------------
	//
	// 元: A ← B → C → D
	// 	      ↑
	//       current
	// nextTemp = C  (Cを一時保存)
	// B.Next = A    (Bの矢印をAに)
	// prev = B      (Bをprevに)
	// current = C   (Cに移動)

	//元のリスト: 1 -> 2 -> 3 -> nil
	//
	// イテレーション 1:
	// current = 1, prev = nil
	// nextTemp = 2
	// 1.Next = nil (1 -> nil)
	// prev = 1, current = 2
	//
	// イテレーション 2:
	// current = 2, prev = 1
	// nextTemp = 3
	// 2.Next = 1 (2 -> 1 -> nil)
	// prev = 2, current = 3
	//
	// イテレーション 3:
	// current = 3, prev = 2
	// nextTemp = nil
	// 3.Next = 2 (3 -> 2 -> 1 -> nil)
	// prev = 3, current = nil
	//
	// 最終結果:
	// l.Head = 3
	// リスト: 3 -> 2 -> 1 -> nil
}

func (l *LinkedList[V]) ReverseRecursive() {
	l.Head = l.reverseRecursiveHelper(l.Head)
}

func (l *LinkedList[V]) reverseRecursiveHelper(node *Node[V]) *Node[V] {
	// Base case: if node is nil or we've reached the last node
	if node == nil || node.Next == nil {
		return node
	}

	// Recursively reverse the rest of the list
	rest := l.reverseRecursiveHelper(node.Next)

	// Reverse the link
	node.Next.Next = node
	node.Next = nil

	// Return the new head (which was the last node)
	return rest
}

func (l *LinkedList[V]) Print() {
	current := l.Head
	for current != nil {
		fmt.Print(current.Val, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	list := &LinkedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Insert(0)
	list.Print()
	list.Remove(2)
	list.Print()
	list.Reverse()
	list.Print()
	list.ReverseRecursive()
	list.Print()
}
