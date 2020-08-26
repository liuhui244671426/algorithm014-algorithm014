package Week_01

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

//

// 示例：

// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func main() {
// l1 := &ListNode{
// 	Val: 1,
// 	Next: &ListNode{
// 		Val: 2,
// 		Next: &ListNode{
// 			Val:  4,
// 			Next: nil,
// 		},
// 	},
// }
// l2 := &ListNode{
// 	Val: 1,
// 	Next: &ListNode{
// 		Val: 3,
// 		Next: &ListNode{
// 			Val:  4,
// 			Next: nil,
// 		},
// 	},
// }

// mergeTwoLists(l1, l2)

//nums := []int{2, 3, 4, 5, 18, 17, 6}
//nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
//ret := maxArea(nums)
//fmt.Println(ret)

// 	b := intersect([]int{1, 2, 2, 1}, []int{2, 2})
// 	fmt.Println(b)
// }

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	node := l1
// 	for node.Next != nil {
// 		fmt.Println(node.Val)
// 		node = node.Next
// 	}
// 	return nil
// }
