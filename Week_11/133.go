package Week_11

//113. 路径总和 II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSum() [][]int {
	var root *TreeNode = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	var sum int = 22
	return pathSum(root, sum)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pathSum(root *TreeNode, sum int) [][]int {
	var result [][]int
	dfs(&result, root, 0, sum, []int{})
	return result
}

func dfs(result *[][]int, root *TreeNode, path int, target int, pathNums []int) {
	if root == nil {
		return
	}
	pathNums = append(pathNums, root.Val)
	path += root.Val

	if root.Left == nil && root.Right == nil && path == target {
		//fmt.Println("pathNums", pathNums)
		dst := make([]int, len(pathNums))
		copy(dst, pathNums) //这里要 copy 不然会出错哦
		*result = append(*result, dst)
		//fmt.Println("path end", *result)
		return
	}

	dfs(result, root.Left, path, target, pathNums)
	dfs(result, root.Right, path, target, pathNums)
	return
}
