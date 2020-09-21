package Week_07

//二叉树的层序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder() [][]int {
	var root *TreeNode
	root = &TreeNode{
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

/*
思路:
可通过 BFS 或 前序遍历 实现.其中前序遍历,之前有做过,所以这次实现 BFS
BFS 实现层序遍历的话,需要队列辅助, 由于应用场景简单,不需要引入 list 包
*/
func levelOrder(root *TreeNode) [][]int {

	var queue []*TreeNode
	var res [][]int
	if root == nil {
		return res
	}

	queue = append(queue, root)

	for len(queue) > 0 {
		var tmp []int = []int{}
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:] //pop first

			if node.Left != nil {
				queue = append(queue, node.Left) //push left
			}
			if node.Right != nil {
				queue = append(queue, node.Right) //push right
			}
			tmp = append(tmp, node.Val)
		}
		res = append(res, tmp)
	}

	return res
}
