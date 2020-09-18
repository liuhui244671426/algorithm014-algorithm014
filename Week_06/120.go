package Week_06

//三角形最小路径和

func MinimumTotal() int {
	var triangle [][]int
	triangle = [][]int{
		{3},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	return minimumTotal(triangle)
}

/*
自底向上 dp
*/
func minimumTotal(triangle [][]int) int {
	var dp [][]int = make([][]int, len(triangle)+1)
	for i := 0; i < len(triangle)+1; i++ {
		dp[i] = make([]int, len(triangle)+1)
	}
	//fmt.Println(triangle)
	for y := len(triangle) - 1; y >= 0; y-- {
		for x := 0; x < len(triangle[y]); x++ {
			//fmt.Println(y, x, triangle[y][x])
			dp[y][x] = min(dp[y+1][x], dp[y+1][x+1]) + triangle[y][x]
		}
	}
	//fmt.Println(dp)
	return dp[0][0]
}
