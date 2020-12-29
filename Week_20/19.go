package Week_20

//19. 删除链表的倒数第N个节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd() *ListNode {
	var list *ListNode = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	return removeNthFromEnd(list, 2)
}

//快慢指针
//fast和 slow 相差 n 的距离
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	fast, slow := head, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}
	slow.Next = slow.Next.Next //删除 n 位
	return dummy.Next
}
