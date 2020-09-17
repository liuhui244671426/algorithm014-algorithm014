package Week_06

//不同路径 II

func UniquePathsWithObstacles() int {
	var obstacleGrid [][]int
	obstacleGrid = [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	// obstacleGrid = [][]int{
	// 	{1, 0},
	// }
	return uniquePathsWithObstacles(obstacleGrid)
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	len_x := len(obstacleGrid[0])
	len_y := len(obstacleGrid)
	if len_y == 0 || len_x == 0 {
		return 0
	}
	var dp [][]int = make([][]int, len_y)
	for i, _ := range obstacleGrid {
		dp[i] = make([]int, len_x)
	}
	//fmt.Println(dp)
	for y, yV := range obstacleGrid {
		for x, xV := range yV {
			//fmt.Println("x,y", x, y, xV)
			if xV == 0 {
				if x == 0 && y == 0 {
					dp[y][x] = 1 //初始化 f(0,0)=1
				} else if x == 0 {
					dp[y][x] = dp[y-1][x]
				} else if y == 0 {
					dp[y][x] = dp[y][x-1]
				} else {
					dp[y][x] = dp[y][x-1] + dp[y-1][x]
				}
			} else {
				//跳过 障碍物
			}

		}
	}
	//fmt.Println(dp)
	return dp[len_y-1][len_x-1]
}
