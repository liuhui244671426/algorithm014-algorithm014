package Week_07

//200. 岛屿数量
//思路:连通性->并查集
func NumIslands() int {
	var grid [][]byte
	grid = [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	// grid = [][]byte{
	// 	{'1', '1', '0', '0', '0'},
	// 	{'1', '1', '0', '0', '0'},
	// 	{'0', '0', '1', '0', '0'},
	// 	{'0', '0', '0', '1', '1'},
	// }
	return numIslands(grid)
}

func numIslands(grid [][]byte) int {
	// for _, v1 := range grid {
	// 	fmt.Println(v1)
	// }

	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	dummy := 0
	var uf *UnionFind
	uf = uf.New(row * col)
	// fmt.Println("row, col", row, col)
	for y, yv := range grid {
		for x, xv := range yv {
			if xv == '0' {
				dummy++
			} else {
				//向下
				if y+1 < row && grid[y+1][x] == '1' {
					uf.Union(index(col, x, y), index(col, x, y+1)) //把向下的元素合并
					//grid[y+1][x] = '2'
				}
				//向右
				if x+1 < col && grid[y][x+1] == '1' {
					uf.Union(index(col, x, y), index(col, x+1, y)) //把向右的元素合并
					//grid[y][x+1] = '2'
				}
			}
		}
	}
	// fmt.Println("union end", uf.count, dummy)
	// for _, v1 := range grid {
	// 	fmt.Println(v1)
	// }
	return uf.Count() - dummy
}

//把二维变成一维
func index(col, x, y int) int {
	return y*col + x
}
