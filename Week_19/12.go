package Week_19

import (
	"strings"
)

//面试题 08.12. 八皇后

func SolveNQueens() [][]string {
	var n int = 4
	return solveNQueens(n)
}

func solveNQueens(n int) [][]string {
	var chess [][]string = make([][]string, n)
	var ans [][]string = make([][]string, 0)
	for i := 0; i < n; i++ {
		chess[i] = make([]string, n)
		for j := 0; j < n; j++ {
			chess[i][j] = "."
		}
	}
	backtracking1(&ans, chess, 0)
	return ans
}

func backtracking1(ans *[][]string, chess [][]string, row int) {
	if row == len(chess) {
		var tmp []string = make([]string, len(chess))
		for i := 0; i < len(chess); i++ {
			tmp[i] = strings.Join(chess[i], "") //最终结果转成字符串
		}
		*ans = append(*ans, tmp)
		return
	}
	for col := 0; col < len(chess); col++ {
		if valid(chess, row, col) {
			chess[row][col] = "Q"
			backtracking1(ans, chess, row+1)
			chess[row][col] = "."
		}
	}
}

func valid(grid [][]string, row int, col int) bool {
	//检测每一列
	for i := 0; i < col; i++ {
		if grid[row][i] == "Q" {
			return false
		}
	}
	//检测每一行
	for i := 0; i < row; i++ {
		if grid[i][col] == "Q" {
			return false
		}
	}
	//检测右上角
	for i, j := row-1, col+1; i >= 0 && j < len(grid); i, j = i-1, j+1 {
		if grid[i][j] == "Q" {
			return false
		}
	}
	//检测左上角
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if grid[i][j] == "Q" {
			return false
		}
	}
	return true
}
