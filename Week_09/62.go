package Week_09

//62. 不同路径

func UniquePaths() int {
	var m int = 3
	var n int = 7
	return uniquePaths(m, n)
}

/*
1.重复子问题
2.无后效性
3.最优子结构

dp 方程
1.初始化状态
2.状态参数
3.决策

*/
func uniquePaths(m int, n int) int {
	//存在两个 状态参数 ,所以 dp 需要一个二维素组
	var dp [][]int = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 { //只能走一条路
				dp[i][j] = 1
			} else {
				//动态转移方程
				//当前节点一定是 "上"和"左" 的节点走过来的
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[n-1][m-1] //初始状态
}

// func uniquePaths1(m int, n int) int {
// 	var s int = n
// 	if m < n {
// 		s = m
// 	}
// 	var cur []int
// 	for i := 0; i < s; i++ {
// 		cur = append(cur, 1)
// 	}

// 	for i := 1; i < n; i++ {
// 		for j := 1; j < m; j++ {
// 			cur[j] += cur[j-1]
// 		}
// 	}
// 	fmt.Println(cur)
// 	return cur[n-1] //初始状态
// }
