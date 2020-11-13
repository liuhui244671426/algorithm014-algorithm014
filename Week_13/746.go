package Week_13

//746. 使用最小花费爬楼梯

func MinCostClimbingStairs() int {
	var cost []int

	cost = []int{
		10, 15, 20,
	}
	cost = []int{
		1, 100, 1, 1, 1, 100, 1, 1, 100, 1,
	}
	return minCostClimbingStairs(cost)
}

/*
读题,有求最值,第一个想到动态规划
*/
func minCostClimbingStairs(cost []int) int {
	/*
		//1.dp
		var dp []int = make([]int, len(cost))
		dp[0] = cost[0]
		dp[1] = cost[1]
		for i := 2; i < len(cost); i++ {
			dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
		}
		return min(dp[len(dp)-1], dp[len(dp)-2])
	*/
	//1.压缩空间的 dp
	for i := 2; i < len(cost); i++ {
		cost[i] = min(cost[i-2], cost[i-1]) + cost[i]
	}
	return min(cost[len(cost)-2], cost[len(cost)-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
