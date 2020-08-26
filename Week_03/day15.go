package Week_03

import (
	. "homework/Mstruct"
)

func inorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var result []int

	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left //取左子树,进行到底
		} else {
			top := stack[len(stack)-1]       //栈顶
			result = append(result, top.Val) //出栈
			root = top.Right                 //右子树
			stack = stack[:len(stack)-1]     //出栈
		}
	}
	return result
}
