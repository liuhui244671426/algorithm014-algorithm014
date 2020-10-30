package Week_11

//463. 岛屿的周长

func IslandPerimeter() int {
	var grid [][]int
	grid = [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	return islandPerimeter(grid)
}

func islandPerimeter(grid [][]int) int {
	var lenRow int = len(grid) - 1
	var lenCol int = len(grid[0]) - 1
	var cnt int = 0
	for i := 0; i <= lenRow; i++ {
		for j := 0; j <= lenCol; j++ {
			if grid[i][j] == 1 {
				//fmt.Println("start", cnt, i, j)
				cnt += 4                        //先加 4 个边
				if j > 0 && grid[i][j-1] == 1 { //left
					// fmt.Println("left", grid[i][j-1])
					cnt--
				}
				if j+1 <= lenCol && grid[i][j+1] == 1 { //right
					// fmt.Println("right", grid[i][j+1])
					cnt--
				}
				if i > 0 && grid[i-1][j] == 1 { //up
					// fmt.Println("up", grid[i-1][j])
					cnt--
				}
				if i+1 <= lenRow && grid[i+1][j] == 1 { //down
					// fmt.Println("down", grid[i+1][j])
					cnt--
				}
				//fmt.Println("end cnt", cnt)
			}
		}
	}
	return cnt
}
