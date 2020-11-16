package Week_14

//102. 二叉树的层序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder() [][]int {
	var root *TreeNode = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	return levelOrder(root)
}

//层序遍历->队列
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		lenQ := len(queue)
		var level []int = []int{}

		for i := 0; i < lenQ; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if level != nil {
			//fmt.Println(level)
			result = append(result, level)
		}

	}
	return result
}
