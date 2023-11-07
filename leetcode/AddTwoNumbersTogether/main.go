package main

import (
	"fmt"
)

// 两数相加
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		//fmt.Printf("sum is: %d, carry is: %d\n", sum, carry)
		//fmt.Println(head)
		if head == nil {
			head = &ListNode{Val: sum}
			//fmt.Printf("head.Val: %d\n", head.Val)
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
		//fmt.Printf("-------------\n")
		//fmt.Println(l1)
		//fmt.Printf("-------------\n")
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

func main() {
	// 创建两个链表
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	// 调用相加函数
	result := addTwoNumbers(l1, l2)

	// 打印结果链表
	/*
		这个循环的目的是遍历 result 链表中的每个节点，依次输出它们的值。
		在每次循环迭代中，我们首先打印当前节点的值 (result.Val)，然后将 result 更新为下一个节点
		(result = result.Next)，以便进行下一次循环迭代。这个过程一直持续到链表的末尾，即 result 变成 nil。
		通过这种方式，我们能够遍历并打印结果链表中的每个元素。这是处理链表数据结构的常见方式。
	*/
	// result = &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}}
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
	fmt.Println()
}
