/** * Definition for singly-linked list.

* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func deleteDuplicates(head *ListNode) *ListNode {
	h:=&ListNode{}
	h.Next=head
	cur:=head
	pre:=h
	for cur!=nil{
		flag:=false
		for cur.Next!=nil && cur.Val==cur.Next.Val{
			cur=cur.Next
			flag=true
		}

		if flag{
			pre.Next=cur.Next
		}else{
			pre.Next=cur
			pre=cur
		}
		cur=cur.Next
	}
	return h.Next
}