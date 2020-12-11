package Week_17

//64. 最小路径和

func MinPathSum() int {
	var grid [][]int
	grid = [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	grid = [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	grid = [][]int{
		{1, 2},
		{1, 1},
	}
	return minPathSum(grid)
}

func minPathSum(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 && j != 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j] //第一行,只能从左边 累加
			}
			if j == 0 && i != 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j] //第一列,只能从上边 累加
			}
			if j != 0 && i != 0 {
				//动态转移方程: 从上边和从左边,取一个最小值,然后和本值 累加
				grid[i][j] = min(grid[i][j-1], grid[i-1][j]) + grid[i][j]
			}
		}
	}
	//fmt.Println(grid)
	return grid[rows-1][cols-1]
}
