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
	if (deep+1)%k == 0 || deep < k {
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
	for i := 8; i > 1; i-- {
		node.add(NewNode(i))
	}
	fmt.Println(node)
	node, _ = reverse(node, 1, 3)
	fmt.Println(node)
}

/*
这其实是⼀道变形的链表反转题，⼤致描述如下 给定⼀个单链表的头节点 head,实现⼀
个调整单链表的函数，使得每K个节点之间为⼀组进⾏逆序，并且从链表的尾部开始组
起，头部剩余节点数量不够⼀组的不需要逆序。（不能使⽤队列或者栈作为辅助）
例如：
链表: 1->2->3->4->5->6->7->8->null, K = 3 。那么 6->7->8 ， 3->4->5 ， 1->2 各
位⼀组。调整后： 1->2->5->4->3->8->7->6->null 。其中 1，2不调整，因为不够⼀组。
解析
原⽂：https://juejin.im/post/5d4f76325188253b49244dd0
*/
