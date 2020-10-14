package Week_09

import "fmt"

//63. 不同路径 II

func UniquePathsWithObstacles() int {
	var obstacleGrid [][]int
	obstacleGrid = [][]int{
		{1, 0},
	}
	obstacleGrid = [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	obstacleGrid = [][]int{
		{0, 0}, {1, 1}, {0, 0},
	}
	fmt.Println("init", obstacleGrid)
	return uniquePathsWithObstacles(obstacleGrid)
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var rows int = len(obstacleGrid)
	var cols int = len(obstacleGrid[0])
	//出口就是障碍物...
	if rows == 0 || cols == 0 || obstacleGrid[0][0] == 1{
		return 0
	}
	//初始化二维数组
	var dp [][]int = make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}
	//将第一行和第一列,无障碍物的格子,设置成 1
	for i := 0; i < rows && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < cols && obstacleGrid[0][i] == 0; i++ {
		dp[0][i] = 1
	}

	//开始
	for row := 1; row < rows; row++ {
		for col := 1; col < cols; col++ {
			if obstacleGrid[row][col] == 0 {
				dp[row][col] = dp[row-1][col] + dp[row][col-1]
			}
		}
	}
	//fmt.Println("result dp", dp)
	return dp[rows-1][cols-1]
}
