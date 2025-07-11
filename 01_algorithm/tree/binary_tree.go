package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert inserts a value into the binary tree.
func (n *Node) Insert(value int) {
	if value <= n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

// InOrderTraversal performs an in-order traversal of the binary tree.
func (n *Node) InOrderTraversal() {
	if n == nil {
		return
	}
	n.Left.InOrderTraversal()
	fmt.Print(n.Value, " ")
	n.Right.InOrderTraversal()
}

func main() {
	root := &Node{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)

	fmt.Print("In-order traversal: ")
	root.InOrderTraversal()
}
