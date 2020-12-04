package Week_16

//120. 三角形最小路径和

func MinimumTotal() int {
	var triangle [][]int = [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	return minimumTotal(triangle)
}

/*
1.bfs
2.dfs
3.dp 特征：1.有最字 2.无后效性 3.重复子结构
*/
func minimumTotal(triangle [][]int) int {
	if len(triangle) < 1 {
		return 0
	}
	//从倒数第 2 层推,太秒了.
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] = min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
		}
	}
	return triangle[0][0]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
