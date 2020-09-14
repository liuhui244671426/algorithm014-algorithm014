package Week_06

//跳跃游戏 II
func Jump() int {
	var nums []int

	//nums = []int{1, 1, 1, 1}
	nums = []int{2, 3, 1, 1, 4}
	//nums = []int{4, 1, 1, 3, 1, 1, 1}
	//nums = []int{1, 2, 3}
	//nums = []int{1, 1, 1, 1}
	//nums = []int{1, 2, 1, 1, 1}
	//nums = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0}
	return jump(nums)
}

//贪心算法
func jump(nums []int) int {
	var ans int = 0
	var end int = 0    //边界
	var maxPos int = 0 //最大 postion
	for i := 0; i < len(nums)-1; i++ {
		if maxPos < nums[i]+i {
			maxPos = nums[i] + i //贪心取最大
		}
		if i == end {
			end = maxPos
			ans++
		}

	}
	return ans
}
