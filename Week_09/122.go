package Week_09

//122. 买卖股票的最佳时机 II

func MaxProfit122() int {
	var prices []int
	prices = []int{7, 1, 5, 3, 6, 4}
	//prices = []int{7, 6, 4, 3, 1}
	return maxProfit122_2(prices)
}

//贪心法
func maxProfit122_1(prices []int) int {
	var len int = len(prices)
	var result int = 0
	for i := 1; i < len; i++ {
		if prices[i] > prices[i-1] {
			result += prices[i] - prices[i-1]
		}
	}
	return result
}

//第一版 dp
func maxProfit122_2(prices []int) int {
	var len int = len(prices)
	var dp [][]int = make([][]int, len)
	for i := 0; i < len; i++ {
		dp[i] = make([]int, 2) //1:持有 0:不持有
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) //卖
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) //买
	}
	//fmt.Println(dp)
	return dp[len-1][0]
}

//第二版 dp,压缩空间复杂度
func maxProfit122_3(prices []int) int {
	var cash int = 0          //不持有
	var hold int = -prices[0] //持有
	var preCash int = cash
	var preHold int = hold

	for i := 1; i < len(prices); i++ {
		cash = max(preCash, preHold+prices[i]) //卖
		hold = max(preHold, preCash-prices[i]) //买
		preCash, preHold = cash, hold
	}
	//fmt.Println(dp)
	return cash
}
