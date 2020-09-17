package Week_06

//不同路径
//只能向右和向下
func UniquePaths() int {
	var m int = 7
	var n int = 3
	return uniquePaths(m, n)
}

//动态规划
func uniquePaths(m int, n int) int {
	var dp [][]int = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	for y := 0; y < n; y++ { //y
		for x := 0; x < m; x++ { //x
			if x == 0 || y == 0 {
				dp[y][x] = 1
			} else {
				dp[y][x] = dp[y-1][x] + dp[y][x-1] //状态转移方程 f(x,y) = f(x-1,y)+f(x,y-1)
			}
		}
	}
	return dp[n-1][m-1]
}
