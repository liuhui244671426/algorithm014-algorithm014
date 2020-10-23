package Week_10

//240. 搜索二维矩阵 II

func SearchMatrix() bool {
	var matrix [][]int
	var target int = 25
	matrix = [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	matrix = [][]int{}
	return searchMatrix(matrix, target)
}

/*
解法总结:
1.暴力法:查找每一行每一列(M*N)
2.分治法:在每一行进行二分查找O(MlogN)
3.走路法:这是自创名.基于对数据结构的分析,从右上角开始走,若target 小于当前则向左移动一位,若 target 大于当前则向右移动一位,若相同则跳出O(M+N)
*/
//走路法
func searchMatrix(matrix [][]int, target int) bool {
	var rows int = len(matrix)
	if rows == 0 {
		return false
	}
	var cols int = len(matrix[0])
	if cols == 0 {
		return false
	}
	var left int = cols - 1 //左
	var up int = 0          //右

	//fmt.Println("d1", left, up, matrix[up][left])
	for (left >= 0) && (up <= rows-1) {
		//fmt.Println("d2", left, up, matrix[up][left])
		if target > matrix[up][left] {
			up = up + 1
		} else if target < matrix[up][left] {
			left = left - 1
		} else if target == matrix[up][left] {
			return true
		}
	}
	//fmt.Println("d3", left, up)
	return false
}
