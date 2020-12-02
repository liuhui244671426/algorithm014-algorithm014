package Week_16

//63. 不同路径 II
/*基于62题，有了障碍物，62的dp方程组 dp[i][j] = dp[i+1][j] + dp[i][j+1]
到障碍物的数量是0
*/
func UniquePathsWithObstacles() int {
	var obstacleGrid [][]int = [][]int{
		{0,0,0},
		{0,1,0},
		{0,0,0},
	}
	// obstacleGrid = [][]int{
	// 	{0,1},
	// 	{0,0},
	// }
	return uniquePathsWithObstacles(obstacleGrid)
}



func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 空数组不用计算
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	//(0,0)是障碍物。。。
	if m < 1 || n < 1 || 1 == obstacleGrid[0][0] {
        return 0
    }
	// var dp [][]int = make([][]int, m)//空间优化，可以将这个变量去掉
	for i := 0; i < m; i++ {
		// dp[i] = make([]int, n)
		for j := 0; j < n; j++ {			
			if i==0 &&j==0{
				obstacleGrid[i][j] = 1
			} else {
				//遇到障碍物，设置0，然后跳开
				if obstacleGrid[i][j] == 1 {
					obstacleGrid[i][j] = 0
					continue	
				}
				if i ==0 && j != 0 {
					obstacleGrid[i][j] = obstacleGrid[i][j-1]//复制j-1的值
				} else if i!=0 &&j==0{
					obstacleGrid[i][j] = obstacleGrid[i-1][j]//复制i-1的值
				} else {
					obstacleGrid[i][j] = obstacleGrid[i][j-1] + obstacleGrid[i-1][j]
				}

			}
		}
	}
	// fmt.Println(dp)
	return obstacleGrid[m-1][n-1]
}