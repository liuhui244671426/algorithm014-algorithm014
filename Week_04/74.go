package Week_04

//搜索二维矩阵
//思路: 1.遍历行 2.如果 target 进入当前行的数值范围 3.二分查找
func SearchMatrix() bool {
	var matrix [][]int = [][]int{
		{},
	}
	var target int = 1
	return searchMatrix(matrix, target)
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	left := 0
	len_col := len(matrix[left])
	len_lie := len(matrix)
	right := len_col - 1
	//fmt.Println("col, lie", len_col, len_lie)
	for i := 0; i < len_lie; i++ {
		//fmt.Println("i", i)
		if len(matrix[i]) == 0 {
			return false
		}
		if target >= matrix[i][left] && target <= matrix[i][right] {

			for left <= right {
				mid := left + (right-left)/2
				//fmt.Println("left, mid, right", left, mid, right, matrix[i][mid], target)
				if matrix[i][mid] > target {
					right = mid - 1
				} else if matrix[i][mid] < target {
					left = mid + 1
				} else if matrix[i][mid] == target {
					return true
				}
			}

		}
	}
	return false
}
