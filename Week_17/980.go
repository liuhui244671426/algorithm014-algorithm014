package Week_17

//980. 不同路径 III
/*
1 表示起始方格。且只有一个起始方格。
2 表示结束方格，且只有一个结束方格。
0 表示我们可以走过的空方格。
-1 表示我们无法跨越的障碍。
*/
func UniquePathsIII() int {
	var grid [][]int = [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, -1},
	}

	return uniquePathsIII(grid)
}

func uniquePathsIII(grid [][]int) int {
	x, y, step := 0, 0, 1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			//找到开始的起点
			if grid[i][j] == 1 {
				x = j
				y = i
				continue
			}
			//对于空白的格子,记录它的总数
			if grid[i][j] == 0 {
				step++
			}
		}
	}

	return dfsUniquePathIII(x, y, step, grid)
}

func dfsUniquePathIII(x, y, step int, grid [][]int) int {
	//越界条件
	if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) || grid[y][x] == -1 {
		return 0
	}
	//到终点了
	if grid[y][x] == 2 {
		if step == 0 {
			return 1
		} else {
			return 0
		}
	}
	//选择
	grid[y][x] = -1
	res := 0
	//对上下左右 四个方向进行深度优先
	res += dfsUniquePathIII(x-1, y, step-1, grid)
	res += dfsUniquePathIII(x+1, y, step-1, grid)
	res += dfsUniquePathIII(x, y-1, step-1, grid)
	res += dfsUniquePathIII(x, y+1, step-1, grid)
	//撤销
	grid[y][x] = 0
	return res
}
