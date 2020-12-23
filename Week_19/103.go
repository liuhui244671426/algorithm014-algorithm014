package Week_19

//103. 二叉树的锯齿形层序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ZigzagLevelOrder() [][]int {
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
	return zigzagLevelOrder(root)
}

//bfs
/*
优化方法:
1.可以不用翻转,比如不用下面的 reverse 方法,通过 length-i-1 进行翻转
2.可以批量裁剪 queue 的数据,减少内存操作
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := 0

	for len(queue) > 0 {
		length := len(queue)
		tmp := make([]int, length)
		for i := 0; i < length; i++ {
			node := queue[i]
			//queue = queue[1:]
			//tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if level%2 == 0 {
				tmp[i] = node.Val
			} else {
				tmp[length-i-1] = node.Val //按照下标,reverse
			}
		}
		level++
		ans = append(ans, tmp)
		queue = queue[length:] //批量减数据
	}
	return ans
}

func reverse(ints []int) []int {
	l, r := 0, len(ints)-1
	for l < r {
		ints[l], ints[r] = ints[r], ints[l]
		l++
		r--
	}
	return ints
}
