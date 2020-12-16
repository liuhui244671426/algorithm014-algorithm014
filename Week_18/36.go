package Week_18

//36. 有效的数独

func IsValidSudoku() bool {
	var board [][]byte = [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	board = [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	return isValidSudoku(board)
}

func isValidSudoku(board [][]byte) bool {
	//9X9 grid
	//hash table
	var rows [10]int
	var cols [10]int
	var boxs [10]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			currentValue := 1 << (board[i][j] - '0')
			boxIndex := j/3 + (i/3)*3
			//A:没有优化空间
			// if rows[i][currentValue] == '1' || cols[j][currentValue] == '1' || boxs[boxIndex][currentValue] == '1' {
			// 	return false
			// }
			// rows[i][currentValue] = '1'
			// cols[j][currentValue] = '1'
			// boxs[boxIndex][currentValue] = '1'

			//B:二进制优化空间 ,位运算的或是|
			if (rows[i]&currentValue)|(cols[j]&currentValue)|(boxs[boxIndex]&currentValue) != 0 {
				return false
			}
			rows[i] |= currentValue
			cols[j] |= currentValue
			boxs[boxIndex] |= currentValue
		}
	}
	return true
}
