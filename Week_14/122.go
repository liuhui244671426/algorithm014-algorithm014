package Week_14

//122. 买卖股票的最佳时机 II

func MaxProfit() int {
	var prices []int = []int{7, 1, 5, 3, 6, 4}
	return maxProfit(prices)
}

//最值-->动态规划
func maxProfit(prices []int) int {
	var dp [][]int = make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	//是否持有股票时的现金数
	//状态转移：cash → hold → cash → hold → cash → hold → cash
	dp[0][0] = 0          //持有现金
	dp[0][1] = -prices[0] //持有股票

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) //有现金,买
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) //有股票,卖
		//fmt.Println(dp[i])
	}
	return dp[len(dp)-1][0]
}
