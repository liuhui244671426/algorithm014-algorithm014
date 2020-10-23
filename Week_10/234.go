package Week_10

//234. 回文链表

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func IsPalindrome() bool {
	var head *ListNode
	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
				},
			},
		},
	}
	return isPalindrome(head)
}
/*
1.栈:效率一般.
2.双指针:
*/ 
func isPalindrome(head *ListNode) bool {
	var stack []*ListNode
	node := head
	for node != nil {
		//fmt.Println(node.Val)
		stack = append(stack, node)
		node = node.Next
	}
	for head != nil {
		if head.Val != stack[len(stack)-1].Val {
			return false
		}
		head = head.Next
		stack = stack[:len(stack)-1]
	}
	return true
}
