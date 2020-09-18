package Week_06

//零钱兑换
import (
	"fmt"
	"math"
	"sort"
)

func CoinChange() int {
	var coins []int
	var amount int

	/*
		[]

	*/
	// coins = []int{1, 2, 5}
	// amount = 11

	coins = []int{10, 7, 1}
	amount = 14
	return coinChange(coins, amount)
}

/*
func coinChange(coins []int, amount int) int {
	dp := make([][]int, len(coins)+1) //定义状态
	for i := 0; i <= len(coins); i++ {
		dp[i] = make([]int, amount+1)
	}
	//初始条件
	for j := 0; j <= amount; j++ {
		dp[0][j] = amount + 1
	}
	dp[0][0] = 0
	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			print(dp)
			//状态转移方程
			if j >= coins[i-1] {
				fmt.Println("min", i, j, coins[i-1])
				dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i-1]]+1)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	if dp[len(coins)][amount] > amount {
		return -1
	} else {
		return dp[len(coins)][amount]
	}
}
*/
func print(s [][]int) {
	for _, va := range s {
		fmt.Println(va)
	}
}

/*
//贪心算法:在有些情况下,失败!需改进
func coinChange(coins []int, amount int) int {
	var nums int = 0
	var i = len(coins) - 1 //多种面额
	sort.Ints(coins)       //从大到小排序

	for i >= 0 {
		//fmt.Println("coins", coins[i])
		if coins[i] <= amount {
			c := amount / coins[i]
			amount = amount - coins[i]*c
			fmt.Println("amount coins c ", amount, coins[i], c, nums)
			if amount >= 0 {
				nums += c
			}
		}
		//换下一个金额
		i--
	}
	fmt.Println(amount)
	if amount > 0 {
		nums = -1
	}
	return nums
}
*/

/*
贪心+递归
*/
var ans int

func coinChange(coins []int, amount int) int {
	ans = math.MaxInt32
	//fmt.Println(coins, amount)
	sort.Ints(coins)                              //从大到小排序
	findMin(coins, amount, len(coins)-1, 0, &ans) //先扣大额
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}
func findMin(coins []int, amount int, index int, count int, ans *int) {
	//fmt.Println("count", count, index)
	if amount == 0 { //余额是 0,则代表兑换完
		if count < *ans {
			*ans = count
		}
		return
	}
	if index < 0 {
		//fmt.Println("ans", *ans)
		return
	}
	/*
		1.用乘法加速 面额兑换的过程 k
		2.k>=0
	*/
	for k := amount / coins[index]; k >= 0 && k+count < *ans; k-- {
		//fmt.Println("for", coins, amount-k*coins[index], index-1, count+k)
		findMin(coins, amount-k*coins[index], index-1, count+k, ans) //用小额
	}
}
