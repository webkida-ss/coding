package main

import "fmt"

type Node[V any] struct {
	Val  V
	Next *Node[V]
}

type LinkedList[V any] struct {
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
	list.Print()

}
