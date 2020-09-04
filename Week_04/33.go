package Week_04

//二分查找
func Search() int {
	var nums []int = []int{1, 3}
	var target int = 0
	return search(nums, target)
}

//O(logN)
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	//fmt.Println("nums", nums)

	for left <= right {
		mid := left + ((right - left) / 2)
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[left] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
