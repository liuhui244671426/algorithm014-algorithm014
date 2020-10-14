package Week_09

//198. 打家劫舍

func Rob() int {
	var nums []int
	nums = []int{2, 7, 9, 3, 1}
	//nums = []int{2, 1}
	return rob(nums)
}

/*
时间复杂度O(n)
空间复杂度O(1)

dp 方程 : dp[i] = max(dp[i-1], dp[i-2]+nums[i])
基于 dp 方程,可以实现时间复杂度O(n),空间复杂度O(n^2)的解法
但可以进一步压缩空间,我们发现 dp 方程只与 dp[i-1] dp[i-2] 有关,所以可以压缩空间
故 prev = dp[i-2] , curr = dp[i-1]
prev , curr = curr , max(curr, prev+nums[i])

*/
func rob(nums []int) int {
	var prev int = 0
	var curr int = 0
	for i := 0; i < len(nums); i++ {
		prev, curr = curr, max(curr, prev+nums[i])
	}
	return curr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
