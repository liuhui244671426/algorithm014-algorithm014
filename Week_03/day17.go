package Week_03
//https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
//从尾到头打印链表

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
func reversePrint(head *ListNode) []int {
	var stack []int
	var reverseStack []int
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	for i := len(stack) - 1; i >= 0; i-- {
		reverseStack = append(reverseStack, stack[i])
	}
	return reverseStack
}

func ReversePrint(head *ListNode) []int {
	return reversePrint(head)
}
