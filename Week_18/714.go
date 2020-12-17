package Week_18

//714. 买卖股票的最佳时机含手续费

func MaxProfit() int {
	var prices []int = []int{1, 3, 2, 8, 4, 9}
	var fee int = 2
	return maxProfit(prices, fee)
}

func maxProfit(prices []int, fee int) int {
	dp := make([][2]int, len(prices))
	dp[0][1] = -prices[0] //持有
	dp[0][0] = 0          //不持有
	for i := 1; i < len(prices); i++ {
		/*
			对于今天不持有，可以从两个状态转移过来：1. 昨天也不持有；2. 昨天持有，今天卖出。两者取较大值。
			对于今天持有，可以从两个状态转移过来：1. 昨天也持有；2. 昨天不持有，今天买入。两者取较大值。
		*/
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return max(dp[len(prices)-1][0], dp[len(prices)-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
