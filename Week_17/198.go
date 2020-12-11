package Week_17

//198. 打家劫舍

func Rob() int {
	var nums []int = []int{2, 7, 9, 3, 1}
	nums = []int{2, 1, 1, 2}
	return rob(nums)
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i-1], dp[i-1])
	}
	//fmt.Println(nums, dp)
	return dp[len(nums)]
}
