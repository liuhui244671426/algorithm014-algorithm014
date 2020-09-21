package Week_07

//搜索二维矩阵 II

func SearchMatrix() bool {
	var matrix [][]int
	var target int
	// 数组特性
	// 每行的元素从左到右升序排列。
	// 每列的元素从上到下升序排列。
	matrix = [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target = 25

	// matrix = [][]int{
	// 	{-5},
	// }
	// target = -5
	return searchMatrix(matrix, target)
}

/*
 思路:
 1.根据特性,看完题目就想到了 二分法,但除了二分法或许还有其他方法,不妨试一下.
 2.根据特性,观察所得,从右上角开始
 	a.若 target>current , 向下移动
 	b.若 target<current , 向左移动
 	c.如此往复,直到走出矩阵
*/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	len_x := len(matrix[0]) - 1
	len_y := len(matrix) - 1
	y := 0
	x := len_x
	for y <= len_y && x >= 0 {
		//fmt.Println("y,x", y, x, matrix[y][x])
		if target < matrix[y][x] {
			x--
		} else if target > matrix[y][x] {
			y++
		} else if target == matrix[y][x] {
			return true
		}
	}
	return false
}
