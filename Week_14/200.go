package Week_14

//200. 岛屿数量

func NumIslands() int {
	var grid [][]byte = [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	return numIslands(grid)
}

//思路:把岛变成海,同时 count++
//把岛变海,使用DFS
func numIslands(grid [][]byte) int {
	cnt := 0
	if len(grid) == 0 {
		return cnt
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				cnt++
				modifyLands(grid, i, j)
			}
		}
	}
	return cnt
}

func modifyLands(grid [][]byte, row, col int) {
	//不能越界
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || grid[row][col] == '0' {
		return
	}
	grid[row][col] = '0' //把岛,都变成海
	modifyLands(grid, row+1, col)
	modifyLands(grid, row-1, col)
	modifyLands(grid, row, col+1)
	modifyLands(grid, row, col-1)
}
