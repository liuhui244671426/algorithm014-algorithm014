func climbStairs(n int) int {
	var dp []int = []int{1, 1}
	for i := 2; i <= n; i++ {
		fmt.Println(i, dp[i-1], dp[i-2])
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n]
}