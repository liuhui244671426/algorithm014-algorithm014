package Week_04

//岛屿数量
//思路:
//1.发现岛屿,count+1并就进入 dfs
//2.dfs 需要把 "当前" ,"前" ,"后" ,"左" ,"右" 各个节点都置成水,也就是 0
//注意: golang 中的 byte 如果转成 int ,需要加单引号''
var grid [][]byte = [][]byte{
	{'1', '1', '1', '1', '0'},
	{'1', '1', '0', '1', '0'},
	{'1', '1', '0', '0', '0'},
	{'0', '0', '0', '0', '0'},
}

// var grid [][]byte = [][]byte{
// 	{'1', '1', '0', '0', '0'},
// 	{'1', '1', '0', '0', '0'},
// 	{'0', '0', '1', '0', '0'},
// 	{'0', '0', '0', '1', '1'},
// }

func NumIslands() int {
	return numIslands(grid)
}

func numIslands(grid [][]byte) int {
	count := 0
	len_lie := len(grid) //列
	//fmt.Println(grid, len_lie, len_col)
	if len_lie == 0 {
		return 0
	}
	len_col := len(grid[0]) //行

	for i := 0; i < len_lie; i++ {
		for j := 0; j < len_col; j++ {
			//fmt.Println("ij", i, j)
			if grid[i][j] == '1' {
				count++
				dfs2(grid, i, j)
				//fmt.Println(grid)
			}
		}
	}
	return count
}

func dfs2(grid [][]byte, col, lie int) {
	len_col := len(grid)    //行
	len_lie := len(grid[0]) //列
	if col < 0 || lie < 0 || col >= len_col || lie >= len_lie || grid[col][lie] == '0' {
		return //越界条件,全部回退
	}
	grid[col][lie] = '0' //置 0

	dfs2(grid, col+1, lie) //下
	dfs2(grid, col-1, lie) //上
	dfs2(grid, col, lie-1) //左
	dfs2(grid, col, lie+1) //右
}
