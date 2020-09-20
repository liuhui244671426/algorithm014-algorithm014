package Week_06

//左叶子之和
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SumOfLeftLeaves() int {
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
	root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	return sumOfLeftLeaves(root)
}

func sumOfLeftLeaves(root *TreeNode) int {
	var res int = 0
	if root == nil {
		return 0
	}
	var isLeft bool = false //当前节点是否是 左节点
	if root.Left != nil {
		isLeft = true
	}
	helper2(root, isLeft, &res)
	return res
}

func helper2(root *TreeNode, isLeft bool, res *int) {
	//fmt.Println("val", root.Val)
	if isLeft == true && root.Left == nil && root.Right == nil {
		*res += root.Val
		return
	}
	if root.Left != nil {
		helper2(root.Left, true, res)
	}
	if root.Right != nil {
		helper2(root.Right, false, res)
	}

}
