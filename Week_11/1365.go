package Week_11

//1365. 有多少小于当前数字的数字

func SmallerNumbersThanCurrent() []int {
	var nums []int
	nums = []int{8, 1, 2, 2, 3}
	return smallerNumbersThanCurrent(nums)
}
//todo 优化可以用计数排序
func smallerNumbersThanCurrent(nums []int) []int {
	var res []int

	for i := 0; i < len(nums); i++ {
		minCnt := 0
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue //不比较自己
			}
			if nums[i] > nums[j] {
				minCnt++
			}
		}
		res = append(res, minCnt)
	}
	//fmt.Println(res)
	return res
}
