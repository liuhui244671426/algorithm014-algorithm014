package Week_04

import (
	"fmt"
	"math"
)

//在每个树行中找最大值

func LargestValues() []int {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	return largestValues(root)
}

//用队列 实现 层级遍历
func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root) //push

	for len(stack) > 0 {
		max := math.MinInt32 //题目需要
		length := len(stack)

		for i := 0; i < length; i++ {
			node := stack[0]
			stack = stack[1:] //pop

			if node.Val > max {
				fmt.Println("node max", node.Val, max)
				max = node.Val
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		result = append(result, max)
	}
	return result
}
