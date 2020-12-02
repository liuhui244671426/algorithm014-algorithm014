package Week_16

//70. 爬楼梯

func ClimbStairs() int {
	var n int = 10
	return climbStairs(n)
}

//递归->记忆递归->动态规划->压缩空间的动态规划
func climbStairs(n int) int {
	if n < 2 {
        return n
    }
	x,y := 1,1
	for i := 2; i < n; i++ {
		x,y = y,x+y
	}
	return x+y
}