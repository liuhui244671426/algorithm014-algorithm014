package Week_17

//45. 跳跃游戏 II

func Jump() int {
	var nums []int
	nums = []int{2, 3, 1, 1, 4} //result=2
	return jump(nums)
}

func jump(nums []int) int {
	ans := 0
	start := 0
	end := 1

	for end < len(nums) {
		maxPos := 0
		for i := start; i < end; i++ {
			maxPos = max(maxPos, i+nums[i])
		}
		start = end
		end = maxPos + 1
		ans++
	}
	return ans
}
