package Week_07

//617. 合并二叉树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func MergeTrees() *TreeNode {
	var t1 *TreeNode
	var t2 *TreeNode
	t1 = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	t2 = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	return mergeTrees(t1, t2)
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	} else if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}
