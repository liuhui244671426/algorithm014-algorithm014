package Week_04

//买卖股票的最佳时机 II

func MaxProfit() int {
	var prices []int = []int{
		7, 1, 5, 3, 6, 4,
	}
	return maxProfit(prices)
}

func maxProfit(prices []int) int {
	var x int
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			x += prices[i] - prices[i-1]
		}
	}
	return x
}
