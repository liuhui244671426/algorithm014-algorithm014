//https://leetcode-cn.com/problems/add-digits/submissions/
//各位相加
// 数学计算公式 : (n-1)%9+1
// 数学概念: 数根
package Week_02

func addDigits(num int) int {
	root := (num-1)%9 + 1
	//fmt.Println(root)
	return root
}

func AddDigits(num int) int {
	return addDigits(num)
}
