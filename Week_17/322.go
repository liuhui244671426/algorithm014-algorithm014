package Week_17

import (
	"math"
	"sort"
)

//322. 零钱兑换

func CoinChange() int {
	var coins []int = []int{2, 3, 5}
	var amount int = 54
	return coinChange2(coins, amount)
}

//dp
//可以 贪心+dfs 效率要高于 dp
func coinChange(coins []int, amount int) int {
	if len(coins) == 0 || amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt64
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && dp[i-coins[j]] != math.MaxInt64 {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a > b {
		return b

	}
	return a
}

//贪心 + dfs
func coinChange2(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	ans := math.MaxInt64
	dfs(coins, amount, 0, 0, &ans)

	if ans == math.MaxInt64 {
		return -1
	}
	return ans
}

func dfs(coins []int, amount int, cIndex int, count int, ans *int) {
	//中断条件
	if amount == 0 {
		*ans = min(*ans, count)
		return
	}
	//剪枝
	if cIndex == len(coins) {
		return
	}
	//用乘法加速
	for k := amount / coins[cIndex]; k >= 0 && k+count < *ans; k-- {
		dfs(coins, amount-k*coins[cIndex], cIndex+1, count+k, ans)
	}
}

//不用手动 reverse 了
// func reverse(s []int) []int {
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		s[i], s[j] = s[j], s[i]
// 	}
// 	return s
// }
