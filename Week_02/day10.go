//https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/
//滑动窗口的最大值

func maxSlidingWindow(nums []int, k int) []int {
	//fmt.Println(nums, k)
	leng := len(nums)
	var max_slice []int
	for i := 0; i < leng; i++ {
		if i+k <= leng {
			//fmt.Println(i, i+k, max(nums[i:i+k]))
			max_slice = append(max_slice, max(nums[i:i+k]))
		}
	}
	return max_slice
}

func max(nums []int) int {
	var max int = math.MinInt16
	for _, val := range nums {
		if val > max {
			max = val
		}
	}
	return max
}