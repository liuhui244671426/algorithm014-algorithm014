package Week_10

//322. 零钱兑换

func CoinChange() int {
	var coins []int
	var amount int
	coins = []int{1, 2, 5, 7}
	amount = 14
	return coinChange(coins, amount)
}

func coinChange(coins []int, amount int) int {
	var dp []int = make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1 //凑 amount 金额,最多需要 amount 个 1 元
	}
	dp[0] = 0 //0 元不需要凑
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	//fmt.Println(dp)
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
