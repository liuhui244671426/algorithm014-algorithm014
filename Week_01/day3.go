package Week_01

func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {

	leng := len(nums)
	var ret []int
	for i := 0; i < leng; i++ {
		if nums[i] == target && target != 0 {
			continue
		}
		new_target := target - nums[i]
		res := helper_twoSum(nums, i+1, new_target)
		if res != 0 {
			ret = append(ret, i)
			ret = append(ret, res)
		}
	}
	return ret
}

func helper_twoSum(nums []int, start int, target int) int {
	for i := start; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return 0
}
