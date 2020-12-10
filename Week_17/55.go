package Week_17

//55. 跳跃游戏

func CanJump() bool {
	var nums []int
	nums = []int{2, 3, 1, 1, 4} // true
	nums = []int{3, 2, 1, 0, 4} // false
	nums = []int{3, 2, 1, 1, 4} // true
	return canJump(nums)
}

/*
如果某一个作为 起跳点 的格子可以跳跃的距离是 3，那么表示后面 3 个格子都可以作为 起跳点。
可以对每一个能作为 起跳点 的格子都尝试跳一次，把 能跳到最远的距离 不断更新。
如果可以一直跳到最后，就成功了。
*/
func canJump(nums []int) bool {
	maxStep := 0
	for i := 0; i < len(nums); i++ {
		if i > maxStep {
			return false
		}
		maxStep = max(maxStep, i+nums[i])
	}
	return true
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
