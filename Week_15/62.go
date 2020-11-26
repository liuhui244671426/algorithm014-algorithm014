package Week_15

//62. 不同路径

func UniquePaths() int {
	var m int = 3
	var n int = 2
	return uniquePaths(m, n)
}

/*
为什么要用动态规划？
符合什么条件可以使用动态规划？
1.从题里如果找到“最”字，80%可以动态规划解决
2.无后效性
3.重复字问题
4.最佳字结构

*/
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	//fmt.Println(dp)
	return dp[m-1][n-1]
}
