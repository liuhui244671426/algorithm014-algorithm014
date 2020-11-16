package Week_14

//53. 最大子序和

func MaxSubArray() int {
	var nums []int = []int{-2, 1}
	return maxSubArray(nums)
}

//求最值->动态规划 or 贪心
//选动态规划
func maxSubArray(nums []int) int {

	maxValue := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i], nums[i-1]+nums[i])
		maxValue = max(nums[i], maxValue)
	}
	return maxValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
