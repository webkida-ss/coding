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
	var prev *Node[V]
	current := l.Head

	for current != nil {
		// Store next node
		next := current.Next
		// Reverse the link
		current.Next = prev
		// Move prev and current one step forward
		prev = current
		current = next
	}

	// Update head to point to the last node (which is now first)
	l.Head = prev
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
