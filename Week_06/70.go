package Week_06

// 爬楼梯
/*
思路总结:
1.傻递归
2.cache递归
3.dp
4.优化空间的 dp
*/ 
var cache map[int]int = map[int]int{}

func ClimbStairs() int {
	var n int

	n = 3
	return climbStairs3(n)
}

//递归解 fib f(n) = f(n-1)+f(n-2)
func climbStairs1(n int) int {
	if n <= 2 {
		return n
	}
	if cache[n] > 0 {
		return cache[n]
	} else {
		cache[n] = climbStairs1(n-1) + climbStairs1(n-2)
	}
	return cache[n]
}

/*
DP规划,有优化空间,比如说,dp 空间复杂度压缩
*/
func climbStairs2(n int) int {
	var dp []int = make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

/*
基于 dp ,优化空间复杂度,因为 dp[i] 只与dp[i-1]和dp[i-2]有关,所以其他 dp 空间可以不需要
*/
func climbStairs3(n int) int {
	var one, two int = 1, 1
	for i := 2; i <= n; i++ {
		one, two = two, one+two
	}
	return two
}
