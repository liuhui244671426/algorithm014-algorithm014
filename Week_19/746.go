package Week_19

//746. 使用最小花费爬楼梯

func MinCostClimbingStairs() int {
	var cost []int = []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	// cost = []int{10, 15, 20}
	return minCostClimbingStairs(cost)
}

func minCostClimbingStairs(cost []int) int {

	for i := 2; i < len(cost); i++ {
		cost[i] = min(cost[i-1], cost[i-2]) + cost[i]
	}
	return min(cost[len(cost)-1], cost[len(cost)-2])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
