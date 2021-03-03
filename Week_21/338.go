package Week_21

/*
338. 比特位计数
给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。
*/
func CountBits() []int {
	return countBits(5)
}

//牛逼
func countBits(num int) []int {
	var res []int = make([]int, num+1)
	for i := 1; i <= num; i++ {
		if i&1 == 1 {
			res[i] = res[i-1] + 1
		} else {
			res[i] = res[i/2]
		}
	}
	return res
}
