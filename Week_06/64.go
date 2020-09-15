package Week_06

//最小路径和
func MinPathSum() int {
	var grid [][]int
	grid = [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	grid = [][]int{
		{1, 2, 5},
		{3, 2, 1},
	}

	// grid = [][]int{{0}}
	return minPathSum(grid)
}

//再次基础上可以将空间复杂度再降一降,不申明 dp 变量,直接修改 grid 变量即可
//动态规划
func minPathSum(grid [][]int) int {
	// var dp [][]int = make([][]int, len(grid))
	// for k, v := range grid {
	// 	dp[k] = make([]int, len(v))
	// }
	//fmt.Println(dp)
	for y, v_y := range grid {
		//fmt.Println("y, val", y, v_y)
		for x, _ := range v_y {
			//fmt.Println("x, val", x, v_x, grid[y][x])
			if x == 0 && y == 0 {
				//grid[y][x] = grid[y][x]
			} else if x == 0 {
				grid[y][x] = grid[y][x] + grid[y-1][x]
			} else if y == 0 {
				grid[y][x] = grid[y][x] + grid[y][x-1]
			} else {
				grid[y][x] = min(grid[y-1][x], grid[y][x-1]) + grid[y][x]
			}
		}
	}
	//fmt.Print(dp)
	return grid[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
