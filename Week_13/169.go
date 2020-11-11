package Week_13

//169. 多数元素

func MajorityElement() int {
	var nums []int
	nums = []int{
		2, 2, 1, 1, 1, 2, 2,
	}
	return majorityElement(nums)
}

func majorityElement(nums []int) int {
	var m map[int]int = make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for v, c := range m {
		if c > len(nums)/2 {
			return v
		}
	}
	return 0
}
