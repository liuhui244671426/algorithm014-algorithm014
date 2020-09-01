package Week_04

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

var res [][]int

func levelOrder(root *TreeNode) [][]int {

	res = make([][]int, 0)
	preOrder(0, root)
	return res
}
//前序遍历
func preOrder(level int, root *TreeNode) {
	if root == nil {
		return
	}
	if level == len(res) {
		res = append(res, []int{}) //先设置一个[]
	}
	//fmt.Println(result, level, root.Val)
	//(*result)[level] = append((*result)[level], root.Val) //这是引用变量时,使用的方法,第二版已经把引用变量的去掉
	res[level] = append(res[level], root.Val)
	preOrder(level+1, root.Left)
	preOrder(level+1, root.Right)
}
