package Week_17

//279. 完全平方数

/*

输入: n = 12
输出: 3
解释: 12 = 4 + 4 + 4.

输入: n = 13
输出: 2
解释: 13 = 4 + 9.

*/
func NumSquares() int {
	var n int = 13
	return numSquares(n)
}

//假设最小公式值m = ƒ(n)
//那么n的值满足下列公式 ∑(A[i] * A[i]) = n
//令 k 为满足最小值 m 的时候，最大的平方数  。 令  d + k * k; = n ;  d >= 0;
//注意：一定要是满足m最小的时候的k值,一味的取最大平方数,就是贪心算法了
//得出 f(d) + f(k*k) = f(n);
//显然 f(k*k) = 1; 则  f(d) + 1 = f(n); 因为 d = n - k*k;
//则可以推出ƒ(n - k * k) + 1 = ƒ(n) ;  且 k * k <= n;
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; i-j*j >= 0; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	//fmt.Println(dp)
	return dp[n]
}
