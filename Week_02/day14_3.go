package Week_02

//https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import (
	. "homework/Mstruct"
)

//左-根-右
//中序遍历,左链入栈
func inorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var result []int

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top_idx := len(stack) - 1
		result = append(result, stack[top_idx].Val)
		root = stack[top_idx].Right
		stack = stack[:top_idx]
	}
	return result
}
