package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// RemoveDuplicationNode remove duplicate nodes from linked list
// Example:
//
// Input: 1 --> 2 --> 2 --> 5 --> 3 --> 4 --> 5
// Output: 1 --> 3 --> 4
func RemoveDuplicationNode(root *ListNode) *ListNode {
	if root == nil {
		return root
	}

	m := make(map[int]int, 0)
	listNode := root
	for listNode != nil {
		m[listNode.Val]++
		if listNode.Next != nil && m[listNode.Next.Val] > 1 {
			if listNode.Next.Next != nil {
				listNode.Next = listNode.Next.Next
			} else {
				listNode.Next = nil
			}
		}
		listNode = listNode.Next
	}

	return listNode
}

func main() {
	RemoveDuplicationNode()
}
