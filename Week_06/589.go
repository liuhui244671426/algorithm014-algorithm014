package Week_06

//N叉树的前序遍历

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
type Node struct {
	Val      int
	Children []*Node
}

func Preorder() []int {
	var tree *Node = &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{
						Val:      5,
						Children: []*Node{},
					},
					{
						Val:      6,
						Children: []*Node{},
					},
				},
			},
			{
				Val:      2,
				Children: []*Node{},
			},
			{
				Val:      4,
				Children: []*Node{},
			},
		},
	}
	return preorder(tree)
}

var res []int

func preorder(root *Node) []int {
	res = []int{}
	if root == nil {
		return []int{} //空树,返回空
	}
	helper(root)
	return res
}

func helper(root *Node) {
	//fmt.Println(root.Val, res)
	if root.Children == nil {
		return //没有子节点,返回
	}
	res = append(res, root.Val) //把 val 加入数组
	for _, children := range root.Children {
		helper(children) //循环从左开始递归吧
	}
}
