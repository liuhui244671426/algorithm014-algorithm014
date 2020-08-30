package Week_03

import "math"

import (
	. "homework/Mstruct"
)
//这是自己本周唯一不看题解完成的题目.

func isValidBST(root *TreeNode) bool {
	return helper_isValidBST(root, math.MinInt64, math.MaxInt64)
}

func helper_isValidBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	//fmt.Println(root)
	if root.Val <= min || root.Val >= max {
		return false
	}
	l := helper_isValidBST(root.Left, min, root.Val)
	r := helper_isValidBST(root.Right, root.Val, max)
	return l && r
}
