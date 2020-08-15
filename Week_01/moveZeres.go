//移动零
//双指针法
//https://leetcode-cn.com/problems/move-zeroes/submissions/

func moveZeroes(nums []int) {
	//fmt.Println(nums)
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
		//fmt.Println(nums, j, i)
	}
	//fmt.Println(nums)
}