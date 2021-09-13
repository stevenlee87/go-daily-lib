package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func (n *Node) add(node *Node) {
	node.Next = n.Next
	n.Next = node
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("%d -> ", n.Value) + n.Next.String()
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
		Next:  nil,
	}
}

func main() {
	node := NewNode(1)
	fmt.Println(node)
	if node.Next == nil {
		fmt.Println("node is nil")
	}
}
