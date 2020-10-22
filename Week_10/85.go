package Week_10

//85. 最大矩形

/*
解法:可以从 84 题,演变过来.84 使用了单调栈,计算左右边界
*/
func MaximalRectangle() int {
	var matrix [][]byte
	matrix = [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	return maximalRectangle(matrix)
}

func maximalRectangle(matrix [][]byte) int {
	return 0
}
