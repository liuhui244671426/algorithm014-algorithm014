package Week_03

/**
思路:二分法+快慢指针(i,j)
时间复杂度: O(logN)
*/
func majorityElement(nums []int) int {
	l := len(nums)
	p := l / 2 //一分为二
	j := l - 1 //最后的元素
	var m map[int]int = map[int]int{}
	for i := 0; i <= p; i++ {
		m[nums[i]] += 1
		if i != j-i { //ij 相遇,只算一次
			m[nums[j-i]] += 1
		}
		if m[nums[i]] > p {
			return nums[i]
		}
		if m[nums[j-i]] > p {
			return nums[j-i]
		}
	}
	return 1
}

func MajorityElement(nums []int) int {
	return majorityElement(nums)
}
