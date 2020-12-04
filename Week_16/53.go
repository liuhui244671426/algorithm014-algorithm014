package Week_16

//53. 最大子序和

func MaxSubArray() int {
	var nums []int = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	return maxSubArray(nums)
}

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i]) + nums[i]
		ans = max(dp[i], ans)
	}
	//fmt.Println(dp, ans)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
