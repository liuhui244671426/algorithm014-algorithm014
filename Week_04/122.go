package Week_04

//买卖股票的最佳时机 II

func MaxProfit() int {
	var prices []int = []int{
		7, 1, 5, 3, 6, 4,
	}
	return maxProfit(prices)
}

func maxProfit(prices []int) int {
	j := 0 //slow ptr
	var x int
	for i := 1; i < len(prices); i++ {
		//fmt.Println(j, i)
		if prices[j] > prices[i] {
			j = i
			continue
		}

		x += prices[i] - prices[j]
		j = i

	}
	//fmt.Println(x)
	return x
}
