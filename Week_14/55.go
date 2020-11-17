package Week_14

//55.跳跃游戏

func CanJump() bool {
	var nums []int = []int{
		// 2, 3, 1, 1, 4,
		3, 2, 1, 0, 4,
	}

	return canJump(nums)
}

func canJump(nums []int) bool {
	var cover int = 0
	if len(nums) == 1 {
		return true
	}
	for i := 0; i <= cover; i++ {
		cover = max(i+nums[i], cover)
		//fmt.Println(cover, i+nums[i])
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}
