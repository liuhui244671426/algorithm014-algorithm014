package Week_09

import "fmt"

//64. 最小路径和

func MinPathSum() int {
	var grid [][]int
	grid = [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	fmt.Println("init", grid)
	return minPathSum2(grid)
}

func minPathSum(grid [][]int) int {
	var rows = len(grid)
	var cols = len(grid[0])
	if rows == 0 || cols == 0 {
		return 0
	}

	var dp [][]int = make([][]int, rows) //可以优化
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}

	dp[0][0] = grid[0][0]
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 && j != 0 { //第一行,只能向右走,grid[i][j] + grid[i][j-1]
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 && i != 0 { //第一列,只能向下走,dp[i-1][j] + grid[i][j]
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else if j > 0 && i > 0 {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
			}
		}
	}
	fmt.Println(dp)
	return dp[rows-1][cols-1]
}

/*
minPathSum-->minPathSum2

minPathSum 没有空间复杂度优势
minPathSum2 已压缩空间,取消原来的 dp 二维数组
*/
func minPathSum2(grid [][]int) int {
	var rows = len(grid)
	var cols = len(grid[0])
	if rows == 0 || cols == 0 {
		return 0
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 && j != 0 { //第一行,只能向右走,grid[i][j] + grid[i][j-1]
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 && i != 0 { //第一列,只能向下走,dp[i-1][j] + grid[i][j]
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else if j > 0 && i > 0 {
				grid[i][j] = min(grid[i][j-1], grid[i-1][j]) + grid[i][j]
			}
		}
	}
	//fmt.Println(grid)
	return grid[rows-1][cols-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
