package Week_14

import "math"

//515. 在每个树行中找最大值

func LargestValues() []int {
	var root *TreeNode = &TreeNode{
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

func largestValues(root *TreeNode) []int {
	var ans []int = make([]int, 0)
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		maxInt := math.MinInt64 //节点值,会有负数
		l := len(queue)
		for i := 0; i < l; i++ {//通过 for 聚拢行的 node
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if maxInt < node.Val {
				maxInt = node.Val
			}
		}
		ans = append(ans, maxInt)
	}
	return ans
}
