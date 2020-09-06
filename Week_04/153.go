package Week_04

import "fmt"

//寻找旋转排序数组中的最小值
//思路:
//1.二分查找
//2.mid 不是必须加 1
func FindMin() int {
	var nums []int = []int{
		// 4, 5, 6, 7, 0, 1, 2,
		//3, 4, 5, 1, 2,
		//4, 5, 6, 7, 0, 1, 2,
		// 2, 1,
		//3, 1, 2,
		//2, 3, 4, 5, 1,
		4, 5, 1, 2, 3,
	}
	return findMin(nums)
}

func findMin(nums []int) int {
	length := len(nums)
	left := 0
	right := length - 1
	for left < right {
		mid := left + ((right - left) / 2)
		fmt.Println(nums, left, mid, right)
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
