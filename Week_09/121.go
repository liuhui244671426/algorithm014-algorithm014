package Week_09

//121. 买卖股票的最佳时机

func MaxProfit() int {
	var prices []int
	prices = []int{7, 1, 5, 3, 6, 4}
	//prices = []int{7, 6, 4, 3, 1}
	return maxProfit2(prices)
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
        return 0
    }
	var dp []int = make([]int, len(prices))
	var minPrice = prices[0]

	for i := 1; i < len(prices); i++ {
		minPrice = min(prices[i], minPrice)
		dp[i] = max(dp[i-1], prices[i]-minPrice)
	}
	//fmt.Println(dp)
	return dp[len(prices)-1]
}
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
        return 0
    }
	var minPrice = prices[0]
	var maxprofit int = 0
	for i := 1; i < len(prices); i++ {
		minPrice = min(prices[i], minPrice)
		maxprofit = max(maxprofit, prices[i]-minPrice)
	}
	return maxprofit
}
