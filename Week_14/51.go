package Week_14

import (
	"strings"
)

//51. N 皇后
//todo 2020.10.27 done
//done 2020.11.18
func SolveNQueens() [][]string {
	var n int = 8
	return solveNQueens(n)
}

/*
回溯模板
result = []
def backtrack (路径,选择列表):
	if 满足条件:
		result = add(路径)
		return
	for 选择 in 选择列表:
		做选择
		backtrack(路径, 选择列表)
		撤销选择
*/
var result [][]string

func solveNQueens(n int) [][]string {
	var chess [][]string = make([][]string, n)
	result = make([][]string, 0)
	for i := 0; i < n; i++ {
		chess[i] = make([]string, n)
		for j := 0; j < n; j++ {
			chess[i][j] = "."
		}
	}

	solve(&result, chess, 0)
	return result
}

func solve(result *[][]string, chess [][]string, row int) {
	if row == len(chess) {
		var tmp []string = make([]string, len(chess))
		for i := 0; i < len(chess); i++ {
			tmp[i] = strings.Join(chess[i], "") //最终结果转成字符串
		}
		*result = append(*result, tmp)
		return
	}
	for col := 0; col < len(chess); col++ {
		//可以放 Q
		if valid(chess, row, col) {
			//老生常谈,回溯
			//空间可以再压缩,比如使用 bit
			chess[row][col] = "Q"
			solve(result, chess, row+1)
			chess[row][col] = "."
		}
	}
}

func valid(chess [][]string, row int, col int) bool {
	//行,有 Q
	for i := 0; i < row; i++ {
		if chess[i][col] == "Q" {
			return false
		}
	}
	//列,有 Q
	for i := 0; i < len(chess); i++ {
		if chess[row][i] == "Q" {
			return false
		}
	}
	//GO for 多变量,这样写!!!
	//当前坐标,右上角有 Q
	for i, j := row-1, col+1; i >= 0 && j < len(chess); i, j = i-1, j+1 {
		if chess[i][j] == "Q" {
			return false
		}
	}
	//当前坐标,左下角有 Q
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chess[i][j] == "Q" {
			return false
		}
	}
	return true
}
