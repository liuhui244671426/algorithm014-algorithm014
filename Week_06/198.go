package Week_06

import "fmt"

//打家劫舍

func Rob() int {
	var nums []int
	nums = []int{1, 2, 3, 1}
	//nums = []int{1, 2}
	nums = []int{2, 7, 9, 3, 1}
	return rob(nums)
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	var one int = nums[0]
	var two int = max(nums[1], nums[0])
	for i := 2; i < len(nums); i++ {
		one, two = two, max(two, nums[i]+one)
	}
	fmt.Println(one, two)
	return max(one, two)
}
