package Week_01

//移动零
//双指针法
//https://leetcode-cn.com/problems/move-zeroes/submissions/

func MoveZeroes(nums []int) {
	moveZeroes(nums)
}

func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
