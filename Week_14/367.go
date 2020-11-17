package Week_14

//367. 有效的完全平方数

func IsPerfectSquare() bool {
	var num int = 15

	return isPerfectSquare(num)
}

//完全平方数->等差数列
//还可以试试 牛顿迭代法
func isPerfectSquare(num int) bool {
	n := 1
	for num > 0 {
		num -= n
		n += 2
	}
	return num == 0
}
