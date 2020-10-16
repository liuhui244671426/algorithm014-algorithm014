package Week_09

//300. 最长上升子序列

func LengthOfLIS() int {
	var nums []int
	nums = []int{10, 9, 2, 5, 3, 7, 101, 18}
	nums = []int{2, 2}
	nums = []int{4, 10, 4, 3, 8, 9}
	nums = []int{10, 9, 2, 5, 3, 7, 101, 18}
	return lengthOfLIS(nums)
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var dp []int = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1 //每个元素都是独立的最大子上升序列
	}
	//时间复杂度 O(n^2)
	//空间复杂度 O(n)
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ { //取 i 之前的最大子序列
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	//fmt.Println(dp)

	//这里必须求 dp 中的最大值,而不是 len-1
	var m int = 0
	for i := 0; i < len(dp); i++ {
		if m < dp[i] {
			m = dp[i]
		}
	}
	return m
}
