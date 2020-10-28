package Week_11

//144. 二叉树的前序遍历

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func PreorderTraversal() []int {
	var root *TreeNode
	root = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	return preorderTraversal(root)
}

/*
BFS借助栈完成前序遍历
需要先入栈右子树,左子树才能先遍历
*/
func preorderTraversal(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}
	var stack []*TreeNode
	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

	}
	return result
}
