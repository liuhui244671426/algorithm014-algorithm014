package Week_09

//123. 买卖股票的最佳时机 III

func MaxProfit123() int {
	var prices []int
	prices = []int{7, 1, 5, 3, 6, 4}
	//prices = []int{7, 6, 4, 3, 1}
	prices = []int{3, 3, 5, 0, 0, 3, 1, 4}
	return maxProfit123_1(prices)
}

func maxProfit123_1(prices []int) int {
	var len int = len(prices)
	var dp [][]int = make([][]int, len)
	for i := 0; i < len; i++ {
		dp[i] = make([]int, 5) //0，未交易，1，买入一次，2卖出1次，3买入2次，4卖出2次
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	//实际状态依赖与上一状态，以下情况实际不存
	dp[0][2] = 0
	dp[0][3] = -prices[0]
	dp[0][4] = 0

	for i := 1; i < len; i++ {
		dp[i][0] = 0
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	//fmt.Println(dp)
	return max(dp[len-1][2], dp[len-1][4])
}
