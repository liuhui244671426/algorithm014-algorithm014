package Week_02

import (
	. "homework/Mstruct"
)

//关键点1：使用栈来模拟遍历的顺序，根据栈的先入后出的特点，入栈的顺序跟遍历的顺序相反，这样出栈的时候就可以得到想要的顺序
//关键点2：每一次循环的入栈按照顺序完成之后，在顶层加一个标志位nil,下一个循环如果碰到标识位，则把标识位的下一个节点弹出，这样就能保证弹出的顺序是：中左右

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	// 定义stack模拟二叉树，利用先入后出的特点，出栈顺序为中左右，进栈顺序则为右左中
	var stack []*TreeNode
	// 定义res最终返回的结果
	var res []int
	// 先把root压入栈中
	stack = append(stack, root)
	for len(stack) > 0 {
		// 获取栈的顶层节点
		top := stack[len(stack)-1]
		// 弹出栈的顶层节点
		stack = stack[:len(stack)-1]
		// 如果栈的顶层节点不为空，则入栈右左中的顺序
		// 最后一个入栈的同时压入nil标识位
		// 当top为空时，则取出top的下一个节点
		if top != nil {
			if top.Right != nil {
				stack = append(stack, top.Right)
			}
			if top.Left != nil {
				stack = append(stack, top.Left)
			}
			stack = append(stack, top, nil)
		} else if len(stack) > 0 {
			top = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, top.Val)
		}
	}
	return res
}

func PreorderTraversal(root *TreeNode) []int {
	return preorderTraversal3(root)
}

//#根->左->右
//#栈 先进后出
func preorderTraversal2(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int
	if root == nil {
		return []int{}
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]   //取栈顶
		stack = stack[:len(stack)-1] //
		res = append(res, top.Val)
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return res
}

//
func preorderTraversal3(root *TreeNode) []int {
	var result []int = []int{}
	helper(root, &result)
	return result
}

func helper(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	helper(root.Left, result)
	helper(root.Right, result)

}


