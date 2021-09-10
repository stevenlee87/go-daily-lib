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

// node, _ = reverse(node, 1, 3)
func reverse(node *Node, deep, k int) (*Node, *Node) {
	if node == nil {
		return nil, nil
	}
	head, last := reverse(node.Next, deep+1, k)
	fmt.Printf("head: %v, last: %v\n", head, last)
	if head == nil && last == nil {
		node.Next = nil
		fmt.Println("node.Value -->", node.Value)
		fmt.Println("node :", node)
		return node, node
	}
	if (deep)%k == 0 {
		node.Next = head
		return node, node
	}
	node.Next = last.Next
	last.Next = node
	return head, node
}

func main() {
	node := NewNode(1)
	//fmt.Println(node.Value)
	for i := 9; i > 1; i-- {
		node.add(NewNode(i))
	}
	fmt.Println(node)
	node, _ = reverse(node, 1, 3)
	fmt.Println(node)
}
